// Net utilities - there's quite a lot of ceremony setting up a unix socket in a robust way.
package netutil

import (
	"fmt"
	"net"
	"os"

	"github.com/function61/gokit/os/osutil"
)

// Linux documentation: "Connecting to the socket object requires read/write permission."
// => RW for each desired (owner/group/other)
// not adding permutations that don't make sense: (), (group, other) and (other)
var (
	allowOwner         = osutil.FileMode(osutil.OwnerRW, osutil.GroupNone, osutil.OtherNone)
	allowOwnerAndGroup = osutil.FileMode(osutil.OwnerRW, osutil.GroupRW, osutil.OtherNone)
	allowEveryone      = osutil.FileMode(osutil.OwnerRW, osutil.GroupRW, osutil.OtherRW)
)

func ListenUnixAllowOwner(sockPath string, with func(net.Listener) error) error {
	return ListenUnixWithMode(sockPath, &allowOwner, with)
}

func ListenUnixAllowOwnerAndGroup(sockPath string, with func(net.Listener) error) error {
	return ListenUnixWithMode(sockPath, &allowOwnerAndGroup, with)
}

func ListenUnixAllowEveryone(sockPath string, with func(net.Listener) error) error {
	return ListenUnixWithMode(sockPath, &allowEveryone, with)
}

// pass nil os.FileMode if you don't want to Chmod() the socket file
//
// NOTE: `with()` is responsible for closing the obtained listener. for example the HTTP server
// launched from `with()` must be able to do graceful shutdown and `http.Server.Serve()` always closes the listener.
func ListenUnixWithMode(sockPath string, mode *os.FileMode, with func(net.Listener) error) error {
	withErr := func(err error) error { return fmt.Errorf("ListenUnix: %w", err) }

	exists, err := osutil.Exists(sockPath)
	if err != nil {
		return withErr(fmt.Errorf("exists: %w", err))
	}

	// clean-up old socket
	if exists {
		if err := os.Remove(sockPath); err != nil {
			return withErr(fmt.Errorf("cleanup previous socket: %w", err))
		}
	}

	listener, err := net.Listen("unix", sockPath)
	if err != nil {
		return withErr(err)
	}

	defer func() { // socket file was created. cleanup, so hopefully the next user doesn't need os.Remove()
		_ = os.Remove(sockPath)
	}()

	if mode != nil {
		if err := os.Chmod(sockPath, *mode); err != nil {
			return withErr(err)
		}
	}

	return with(listener)
}
