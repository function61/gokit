package cli

import (
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"
	"github.com/mattn/go-isatty"
)

var (
	logLevelVerbose = false
	discardAttr     = slog.Attr{} // zero `Attr` means discard
)

func configureLogging() {
	logLevel := func() slog.Level {
		if logLevelVerbose {
			return slog.LevelDebug
		} else {
			return slog.LevelInfo
		}
	}()

	addSource := func() bool {
		if logLevelVerbose {
			return true
		} else {
			return false
		}
	}()

	errorStream := os.Stderr
	errorStreamIsUserTerminal := isatty.IsTerminal(errorStream.Fd())

	logHandler := func() slog.Handler {
		if errorStreamIsUserTerminal { // output format optimized to looking at from terminal
			return tint.NewHandler(errorStream, &tint.Options{
				Level:      logLevel,
				AddSource:  addSource,
				TimeFormat: time.TimeOnly, // not using freedom time (`time.Kitchen`)
				// intentionally not giving `ReplaceAttr` because for terminal we can always include times
			})
		} else { // "production" log output
			logAttrReplacer := timeRemoverAttrReplacer
			if !logsShouldOmitTime() {
				logAttrReplacer = nil
			}

			return slog.NewTextHandler(errorStream, &slog.HandlerOptions{
				Level:       logLevel,
				AddSource:   addSource,
				ReplaceAttr: logAttrReplacer,
			})
		}
	}()

	// expecting the apps to just use the global logger
	slog.SetDefault(slog.New(logHandler))
}

// if our logs are redirected to journald or similar which already add timestamps don't add double timestamps
func logsShouldOmitTime() bool {
	// "This permits invoked processes to safely detect whether their standard output or standard
	// error output are connected to the journal."
	// https://www.freedesktop.org/software/systemd/man/systemd.exec.html#%24JOURNAL_STREAM
	systemdJournal := os.Getenv("JOURNAL_STREAM") != ""

	// explicitly asked, e.g. set by orchestrator when running in Docker with log redirection taken care of
	explicitSuppress := os.Getenv("LOGGER_SUPPRESS_TIMESTAMPS") == "1"

	return systemdJournal || explicitSuppress
}

func timeRemoverAttrReplacer(groups []string, a slog.Attr) slog.Attr {
	if a.Key == slog.TimeKey {
		return discardAttr
	} else {
		return a
	}
}
