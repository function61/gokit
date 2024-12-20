package cli

import (
	"context"
	"log"

	"github.com/function61/gokit/log/logex"
	"github.com/function61/gokit/os/osutil"
	"github.com/spf13/cobra"
)

// below intended to be deprecated soon

func RunnerNoArgs(run func(ctx context.Context, logger *log.Logger) error) func(*cobra.Command, []string) {
	return func(_ *cobra.Command, _ []string) {
		logger := logex.StandardLogger()

		osutil.ExitIfError(run(
			osutil.CancelOnInterruptOrTerminate(logger),
			logger))
	}
}

func Runner(run func(ctx context.Context, args []string, logger *log.Logger) error) func(*cobra.Command, []string) {
	return func(_ *cobra.Command, args []string) {
		logger := logex.StandardLogger()

		osutil.ExitIfError(run(
			osutil.CancelOnInterruptOrTerminate(logger),
			args,
			logger))
	}
}
