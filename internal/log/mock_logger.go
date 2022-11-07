package log

import (
	"context"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	stdlog "log"
	"sync"
)

var _ Logger = (*MockLogger)(nil)

type MockLogger struct {
	mu          sync.Mutex // Avoid race conditions for setting last log
	Name        string
	Messages    []string
	LastMessage string
	LastItems   map[string]interface{}
	WasFatal    bool
	Level       Level
	Format      Format
	Events      []string
	LastEvent   string
}

func NewMock(name string) Logger {
	format, level := configFromEnv()
	return &MockLogger{
		Name:        name,
		Messages:    make([]string, 0),
		LastMessage: "",
		Level:       level,
		Format:      format,
		LastItems:   make(map[string]interface{}),
		WasFatal:    false,
		Events:      make([]string, 0),
		LastEvent:   "",
	}
}

func (l *MockLogger) GetStd() *stdlog.Logger {
	return stdlog.Default()
}

func (l *MockLogger) GetLevel() Level {
	return l.Level
}

func (l *MockLogger) GetFormat() Format {
	return l.Format
}

func (l *MockLogger) executeInLock(f func()) {
	l.mu.Lock()
	defer l.mu.Unlock()
	f()
}

func (l *MockLogger) Debug(msg string, args ...interface{}) {
	msg, fields := buildMsgAndArgsMock(msg, args...)
	l.executeInLock(func() {
		l.Level = LevelDebug
		l.Messages = append(l.Messages, msg)
		l.LastMessage = msg
		l.LastItems = fields
		l.WasFatal = false
	})
}

func (l *MockLogger) DebugWithContext(ctx context.Context, msg string, args ...interface{}) {
	msg, fields := buildMsgAndArgsWithContextMock(ctx, msg, args...)
	l.executeInLock(func() {
		l.Level = LevelDebug
		l.Messages = append(l.Messages, msg)
		l.LastMessage = msg
		l.LastItems = fields
		l.WasFatal = false
	})
}

func (l *MockLogger) Info(msg string, args ...interface{}) {
	msg, fields := buildMsgAndArgsMock(msg, args...)
	l.executeInLock(func() {
		l.Level = LevelInfo
		l.Messages = append(l.Messages, msg)
		l.LastMessage = msg
		l.LastItems = fields
		l.WasFatal = false
	})
}

func (l *MockLogger) InfoWithContext(ctx context.Context, msg string, args ...interface{}) {
	msg, fields := buildMsgAndArgsWithContextMock(ctx, msg, args...)
	l.executeInLock(func() {
		l.Level = LevelInfo
		l.Messages = append(l.Messages, msg)
		l.LastMessage = msg
		l.LastItems = fields
		l.WasFatal = false
	})
}

func (l *MockLogger) Error(msg string, args ...interface{}) {
	msg, fields := buildMsgAndArgsMock(msg, args...)
	l.executeInLock(func() {
		l.Level = LevelError
		l.Messages = append(l.Messages, msg)
		l.LastMessage = msg
		l.LastItems = fields
		l.WasFatal = false
	})
}

func (l *MockLogger) ErrorWithContext(ctx context.Context, msg string, args ...interface{}) {
	msg, fields := buildMsgAndArgsWithContextMock(ctx, msg, args...)
	l.executeInLock(func() {
		l.Level = LevelError
		l.Messages = append(l.Messages, msg)
		l.LastMessage = msg
		l.LastItems = fields
		l.WasFatal = false
	})
}

func (l *MockLogger) Fatal(msg string, args ...interface{}) {
	msg, fields := buildMsgAndArgsMock(msg, args...)
	l.executeInLock(func() {
		l.Level = LevelFatal
		l.Messages = append(l.Messages, msg)
		l.LastMessage = msg
		l.LastItems = fields
		l.WasFatal = true
	})
}

func (l *MockLogger) FatalWithContext(ctx context.Context, msg string, args ...interface{}) {
	msg, fields := buildMsgAndArgsWithContextMock(ctx, msg, args...)
	l.executeInLock(func() {
		l.Level = LevelFatal
		l.Messages = append(l.Messages, msg)
		l.LastMessage = msg
		l.LastItems = fields
		l.WasFatal = false
	})
}

func normalizeValuesMock(values ...interface{}) map[string]interface{} {
	fields := make(map[string]interface{})
	if len(values) == 1 {
		switch typedValue := values[0].(type) {
		case error:
			fields["error"] = typedValue.Error()
		case map[string]string:
			for k, v := range typedValue {
				fields[k] = v
			}
		case map[string]interface{}:
			for k, v := range typedValue {
				fields[k] = v
			}
		default:
			fields["extra"] = typedValue
		}
	} else if len(values) > 1 {
		fields["extra"] = values
	}

	return fields
}

func buildMsgAndArgsMock(msg string, args ...interface{}) (string, map[string]interface{}) {
	msg, args = formatMessageWithArgs(msg, args)
	return msg, normalizeValuesMock(args...)
}

func buildMsgAndArgsWithContextMock(ctx context.Context, msg string, args ...interface{}) (string, map[string]interface{}) {
	msg, fields := buildMsgAndArgsMock(msg, args...)
	span, has := tracer.SpanFromContext(ctx)
	if has {
		fields["dd.trace_id"] = span.Context().TraceID()
		fields["dd.span_id"] = span.Context().SpanID()
	}
	return msg, fields
}
