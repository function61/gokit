package userutil

// Privilege elevation. Used for guaranteeing that we run as root, but if possible dropping back down
// to unprivileged user with the possibility to re-enter root again for short durations of privileged work.

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"syscall"

	. "github.com/function61/gokit/builtin"
	"github.com/function61/gokit/sliceutil"
)

func RequireRoot() (*ProofOfRunningAsRoot, error) {
	if !isRoot() {
		return nil, errors.New(`need root (tip: run with "$ sudo ...")`)
	} else {
		return &ProofOfRunningAsRoot{}, nil
	}
}

// used for UnprivilegedUser()
type UserAndGroup struct {
	uid               int
	gid               int
	gidsSupplementary []int
}

func NewUserAndGroup(uid int, gid int) UserAndGroup {
	return UserAndGroup{uid, gid, []int{}}
}

func (u UserAndGroup) Uid() int {
	return u.uid
}

func (u UserAndGroup) Gid() int {
	return u.gid
}

// can serve as "proof token" to functions of current execution context being as privileged user (root)
type ProofOfRunningAsRoot struct{}

// holder of this interface knows:
// - elevating temporarily to privileged user is possible
// - the base state *might be* running as unprivileged user, but could be privileged user also
type PrivilegedWork interface {
	// enter root for only the duration of work
	// proof is only valid during before work returns!
	AsRoot(work func(ProofOfRunningAsRoot) error) error
}

// holder of this interface knows:
// - elevating temporarily to privileged user is possible
// - the base state *is definitely* running as unprivileged user
type UnprivilegedUser interface {
	PrivilegedWork
	// get uid + gid of user who ran "$ sudo ..."
	UnprivilegedUser() UserAndGroup
}

// same as DropToUnprivilegedUser() but UnprivilegedUser is only an optional outcome, the only thing
// guaranteed is PrivilegedWork. this means that:
//
// - we definitely can elevate
// - but we might not be running as unprivileged user (we were not run under '$ sudo' so we cannot
//   jump between unprivileged and privileged contexts)
//
// WARNING: this alters global process state, so you shouldn't be doing anything concurrent.
// (at least where the different operations would be bothered by running in different security context)
func DropToUnprivilegedUserIfPossible() (PrivilegedWork, error) {
	if _, err := RequireRoot(); err != nil {
		return nil, err
	}

	// RequireRoot() doesn't yet guarantee we were run under sudo

	userDetails, err := getInvokingUserUidAndGidIfRunningInSudoWithSupplementary()
	switch {
	case err != nil: // should only happen if sudo ENV var numbers fail to parse
		return nil, ErrorWrap("PrivilegesDropTemporarily", err)
	case userDetails != nil:
		if err := setEffectiveUIDAndGID(*userDetails); err != nil {
			return nil, ErrorWrap("PrivilegesDropTemporarily", err)
		}

		return &runningUnderSudo{*userDetails}, nil
	default:
		return &runningAsRoot{}, nil
	}
}

// - requires that we were run under "$ sudo ..."
// - drops back to sudo-invoking user, so we're not running under full privileges
// - makes elevation possible for short work durations
//
// WARNING: this alters global process state, so you shouldn't be doing anything concurrent.
// (at least where the different operations would be bothered by running in different security context)
func DropToUnprivilegedUser() (UnprivilegedUser, error) {
	result, err := DropToUnprivilegedUserIfPossible()
	if err != nil {
		return nil, err
	}

	if unpriv, is := result.(UnprivilegedUser); is {
		return unpriv, nil
	} else {
		return nil, errors.New("need to be ran from '$ sudo ...' (just root will not do)")
	}
}

func isRoot() bool {
	return os.Geteuid() == 0
}

type runningAsRoot struct{}

var _ PrivilegedWork = (*runningAsRoot)(nil)

func (r *runningAsRoot) AsRoot(work func(ProofOfRunningAsRoot) error) error {
	return work(ProofOfRunningAsRoot{}) // already running as root -> no setup or cleanup required to enter
}

type runningUnderSudo struct {
	unprivileged UserAndGroup
}

var _ UnprivilegedUser = (*runningUnderSudo)(nil)

func (r *runningUnderSudo) UnprivilegedUser() UserAndGroup {
	return r.unprivileged
}

func (r *runningUnderSudo) AsRoot(work func(ProofOfRunningAsRoot) error) error {
	if err := regainRoot(); err != nil {
		return ErrorWrap("AsRoot", err)
	}

	returnAfterDroppingPrivileges := func(errWork error) error {
		if errDrop := setEffectiveUIDAndGID(r.unprivileged); errDrop != nil {
			if errDrop != nil {
				return fmt.Errorf("%w; additionally privilege drop failed: %v", errWork, errDrop)
			} else {
				return fmt.Errorf("AsRoot: work succeeded but privilege drop failed: %w", errDrop)
			}
		} else { // privilege drop succeeded
			// err might be:
			// a) nil (happy path) or
			// b) an error, but adding any error context is not as nothing in AsRoot() failed
			return errWork
		}
	}

	// actual work happens here, with privileged mode
	if err := work(ProofOfRunningAsRoot{}); err != nil {
		return returnAfterDroppingPrivileges(err) // error context not relevant
	}

	return returnAfterDroppingPrivileges(nil)
}

// effective uid and gid are used for files created by this process.
// https://en.wikipedia.org/wiki/User_identifier#Process_attributes
//
// this is useful for e.g. writing user's owned file on directory only root can write to.
//
// Running a process as root, before this function call (from /proc/self/status):
//     Uid:	0	0	0	0
//
//     (values are: "Real, effective, saved set, and filesystem UIDs")
//     https://man7.org/linux/man-pages/man5/proc.5.html
//
// After this function call:
//     Uid:	0	1000	0	1000
//
// => makes changes to (drops privileges of):
// - Real      : ☐
// - Effective : ☑
// - Saved set : ☐
// - Filesystem: ☑
//
// After regaining root (seteuid(0) and setegid(0)) the effect is reversed, i.e. back to starting situation.
func setEffectiveUIDAndGID(user UserAndGroup) error {
	return ErrorWrap("setEffectiveUIDAndGID", func() error {
		// gid has to be set first (otherwise we'd not be authorized to change it)
		if err := syscall.Setegid(user.gid); err != nil {
			return ErrorWrap("Setegid", err)
		}

		if err := syscall.Setgroups(user.gidsSupplementary); err != nil {
			return ErrorWrap("Setgroups", err)
		}

		if err := syscall.Seteuid(user.uid); err != nil {
			return ErrorWrap("Seteuid", err)
		}

		return nil
	}())
}

func regainRoot() error {
	return ErrorWrap("regainRoot", func() error {
		// uid has to be set first (otherwise we'd not be authorized to change it)
		if err := syscall.Seteuid(0); err != nil {
			return fmt.Errorf("Seteuid root: %w", err)
		}

		if err := syscall.Setegid(0); err != nil {
			return fmt.Errorf("Setegid root: %w", err)
		}

		return nil
	}())
}

// if running under '$ sudo', return invoking user's uid:gid pair.
// returns nil, nil if not running under sudo
func getInvokingUserUidAndGidIfRunningInSudoWithSupplementary() (*UserAndGroup, error) {
	// documented in https://www.sudo.ws/docs/man/sudo.man/#SUDO_UID
	if os.Getenv("SUDO_UID") == "" { // not running under sudo
		return nil, nil
	}

	uidSudo, err := strconv.Atoi(os.Getenv("SUDO_UID"))
	if err != nil {
		return nil, ErrorWrap("getInvokingUserUidAndGidIfRunningInSudoWithoutSupplementary", err)
	}

	gidSudo, err := strconv.Atoi(os.Getenv("SUDO_GID"))
	if err != nil {
		return nil, ErrorWrap("getInvokingUserUidAndGidIfRunningInSudoWithoutSupplementary", err)
	}

	// resolve also supplementary gids
	gidsSupplementarySudo, err := resolveSupplementaryGids(os.Getenv("SUDO_USER"))
	if err != nil {
		return nil, ErrorWrap("getInvokingUserUidAndGidIfRunningInSudoWithSupplementary", err)
	}

	return &UserAndGroup{uidSudo, gidSudo, gidsSupplementarySudo}, nil
}

// FIXME: copy-pasted from function22
func resolveSupplementaryGids(username string) ([]int, error) {
	if err := ErrorIfUnset(username == "", "username"); err != nil {
		return nil, err
	}

	// NOTE: Go deceptively has https://pkg.go.dev/os/user?utm_source=godoc#User.GroupIds
	//       which looks like something that we could use, BUT it only works with cgo.
	//       See https://github.com/golang/go/issues/19395

	groups := []int{}

	groupFile, err := os.Open("/etc/group")
	if err != nil {
		return nil, err
	}
	defer groupFile.Close()

	groupLines := bufio.NewScanner(groupFile)

	for groupLines.Scan() {
		// docker:x:129:joonas
		// docker:x:129:joonas,anotheruser
		parts := strings.Split(groupLines.Text(), ":")
		if len(parts) < 4 {
			return nil, fmt.Errorf("/etc/group invalid parts number: %d", len(parts))
		}

		// "joonas,anotheruser" => ["joonas", "anotheruser"]
		lineUsernames := strings.Split(parts[3], ",")
		if sliceutil.ContainsString(lineUsernames, username) {
			gid, err := strconv.Atoi(parts[2])
			if err != nil {
				return nil, err
			}

			groups = append(groups, gid)
		}
	}
	if err := groupLines.Err(); err != nil {
		return nil, err
	}

	return groups, nil
}
