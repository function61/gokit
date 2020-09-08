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
	servicename          string
	args                 []string
	description          string
	docs                 []string
	envs                 []string
	requireNetworkOnline bool
	selfAbsolutePath     string
	err                  error // if error reading selfAbsolutePath
}

type optFn func(*serviceFile)

func SystemdServiceFile(servicename string, description string, opts ...optFn) serviceFile {
	selfAbsolutePath, err := filepath.Abs(os.Args[0])

	// why LOGGER_SUPPRESS_TIMESTAMPS? journalctl adds its own timestamps to each line,
	// so this is redundant data. respected by github.com/function61/gokit/logex
	sf := serviceFile{
		servicename:      servicename,
		description:      description,
		envs:             []string{"LOGGER_SUPPRESS_TIMESTAMPS=1"},
		selfAbsolutePath: selfAbsolutePath,
		err:              err,
	}

	for _, opt := range opts {
		opt(&sf)
	}

	return sf
}

func GetHints(sf serviceFile) string {
	return strings.Join([]string{
		"Wrote unit file to " + unitfilePath(sf),
		"Run to enable on boot & to start now:",
		"	$ systemctl enable " + sf.servicename,
		"	$ systemctl start " + sf.servicename,
		"	$ systemctl status " + sf.servicename,
	}, "\n")
}

func Install(sf serviceFile) error {
	if sf.err != nil {
		return sf.err
	}

	if _, errStat := os.Stat(unitfilePath(sf)); errStat == nil || !os.IsNotExist(errStat) {
		return fmt.Errorf("systemd service file %s already exists", unitfilePath(sf))
	}

	// https://unix.stackexchange.com/questions/433886/what-are-the-correct-permissions-for-a-systemd-service
	if err := ioutil.WriteFile(unitfilePath(sf), []byte(serialize(sf)), 0644); err != nil {
		return err
	}

	return nil
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

	l("")
	l("[Install]")
	l("WantedBy=multi-user.target")
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

func unitfilePath(sf serviceFile) string {
	return "/etc/systemd/system/" + sf.servicename + ".service"
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
