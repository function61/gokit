// Net utilities
package netutil

import (
	"context"
	"fmt"
	"net"
	"os"

	"github.com/function61/gokit/os/osutil"
)

// Linux documentation: "Connecting to the socket object requires read/write permission."
// => RW for each desired (owner/group/other)
// not adding permutations that don't make sense: (), (group, other) and (other)
var (
	UnixAllowOwner         = osutil.FileMode(osutil.OwnerRW, osutil.GroupNone, osutil.OtherNone)
	UnixAllowOwnerAndGroup = osutil.FileMode(osutil.OwnerRW, osutil.GroupRW, osutil.OtherNone)
	UnixAllowEveryone      = osutil.FileMode(osutil.OwnerRW, osutil.GroupRW, osutil.OtherRW)
)

// there's quite a lot of ceremony setting up a unix socket in a robust way.
func ListenUnix(
	ctx context.Context,
	sockPath string,
	mode *os.FileMode,
	with func(listener net.Listener) error,
) error {
	exists, err := osutil.Exists(sockPath)
	if err != nil {
		return fmt.Errorf("ListenUnix: exists: %w", err)
	}

	// clean-up old socket
	if exists {
		if err := os.Remove(sockPath); err != nil {
			return fmt.Errorf("ListenUnix: cleanup previous: %w", err)
		}
	}

	listener, err := net.Listen("unix", sockPath)
	if err != nil {
		return fmt.Errorf("ListenUnix: %w", err)
	}

	defer func() { // socket file was created. cleanup, so hopefully the next user doesn't need os.Remove()
		_ = os.Remove(sockPath)
	}()

	listenerCtx, cancel := context.WithCancel(ctx) // explained in next block
	defer cancel()

	// stop listener when:
	//
	// a) parent ctx done
	// b) we're exiting due to Chmod() failing
	// c) with() exits due to listener not closed but Accept() failing (kinda grey area - can that happen?)
	//    ^ this signalled by defer cancel()
	go func() {
		<-listenerCtx.Done()
		listener.Close()
	}()

	if mode != nil {
		if err := os.Chmod(sockPath, *mode); err != nil {
			return fmt.Errorf("ListenUnix: %w", err)
		}
	}

	return with(listener)
}
