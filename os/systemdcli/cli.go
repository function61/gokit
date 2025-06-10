// creates an CLI entrypoint for managing a systemd-based service.
package systemdcli

import (
	"context"
	"fmt"
	"os"
	"os/exec"

	"github.com/function61/gokit/os/systemdinstaller"
	"github.com/spf13/cobra"
)

// `makeAdditionalCommands` can be used to give `WithInstallAndUninstallCommands()`
func Entrypoint(serviceName string, makeAdditionalCommands func(string) []*cobra.Command) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "service",
		Aliases: []string{"svc"},
		Short:   "Background service management",
		Args:    cobra.NoArgs,
	}

	if makeAdditionalCommands != nil {
		for _, cmdAdditional := range makeAdditionalCommands(serviceName) {
			cmd.AddCommand(cmdAdditional)
		}
	}

	// returns the raw error from `Run()`
	runSystemctlVerb := func(ctx context.Context, verb string) error {
		verbCmd := exec.CommandContext(ctx, "systemctl", verb, serviceName)
		verbCmd.Stdout = os.Stdout
		verbCmd.Stderr = os.Stderr
		return verbCmd.Run()
	}

	cmd.AddCommand(&cobra.Command{
		Use:   "start",
		Short: "Start the service",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runSystemctlVerb(cmd.Context(), "start")
		},
	})

	cmd.AddCommand(&cobra.Command{
		Use:   "stop",
		Short: "Stop the service",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runSystemctlVerb(cmd.Context(), "stop")
		},
	})

	cmd.AddCommand(&cobra.Command{
		Use:   "restart",
		Short: "Restart the service",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runSystemctlVerb(cmd.Context(), "restart")
		},
	})

	cmd.AddCommand(&cobra.Command{
		Use:   "status",
		Short: "Show status of the service",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			translateNonError := func(err error) error {
				if err != nil {
					// LSB dictates that successful status show of non-running program must return 3:
					// https://github.com/systemd/systemd/blob/997bc9ec568d291961a5ba6b7ef42ef7d4e19bd0/src/systemctl/systemctl-sysv-compat.h#L29
					if errExit, is := err.(*exec.ExitError); is && errExit.ExitCode() == 3 { // map to non-error
						return nil
					} else { // some other error
						return err
					}
				} else {
					return nil
				}
			}

			return translateNonError(runSystemctlVerb(cmd.Context(), "status"))
		},
	})

	cmd.AddCommand(&cobra.Command{
		Use:   "logs",
		Short: "Get logs for the service",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			//nolint:gosec // ok, is expected to not be user input.
			logsCmd := exec.CommandContext(cmd.Context(), "journalctl", "--unit="+serviceName)
			logsCmd.Stdout = os.Stdout
			logsCmd.Stderr = os.Stderr
			return logsCmd.Run()
		},
	})

	return cmd
}

func WithInstallAndUninstallCommands(makeSvc func(string) (*systemdinstaller.ServiceDefinition, error)) func(string) []*cobra.Command {
	return func(serviceName string) []*cobra.Command {
		return []*cobra.Command{
			{
				Use:   "install",
				Short: "Installs the background service",
				Args:  cobra.NoArgs,
				RunE: func(_ *cobra.Command, args []string) error {
					svc, err := makeSvc(serviceName)
					if err != nil {
						return err
					}

					if err := systemdinstaller.Install(*svc); err != nil {
						return err
					}

					fmt.Println(systemdinstaller.EnableAndStartCommandHints(*svc))

					return nil
				},
			},
			// TODO: add uninstall command
		}
	}
}
