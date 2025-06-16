package systemdinstaller

import (
	"os"
	"testing"

	"github.com/function61/gokit/testing/assert"
)

func TestBasic(t *testing.T) {
	sf := Service("testservice", "My cool service", Args("start"))
	sf = fixForTest(sf)

	assert.Equal(t, serialize(sf), `[Unit]
Description=My cool service

[Install]
WantedBy=multi-user.target

[Service]
ExecStart=/home/dummy/testservice_amd64 start
Type=exec
WorkingDirectory=/home/dummy
Restart=always
RestartSec=10s
`)

	assert.Equal(t, EnableAndStartCommandHints(sf), `Wrote unit file to /etc/systemd/system/testservice.service
Run to enable on boot & to start (--)now:
	$ systemctl enable --now testservice
Verify successful start:
	$ systemctl status testservice`)
}

func TestUserService(t *testing.T) {
	defer envVariableDuringTest("HOME", "/home/foobar")()

	sf := UserService("testservice", "My cool service", Args("start"))
	sf = fixForTest(sf)

	assert.Equal(t, serialize(sf), `[Unit]
Description=My cool service

[Install]
WantedBy=default.target

[Service]
ExecStart=/home/dummy/testservice_amd64 start
Type=exec
WorkingDirectory=/home/dummy
Restart=always
RestartSec=10s
`)

	assert.Equal(t, EnableAndStartCommandHints(sf), `Wrote unit file to /home/foobar/.config/systemd/user/testservice.service
Run to enable on boot & to start (--)now:
	$ systemctl --user enable --now testservice
Verify successful start:
	$ systemctl --user status testservice`)
}

func TestRequireNetworkOnline(t *testing.T) {
	sf := Service("testservice", "My cool service", Args("start"), RequireNetworkOnline)
	sf = fixForTest(sf)

	assert.Equal(t, serialize(sf), `[Unit]
Description=My cool service
Wants=network-online.target
After=network-online.target

[Install]
WantedBy=multi-user.target

[Service]
ExecStart=/home/dummy/testservice_amd64 start
Type=exec
WorkingDirectory=/home/dummy
Restart=always
RestartSec=10s
`)
}

func TestDocs(t *testing.T) {
	sf := Service("testservice", "My cool service", Docs("https://function61.com/", "https://github.com/function61/gokit"))
	sf = fixForTest(sf)

	assert.Equal(t, serialize(sf), `[Unit]
Description=My cool service
Documentation=https://function61.com/ https://github.com/function61/gokit

[Install]
WantedBy=multi-user.target

[Service]
ExecStart=/home/dummy/testservice_amd64
Type=exec
WorkingDirectory=/home/dummy
Restart=always
RestartSec=10s
`)
}

func TestEnv(t *testing.T) {
	sf := Service("testservice", "My cool service", Env("HOME", "/root"))
	sf = fixForTest(sf)

	assert.Equal(t, serialize(sf), `[Unit]
Description=My cool service

[Install]
WantedBy=multi-user.target

[Service]
ExecStart=/home/dummy/testservice_amd64
Type=exec
WorkingDirectory=/home/dummy
Restart=always
RestartSec=10s
Environment=HOME=/root
`)
}

func TestWaitNetworkInterface(t *testing.T) {
	sf := Service("testservice", "My cool service", WaitNetworkInterface("tailscale0"))
	sf = fixForTest(sf)

	assert.Equal(t, serialize(sf), `[Unit]
Description=My cool service
After=sys-subsystem-net-devices-tailscale0.device
BindsTo=sys-subsystem-net-devices-tailscale0.device

[Install]
WantedBy=multi-user.target

[Service]
ExecStart=/home/dummy/testservice_amd64
Type=exec
WorkingDirectory=/home/dummy
Restart=always
RestartSec=10s
`)
}

func TestTypeNotify(t *testing.T) {
	sf := Service("testservice", "My cool service", Type(ServiceTypeNotify))
	sf = fixForTest(sf)

	assert.Equal(t, serialize(sf), `[Unit]
Description=My cool service

[Install]
WantedBy=multi-user.target

[Service]
ExecStart=/home/dummy/testservice_amd64
Type=notify
WorkingDirectory=/home/dummy
Restart=always
RestartSec=10s
`)
}

func fixForTest(sf ServiceDefinition) ServiceDefinition {
	sf.selfAbsolutePath = "/home/dummy/testservice_amd64" // need to monkey patch this to get deterministic output
	return sf
}

func envVariableDuringTest(key string, value string) func() {
	oldValue, oldValueExisted := os.LookupEnv(key)

	// change for the duration of the test
	if err := os.Setenv(key, value); err != nil {
		panic(err)
	}

	// cleanup returns variable to old state
	return func() {
		if oldValueExisted {
			if err := os.Setenv(key, oldValue); err != nil {
				panic(err)
			}
		} else { // did not exist (not same as empty value)
			if err := os.Unsetenv(key); err != nil {
				panic(err)
			}
		}
	}
}
