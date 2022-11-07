package log

import (
	"context"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	stdlog "log"
	"os"
)

var _ Logger = (*ZapLogger)(nil)

type ZapLogger struct {
	internal *zap.Logger
	level    Level
	format   Format
}

func NewZapFromEnv(name string) (Logger, func()) {
	format, level := configFromEnv()
	return NewZap(name, format, level, zapcore.Lock(os.Stdout))
}

func NewZap(name string, format Format, level Level, writer zapcore.WriteSyncer) (Logger, func()) {

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	encoderConfig.NameKey = "name"
	encoder := zapcore.NewJSONEncoder(encoderConfig)
	if format == FormatConsole {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	var coreLevel zap.AtomicLevel
	switch level {
	case LevelDebug:
		coreLevel = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	case LevelInfo:
		coreLevel = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	case LevelError:
		coreLevel = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	default:
		coreLevel = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	}

	logger := zap.New(zapcore.NewCore(encoder, writer, coreLevel))
	logger = logger.Named(name)
	zap.ReplaceGlobals(logger)
	_ = zap.RedirectStdLog(logger)

	return &ZapLogger{
			internal: logger,
			level:    level,
			format:   format,
		}, func() {
			_ = logger.Sync()
		}
}

func (z *ZapLogger) GetStd() *stdlog.Logger {
	return zap.NewStdLog(z.internal)
}

func (z *ZapLogger) GetLevel() Level {
	return z.level
}

func (z *ZapLogger) GetFormat() Format {
	return z.format
}

func (z *ZapLogger) Debug(msg string, args ...interface{}) {
	msg, fields := buildMsgAndArgsZap(msg, args...)
	z.internal.Debug(msg, fields...)
}

func (z *ZapLogger) DebugWithContext(ctx context.Context, msg string, args ...interface{}) {
	msg, fields := buildMsgAndArgsWithContextZap(ctx, msg, args...)
	z.internal.Debug(msg, fields...)
}

func (z *ZapLogger) Info(msg string, args ...interface{}) {
	msg, fields := buildMsgAndArgsZap(msg, args...)
	z.internal.Info(msg, fields...)
}

func (z *ZapLogger) InfoWithContext(ctx context.Context, msg string, args ...interface{}) {
	msg, fields := buildMsgAndArgsWithContextZap(ctx, msg, args...)
	z.internal.Info(msg, fields...)
}

func (z *ZapLogger) Error(msg string, args ...interface{}) {
	msg, fields := buildMsgAndArgsZap(msg, args...)
	z.internal.Error(msg, fields...)
}

func (z *ZapLogger) ErrorWithContext(ctx context.Context, msg string, args ...interface{}) {
	msg, fields := buildMsgAndArgsWithContextZap(ctx, msg, args...)
	z.internal.Error(msg, fields...)
}

func (z *ZapLogger) Fatal(msg string, args ...interface{}) {
	msg, fields := buildMsgAndArgsZap(msg, args...)
	z.internal.Fatal(msg, fields...)
}

func (z *ZapLogger) FatalWithContext(ctx context.Context, msg string, args ...interface{}) {
	msg, fields := buildMsgAndArgsWithContextZap(ctx, msg, args...)
	z.internal.Fatal(msg, fields...)
}

func normalizeValuesZap(values ...interface{}) []zap.Field {
	fields := make([]zap.Field, 0, len(values))
	if len(values) == 1 {
		switch typedValue := values[0].(type) {
		case error:
			fields = append(fields, zap.Error(typedValue))
		case map[string]string:
			for k, v := range typedValue {
				fields = append(fields, zap.String(k, v))
			}
		case map[string]interface{}:
			for k, v := range typedValue {
				fields = append(fields, zap.Any(k, v))
			}
		default:
			fields = append(fields, zap.Any("extra", typedValue))
		}
	} else if len(values) > 1 {
		fields = append(fields, zap.Any("extra", values))
	}

	return fields
}

func buildMsgAndArgsZap(msg string, args ...interface{}) (string, []zap.Field) {
	msg, args = formatMessageWithArgs(msg, args)
	var fields []zap.Field
	fields = append(fields, normalizeValuesZap(args...)...)
	return msg, fields
}

func buildMsgAndArgsWithContextZap(ctx context.Context, msg string, args ...interface{}) (string, []zap.Field) {
	msg, fields := buildMsgAndArgsZap(msg, args...)
	span, has := tracer.SpanFromContext(ctx)
	if has {
		fields = append(
			fields,
			zap.Uint64("dd.trace_id", span.Context().TraceID()),
			zap.Uint64("dd.span_id", span.Context().SpanID()),
		)
	}
	return msg, fields
}
