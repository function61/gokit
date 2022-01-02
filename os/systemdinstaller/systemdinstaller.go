// Util for programmatically writing systemd unit file so your service can be autostarted
package systemdinstaller

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type serviceFile struct {
	serviceName      string
	args             []string
	description      string
	docs             []string
	envs             []string
	wants            []string
	after            []string
	bindsTo          []string
	userService      bool // systemd user-level service
	selfAbsolutePath string
	err              error // if error reading selfAbsolutePath
}

type Option func(*serviceFile)

func Service(serviceName string, description string, opts ...Option) serviceFile {
	return newService(serviceName, description, opts, false)
}

// user-level service
func UserService(serviceName string, description string, opts ...Option) serviceFile {
	return newService(serviceName, description, opts, true)
}

func newService(serviceName string, description string, opts []Option, userService bool) serviceFile {
	// filepath.Abs(os.Args[0]) fails with PATH-expanded lookups, os.Executable() resolves symlinks (bad for us)
	selfAbsolutePath, err := currentExecutableNoFollowSymlink()

	sf := serviceFile{
		userService:      userService,
		serviceName:      serviceName,
		description:      description,
		selfAbsolutePath: selfAbsolutePath,
		err:              err,
	}

	for _, opt := range opts {
		opt(&sf)
	}

	return sf
}

// installs the service
func Install(sf serviceFile) error {
	if sf.err != nil {
		return sf.err
	}

	filePath, err := unitfilePath(sf)
	if err != nil {
		return fmt.Errorf("unitfilePath: %w", err)
	}

	// user's service dir doesn't always exist
	if sf.userService {
		if err := os.MkdirAll(filepath.Dir(filePath), 0775); err != nil {
			return err
		}
	}

	if _, errStat := os.Stat(filePath); errStat == nil || !os.IsNotExist(errStat) {
		return fmt.Errorf("systemd service file %s already exists", filePath)
	}

	// https://unix.stackexchange.com/questions/433886/what-are-the-correct-permissions-for-a-systemd-service
	if err := ioutil.WriteFile(filePath, []byte(serialize(sf)), 0644); err != nil {
		// try to improve UX
		if errors.Is(err, os.ErrPermission) {
			return fmt.Errorf("%w\nHint: try prefix with '$ sudo ...'", err)
		} else {
			return err
		}
	}

	return nil
}

// gives, among others, command hints for how to start the installed service
func EnableAndStartCommandHints(sf serviceFile) string {
	maybeUserArg := func() string {
		if sf.userService {
			return " --user"
		} else {
			return ""
		}
	}()

	filePath, err := unitfilePath(sf)
	if err != nil { // shouldn't, because this is usually called after successful Install()
		return fmt.Sprintf("ERROR: %s", err.Error())
	}

	return strings.Join([]string{
		"Wrote unit file to " + filePath,
		"Run to enable on boot & to start (--)now:",
		"	$ systemctl" + maybeUserArg + " enable --now " + sf.serviceName,
		"Verify successful start:",
		"	$ systemctl" + maybeUserArg + " status " + sf.serviceName,
	}, "\n")
}

func serialize(sf serviceFile) string {
	lines := []string{}

	l := func(line string) {
		lines = append(lines, line)
	}

	l("[Unit]")
	l("Description=" + sf.description)

	if len(sf.docs) > 0 {
		l("Documentation=" + strings.Join(sf.docs, " "))
	}

	for _, wants := range sf.wants {
		l("Wants=" + wants)
	}

	for _, after := range sf.after {
		l("After=" + after)
	}

	for _, bindsTo := range sf.bindsTo {
		l("BindsTo=" + bindsTo)
	}

	wantedBy := func() string {
		if sf.userService { // https://unix.stackexchange.com/a/251225
			return "default.target"
		} else {
			return "multi-user.target"
		}
	}()

	l("")
	l("[Install]")
	l("WantedBy=" + wantedBy)
	l("")
	l("[Service]")
	l("ExecStart=" + strings.Join(append([]string{sf.selfAbsolutePath}, sf.args...), " "))
	l("WorkingDirectory=" + filepath.Dir(sf.selfAbsolutePath))
	l("Restart=always")
	l("RestartSec=10s")
	for _, env := range sf.envs {
		l("Environment=" + env)
	}

	return strings.Join(lines, "\n") + "\n"
}

func unitfilePath(sf serviceFile) (string, error) {
	// "example.service"
	unitFilename := sf.serviceName + ".service"

	if sf.userService {
		userConfigDir, err := os.UserConfigDir()
		if err != nil {
			return "", err
		}

		// ~/.config/systemd/user/example.service
		return filepath.Join(userConfigDir, "systemd", "user", unitFilename), nil
	} else {
		// /etc/systemd/system/example.service
		return "/etc/systemd/system/" + unitFilename, nil
	}
}

// FIXME(security): args are not shell escaped - DO NOT TAKE THIS FROM USER INPUT
func Args(args ...string) Option {
	return func(sf *serviceFile) {
		sf.args = args
	}
}

func Docs(docs ...string) Option {
	return func(sf *serviceFile) {
		sf.docs = docs
	}
}

func Env(key string, value string) Option {
	return func(sf *serviceFile) {
		sf.envs = append(sf.envs, key+"="+value)
	}
}

// https://unix.stackexchange.com/a/126146
func RequireNetworkOnline(sf *serviceFile) {
	Wants("network-online.target")(sf)
	After("network-online.target")(sf)
}

// https://www.freedesktop.org/software/systemd/man/systemd.unit.html#Wants=
func Wants(wants string) Option {
	return func(sf *serviceFile) {
		sf.wants = append(sf.wants, wants)
	}
}

// https://www.freedesktop.org/software/systemd/man/systemd.unit.html#Before=
func After(after string) Option {
	return func(sf *serviceFile) {
		sf.after = append(sf.after, after)
	}
}

// https://www.freedesktop.org/software/systemd/man/systemd.unit.html#BindsTo=
func BindsTo(to string) Option {
	return func(sf *serviceFile) {
		sf.bindsTo = append(sf.bindsTo, to)
	}
}

// systemd automatically dynamically generates units for network devices, so we can wait + bind to them.
// you can find interesting units by invoking $ systemctl list-unit
func WaitNetworkInterface(interfaceName string) Option {
	return func(sf *serviceFile) {
		interfaceDeviceUnit := fmt.Sprintf("sys-subsystem-net-devices-%s.device", interfaceName)

		// https://unix.stackexchange.com/a/417839
		// *BindsTo* makes the *After* even stronger. bind means that if the bound dependency goes
		// down, so should us go too. which is desirable with networked dependencies.
		BindsTo(interfaceDeviceUnit)(sf)
		After(interfaceDeviceUnit)(sf)
	}
}

// same as os.Executable() but does not follow symlinks (i.e. if /usr/bin/bob is returned even if it is a symlink to /tmp/bob).
// (the stdlib function follows symlinks on Linux)
func currentExecutableNoFollowSymlink() (string, error) {
	pathResolved, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}

	// LookPath() may still return relative path
	return filepath.Abs(pathResolved)
}
