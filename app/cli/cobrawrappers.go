// Making CLI commands have some quality without too much boilerplate.
package cli

import (
	"context"
	"os"

	"github.com/function61/gokit/app/dynversion"
	"github.com/function61/gokit/os/osutil"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// wraps the `Execute()` call of the command to inject boilerplate details like `Use`, `Version` and
// handling of error to `Command.Execute()` (such as flag validation, missing command etc.)
func Execute(app *cobra.Command) {
	// dirty to mutate after-the-fact

	app.Use = os.Args[0]
	app.Version = dynversion.Version
	// hide the default "completion" subcommand from polluting UX (it can still be used). https://github.com/spf13/cobra/issues/1507
	app.CompletionOptions = cobra.CompletionOptions{HiddenDefaultCmd: true}

	// cannot `AddLogLevelControls(app.Flags())` here because it gets confusing if:
	// a) the root command is not runnable
	// b) the root command is runnable BUT it doesn't do logging (or there is no debug-level logs to suppress)

	osutil.ExitIfError(app.Execute())
}

// fixes problems of cobra commands' bare run callbacks with regards to application quality:
// 1. logging not configured
// 2. no interrupt handling
// 3. no error handling
func WrapRun(run func(ctx context.Context, args []string) error) func(*cobra.Command, []string) {
	return func(_ *cobra.Command, args []string) {
		// handle logging
		configureLogging()

		// handle interrupts
		ctx := notifyContextInterruptOrTerminate()

		// run the actual code (jump from CLI context to higher-level application context)
		// this can be kinda read as:
		//  output = logic(input)
		err := run(ctx, args)

		// handle errors
		osutil.ExitIfError(err)
	}
}

// adds CLI flags that control the logging level
func AddLogLevelControls(flags *pflag.FlagSet) {
	flags.BoolVarP(&logLevelVerbose, "verbose", "v", logLevelVerbose, "Include debug-level logs")

	// TODO: maybe add a "quiet" level as well
}
