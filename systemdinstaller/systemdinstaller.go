// Util for programmatically writing systemd unit file so your service can be autostarted
package systemdinstaller

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type serviceFile struct {
	servicename          string
	args                 []string
	description          string
	requireNetworkOnline bool
	selfAbsolutePath     string
	err                  error // if error reading selfAbsolutePath
}

type optFn func(*serviceFile)

func SystemdServiceFile(servicename string, description string, opts ...optFn) serviceFile {
	selfAbsolutePath, err := filepath.Abs(os.Args[0])

	sf := serviceFile{
		servicename:      servicename,
		description:      description,
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

// FIXME(security): args are not shell escaped - DO NOT TAKE THIS FROM USER INPUT
func Install(sf serviceFile) error {
	if sf.err != nil {
		return sf.err
	}

	if _, errStat := os.Stat(unitfilePath(sf)); errStat == nil || !os.IsNotExist(errStat) {
		return errors.New("systemd service file already exists!")
	}

	if err := ioutil.WriteFile(unitfilePath(sf), []byte(serialize(sf)), 0755); err != nil {
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
	l("Environment=LOGGER_SUPPRESS_TIMESTAMPS=1")
	// why LOGGER_SUPPRESS_TIMESTAMPS? journalctl adds its own timestamps to each line,
	// so this is redundant data

	return strings.Join(lines, "\n") + "\n"
}

func unitfilePath(sf serviceFile) string {
	return "/etc/systemd/system/" + sf.servicename + ".service"
}

func Args(args ...string) optFn {
	return func(sf *serviceFile) {
		sf.args = args
	}
}

func RequireNetworkOnline(sf *serviceFile) {
	sf.requireNetworkOnline = true
}
