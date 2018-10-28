package systemdinstaller

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// FIXME(security): args are not shell escaped - DO NOT TAKE THIS FROM USER INPUT
func InstallSystemdServiceFile(servicename string, args []string, description string) (string, error) {
	unitfilePath := "/etc/systemd/system/" + servicename + ".service"

	unitTemplate := `[Unit]
Description=%s

[Install]
WantedBy=multi-user.target

[Service]
ExecStart=%s
WorkingDirectory=%s
Restart=always
RestartSec=10s
`

	selfAbsolutePath, errAbs := filepath.Abs(os.Args[0])
	if errAbs != nil {
		return "", errAbs
	}

	cmd := append([]string{selfAbsolutePath}, args...)

	unitContent := fmt.Sprintf(
		unitTemplate,
		description,
		strings.Join(cmd, " "),
		filepath.Dir(selfAbsolutePath))

	if _, errStat := os.Stat(unitfilePath); errStat == nil || !os.IsNotExist(errStat) {
		return "", errors.New("systemd service file already exists!")
	}

	if err := ioutil.WriteFile(unitfilePath, []byte(unitContent), 0755); err != nil {
		return "", err
	}

	hints := []string{
		"Wrote unit file to " + unitfilePath,
		"Run to enable on boot & to start now:",
		"	$ systemctl enable " + servicename,
		"	$ systemctl start " + servicename,
	}

	return strings.Join(hints, "\n"), nil
}
