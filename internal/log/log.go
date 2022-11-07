package log

import (
	"context"
	"fmt"
	stdlog "log"
	"strings"
)

type Logger interface {
	GetStd() *stdlog.Logger
	GetLevel() Level
	GetFormat() Format

	Debug(msg string, args ...interface{})
	DebugWithContext(ctx context.Context, msg string, args ...interface{})
	Info(msg string, args ...interface{})
	InfoWithContext(ctx context.Context, msg string, args ...interface{})
	Error(msg string, args ...interface{})
	ErrorWithContext(ctx context.Context, msg string, args ...interface{})
	Fatal(msg string, args ...interface{})
	FatalWithContext(ctx context.Context, msg string, args ...interface{})
}

func formatMessageWithArgs(msg string, args []interface{}) (string, []interface{}) {
	thresholds := strings.Count(msg, "%")
	if thresholds > len(args) {
		thresholds = len(args)
	}
	if thresholds > 0 {
		msg = fmt.Sprintf(msg, args[:thresholds]...)
		if len(args) > thresholds {
			args = args[thresholds:]
		} else {
			args = make([]interface{}, 0)
		}
	} else if args == nil {
		args = make([]interface{}, 0)
	}
	return msg, args
}
