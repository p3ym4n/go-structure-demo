package log

import (
	"os"
	"strings"
)

type Level string

const (
	LevelDebug Level = "DEBUG"
	LevelInfo  Level = "INFO"
	LevelError Level = "ERROR"
	LevelFatal Level = "FATAL"
)

type Format string

const (
	FormatConsole Format = "CONSOLE"
	FormatJSON    Format = "JSON"
)

func configFromEnv() (Format, Level) {
	level := Level(strings.ToUpper(os.Getenv("LOG_LEVEL")))
	if level != LevelDebug && level != LevelInfo && level != LevelError && level != LevelFatal {
		level = LevelInfo
	}

	format := Format(strings.ToUpper(os.Getenv("LOG_FORMAT")))
	if format != FormatConsole && format != FormatJSON {
		format = FormatJSON
	}

	return format, level
}
