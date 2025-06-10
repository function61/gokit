// Making CLI commands have some quality without too much boilerplate.
package cli

import (
	"os"

	"github.com/function61/gokit/app/dynversion"
	"github.com/function61/gokit/os/osutil"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// wraps the `Execute()` call of the command to inject boilerplate details for root command:
// - `Use` (for root cmd should be `os.Args[0]`)
// - `Version` (should be program's built-in version info)
// - UX improvements (some Cobra defaults are in poor taste)
// - configure logging
// - deliver context & cancel it on signals
func Execute(app *cobra.Command) {
	// dirty to mutate after-the-fact

	app.Use = os.Args[0]
	app.Version = dynversion.Version

	// hide the default "completion" subcommand from polluting UX (it can still be used). https://github.com/spf13/cobra/issues/1507
	app.CompletionOptions = cobra.CompletionOptions{HiddenDefaultCmd: true}

	// disable alternate `help <command>` route from polluting UX.
	// the de facto `--help` can still be used and pro-tip for that is prominently visible.
	// https://github.com/spf13/cobra/issues/587#issuecomment-843747825
	app.SetHelpCommand(&cobra.Command{Hidden: true})

	// don't display error internally by Cobra. we do better job of displaying it here (and most importantly, setting exit code)
	app.SilenceErrors = true

	// child commands will "inherit" this.
	// the `cmd` will be the child command (and not root), which is the one we need.
	app.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		// Cobra would displays CLI usage help for any error happened inside (= higher-level app errors) `Run` func.
		//
		// in reality we want the usage to only appear on failed input validation, because it would be confusing to
		// show CLI usage help for app-level errors happening above the CLI layer.
		//
		// since we arrived at the run family of funcs we know the error won't be due to input validation anymore.
		cmd.SilenceUsage = true

		// handle logging configuration. needs to be done here because `AddLogLevelControls()` instructs Cobra to
		// possibly modify our global log level verbosity flag but it is only mutated just before run.
		configureLogging()
	}

	// cannot `AddLogLevelControls(app.Flags())` here because it gets confusing if:
	// a) the root command is not runnable
	// b) the root command is runnable BUT it doesn't do logging (or there is no debug-level logs to suppress)

	// handle signals
	ctx := notifyContextInterruptOrTerminate()

	// this is where the magic happens
	_, err := app.ExecuteContextC(ctx)

	osutil.ExitIfError(err)
}

// adds CLI flags that control the logging level
func AddLogLevelControls(flags *pflag.FlagSet) {
	flags.BoolVarP(&logLevelVerbose, "verbose", "v", logLevelVerbose, "Include debug-level logs")

	// TODO: maybe add a "quiet" level as well
}
