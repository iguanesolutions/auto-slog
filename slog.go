package autoslog

import (
	"log/slog"
	"os"
	"strings"

	sysd "github.com/iguanesolutions/go-systemd/v5"
	sysdjdslog "github.com/iguanesolutions/go-systemd/v5/journald/slog"
	"github.com/mattn/go-isatty"
)

// LogLevel is a wrapper to get a log level from a string. Case insensitive:
// * DEBUG
// * INFO
// * WARN
// * ERROR
func LogLevel(logLevelRequested string) slog.Level {
	switch strings.ToUpper(logLevelRequested) {
	case slog.LevelDebug.String():
		return slog.LevelDebug
	case slog.LevelWarn.String():
		return slog.LevelWarn
	case slog.LevelError.String():
		return slog.LevelError
	case slog.LevelInfo.String(), "":
		fallthrough
	default:
		return slog.LevelInfo
	}
}

// NewLogger spawn a new structured logger with automatically choosen handler:
// * If the output is a terminal, a standard text handler will be used
// * If the program has bee started by systemd, a custom journald handler will be used
// * Otherwise a standard JSON handler will be used
func NewLogger(logLevel slog.Level) (logger *slog.Logger) {
	// If output is a terminal, use a regular text logger
	if isatty.IsTerminal(os.Stdout.Fd()) {
		return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: logLevel,
		}))
	}
	// If started by systemd, use a custom journald slog handler
	if _, sysdStarted := sysd.GetInvocationID(); sysdStarted {
		return slog.New(sysdjdslog.NewHandler(logLevel))
	}
	// Otherwise, use a JSON logger
	return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: logLevel,
	}))
}
