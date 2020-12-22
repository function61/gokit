// Util for programmatically writing systemd unit file so your service can be autostarted
package systemdinstaller

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type serviceFile struct {
	serviceName          string
	args                 []string
	description          string
	docs                 []string
	envs                 []string
	requireNetworkOnline bool
	userService          bool // systemd user-level service
	selfAbsolutePath     string
	err                  error // if error reading selfAbsolutePath
}

type optFn func(*serviceFile)

func Service(serviceName string, description string, opts ...optFn) serviceFile {
	return newService(serviceName, description, opts, false)
}

// user-level service
func UserService(serviceName string, description string, opts ...optFn) serviceFile {
	return newService(serviceName, description, opts, true)
}

func newService(serviceName string, description string, opts []optFn, userService bool) serviceFile {
	selfAbsolutePath, err := filepath.Abs(os.Args[0])

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
		return err
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
		"Run to enable on boot & to start now:",
		"	$ systemctl" + maybeUserArg + " enable " + sf.serviceName,
		"	$ systemctl" + maybeUserArg + " start " + sf.serviceName,
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

	// https://unix.stackexchange.com/a/126146
	if sf.requireNetworkOnline {
		l("Wants=network-online.target")
		l("After=network-online.target")
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
func Args(args ...string) optFn {
	return func(sf *serviceFile) {
		sf.args = args
	}
}

func Docs(docs ...string) optFn {
	return func(sf *serviceFile) {
		sf.docs = docs
	}
}

func Env(key string, value string) optFn {
	return func(sf *serviceFile) {
		sf.envs = append(sf.envs, key+"="+value)
	}
}

func RequireNetworkOnline(sf *serviceFile) {
	sf.requireNetworkOnline = true
}
