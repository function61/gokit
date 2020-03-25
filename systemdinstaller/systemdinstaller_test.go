package systemdinstaller

import (
	"testing"

	"github.com/function61/gokit/assert"
)

func TestBasic(t *testing.T) {
	sf := SystemdServiceFile("testservice", "My cool service", Args("start"))
	sf = fixForTest(sf)

	assert.EqualString(t, serialize(sf), `[Unit]
Description=My cool service

[Install]
WantedBy=multi-user.target

[Service]
ExecStart=/home/dummy/testservice_amd64 start
WorkingDirectory=/home/dummy
Restart=always
RestartSec=10s
Environment=LOGGER_SUPPRESS_TIMESTAMPS=1
`)

	assert.EqualString(t, GetHints(sf), `Wrote unit file to /etc/systemd/system/testservice.service
Run to enable on boot & to start now:
	$ systemctl enable testservice
	$ systemctl start testservice
	$ systemctl status testservice`)
}

func TestRequireNetworkOnline(t *testing.T) {
	sf := SystemdServiceFile("testservice", "My cool service", Args("start"), RequireNetworkOnline)
	sf = fixForTest(sf)

	assert.EqualString(t, serialize(sf), `[Unit]
Description=My cool service
Wants=network-online.target
After=network-online.target

[Install]
WantedBy=multi-user.target

[Service]
ExecStart=/home/dummy/testservice_amd64 start
WorkingDirectory=/home/dummy
Restart=always
RestartSec=10s
Environment=LOGGER_SUPPRESS_TIMESTAMPS=1
`)
}

func TestDocs(t *testing.T) {
	sf := SystemdServiceFile("testservice", "My cool service", Docs("https://function61.com/", "https://github.com/function61/gokit"))
	sf = fixForTest(sf)

	assert.EqualString(t, serialize(sf), `[Unit]
Description=My cool service
Documentation=https://function61.com/ https://github.com/function61/gokit

[Install]
WantedBy=multi-user.target

[Service]
ExecStart=/home/dummy/testservice_amd64
WorkingDirectory=/home/dummy
Restart=always
RestartSec=10s
Environment=LOGGER_SUPPRESS_TIMESTAMPS=1
`)
}

func fixForTest(sf serviceFile) serviceFile {
	sf.selfAbsolutePath = "/home/dummy/testservice_amd64" // need to monkey patch this to get deterministic output
	return sf
}
