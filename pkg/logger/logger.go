package logger

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog"
)

type Logger struct {
	logger zerolog.Logger
}

func New(level string) *Logger {
	var lvl zerolog.Level
	switch strings.ToLower(level) {
	case "debug":
		lvl = zerolog.DebugLevel
	case "info":
		lvl = zerolog.InfoLevel
	case "warn":
		lvl = zerolog.WarnLevel
	case "error":
		lvl = zerolog.ErrorLevel
	default:
		lvl = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(lvl)
	zlogger := zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()

	return &Logger{logger: zlogger}
}

// Public logging methods — compatible with interface
func (l *Logger) Debug(msg string, args ...any) {
	l.logger.Debug().Msgf("%s", format(msg, args...))
}

func (l *Logger) Info(msg string, args ...any) {
	l.logger.Info().Msgf("%s", format(msg, args...))
}

func (l *Logger) Warn(msg string, args ...any) {
	l.logger.Warn().Msgf("%s", format(msg, args...))
}

func (l *Logger) Error(msg string, args ...any) {
	l.logger.Error().Msgf("%s", format(msg, args...))
}

func (l *Logger) Fatal(msg string, args ...any) {
	l.logger.Fatal().Msgf("%s", format(msg, args...))
	os.Exit(1)
}

// ---------- Helper ----------
func format(msg string, args ...any) string {
	if len(args) == 0 {
		return msg
	}
	return fmt.Sprintf(msg, args...)
}
