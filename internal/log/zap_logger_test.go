package log

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zapcore"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/mocktracer"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"os"
	"testing"
)

func TestNewZapFromEnv(t *testing.T) {
	logger, closer := NewZapFromEnv("some names")
	assert.IsType(t, func() {}, closer)
	assert.NotNil(t, logger)
}

func TestNewZap(t *testing.T) {
	format := FormatConsole
	level := LevelError
	logger, closer := NewZap("some names", format, level, zapcore.Lock(os.Stdout))
	assert.IsType(t, func() {}, closer)
	assert.NotNil(t, logger)
}

func TestZapLogger_Debug(t *testing.T) {

	out := new(bytes.Buffer)
	writer := zapcore.AddSync(out)
	name := "service-name"
	logger, syncer := NewZap(name, FormatJSON, LevelDebug, writer)
	defer syncer()

	testCases := []struct {
		name      string
		msg       string
		args      []interface{}
		shouldHas map[string]interface{}
	}{
		{
			name: "simple_msg",
			msg:  "just a simple msg",
			args: nil,
			shouldHas: map[string]interface{}{
				"level": "debug",
				"msg":   "just a simple msg",
				"name":  name,
			},
		},
		{
			name: "formatted_msg",
			msg:  "just a simple msg %s",
			args: []interface{}{"the replace"},
			shouldHas: map[string]interface{}{
				"level": "debug",
				"msg":   "just a simple msg the replace",
				"name":  name,
			},
		},
		{
			name: "formatted_msg_with_key_value_arg",
			msg:  "just a simple msg %s",
			args: []interface{}{"the replace", map[string]interface{}{
				"key":  "value",
				"key2": "value2",
			}},
			shouldHas: map[string]interface{}{
				"level": "debug",
				"msg":   "just a simple msg the replace",
				"name":  name,
				"key":   "value",
				"key2":  "value2",
			},
		},
		{
			name: "formatted_msg_with_error_arg",
			msg:  "just a simple msg %s",
			args: []interface{}{"the replace", errors.New("the new error")},
			shouldHas: map[string]interface{}{
				"level": "debug",
				"msg":   "just a simple msg the replace",
				"name":  name,
				"error": "the new error",
			},
		},
		{
			name: "formatted_msg_with_string_arg",
			msg:  "just a simple msg %s",
			args: []interface{}{"the replace", "some random string"},
			shouldHas: map[string]interface{}{
				"level": "debug",
				"msg":   "just a simple msg the replace",
				"name":  name,
				"extra": "some random string",
			},
		},
		{
			name: "formatted_msg_with_number_arg",
			msg:  "just a simple msg %s",
			args: []interface{}{"the replace", 15},
			shouldHas: map[string]interface{}{
				"level": "debug",
				"msg":   "just a simple msg the replace",
				"name":  name,
				"extra": float64(15),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out.Reset()
			logger.Debug(tc.msg, tc.args...)
			items := make(map[string]interface{})
			err := json.Unmarshal(out.Bytes(), &items)
			assert.Nil(t, err)
			for k, v := range tc.shouldHas {
				res, has := items[k]
				assert.True(t, has)
				assert.Equal(t, v, res)
			}
		})
	}
}

func TestZapLogger_DebugWithContext(t *testing.T) {

	out := new(bytes.Buffer)
	writer := zapcore.AddSync(out)
	name := "service-name"
	logger, syncer := NewZap(name, FormatJSON, LevelDebug, writer)
	defer syncer()

	testCases := []struct {
		name      string
		msg       string
		args      []interface{}
		shouldHas map[string]interface{}
	}{
		{
			name: "simple_msg",
			msg:  "just a simple msg",
			args: nil,
			shouldHas: map[string]interface{}{
				"level": "debug",
				"msg":   "just a simple msg",
				"name":  name,
			},
		},
		{
			name: "formatted_msg",
			msg:  "just a simple msg %s",
			args: []interface{}{"the replace"},
			shouldHas: map[string]interface{}{
				"level": "debug",
				"msg":   "just a simple msg the replace",
				"name":  name,
			},
		},
		{
			name: "formatted_msg_with_key_value_arg",
			msg:  "just a simple msg %s",
			args: []interface{}{"the replace", map[string]interface{}{
				"key":  "value",
				"key2": "value2",
			}},
			shouldHas: map[string]interface{}{
				"level": "debug",
				"msg":   "just a simple msg the replace",
				"name":  name,
				"key":   "value",
				"key2":  "value2",
			},
		},
		{
			name: "formatted_msg_with_error_arg",
			msg:  "just a simple msg %s",
			args: []interface{}{"the replace", errors.New("the new error")},
			shouldHas: map[string]interface{}{
				"level": "debug",
				"msg":   "just a simple msg the replace",
				"name":  name,
				"error": "the new error",
			},
		},
		{
			name: "formatted_msg_with_string_arg",
			msg:  "just a simple msg %s",
			args: []interface{}{"the replace", "some random string"},
			shouldHas: map[string]interface{}{
				"level": "debug",
				"msg":   "just a simple msg the replace",
				"name":  name,
				"extra": "some random string",
			},
		},
		{
			name: "formatted_msg_with_number_arg",
			msg:  "just a simple msg %s",
			args: []interface{}{"the replace", 15},
			shouldHas: map[string]interface{}{
				"level": "debug",
				"msg":   "just a simple msg the replace",
				"name":  name,
				"extra": float64(15),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out.Reset()
			mt := mocktracer.Start()
			defer mt.Stop()
			span := tracer.StartSpan("test")
			ctx := tracer.ContextWithSpan(context.Background(), span)
			logger.DebugWithContext(ctx, tc.msg, tc.args...)
			items := make(map[string]interface{})
			err := json.Unmarshal(out.Bytes(), &items)
			assert.Nil(t, err)
			for k, v := range tc.shouldHas {
				res, has := items[k]
				assert.True(t, has)
				assert.Equal(t, v, res)
			}
			assert.Equal(t, float64(span.Context().SpanID()), items["dd.span_id"])
			assert.Equal(t, float64(span.Context().TraceID()), items["dd.trace_id"])
		})
	}
}

func TestZapLogger_Info(t *testing.T) {

	out := new(bytes.Buffer)
	writer := zapcore.AddSync(out)
	name := "service-name"
	logger, syncer := NewZap(name, FormatJSON, LevelDebug, writer)
	defer syncer()

	testCases := []struct {
		name      string
		msg       string
		args      []interface{}
		shouldHas map[string]interface{}
	}{
		{
			name: "simple_msg",
			msg:  "just a simple msg",
			args: nil,
			shouldHas: map[string]interface{}{
				"level": "info",
				"msg":   "just a simple msg",
				"name":  name,
			},
		},
		{
			name: "formatted_msg",
			msg:  "just a simple msg %s",
			args: []interface{}{"the replace"},
			shouldHas: map[string]interface{}{
				"level": "info",
				"msg":   "just a simple msg the replace",
				"name":  name,
			},
		},
		{
			name: "formatted_msg_with_key_value_arg",
			msg:  "just a simple msg %s",
			args: []interface{}{"the replace", map[string]interface{}{
				"key":  "value",
				"key2": "value2",
			}},
			shouldHas: map[string]interface{}{
				"level": "info",
				"msg":   "just a simple msg the replace",
				"name":  name,
				"key":   "value",
				"key2":  "value2",
			},
		},
		{
			name: "formatted_msg_with_error_arg",
			msg:  "just a simple msg %s",
			args: []interface{}{"the replace", errors.New("the new error")},
			shouldHas: map[string]interface{}{
				"level": "info",
				"msg":   "just a simple msg the replace",
				"name":  name,
				"error": "the new error",
			},
		},
		{
			name: "formatted_msg_with_string_arg",
			msg:  "just a simple msg %s",
			args: []interface{}{"the replace", "some random string"},
			shouldHas: map[string]interface{}{
				"level": "info",
				"msg":   "just a simple msg the replace",
				"name":  name,
				"extra": "some random string",
			},
		},
		{
			name: "formatted_msg_with_number_arg",
			msg:  "just a simple msg %s",
			args: []interface{}{"the replace", 15},
			shouldHas: map[string]interface{}{
				"level": "info",
				"msg":   "just a simple msg the replace",
				"name":  name,
				"extra": float64(15),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out.Reset()
			logger.Info(tc.msg, tc.args...)
			items := make(map[string]interface{})
			err := json.Unmarshal(out.Bytes(), &items)
			assert.Nil(t, err)
			for k, v := range tc.shouldHas {
				res, has := items[k]
				assert.True(t, has)
				assert.Equal(t, v, res)
			}
		})
	}
}

func TestZapLogger_InfoWithContext(t *testing.T) {

	out := new(bytes.Buffer)
	writer := zapcore.AddSync(out)
	name := "service-name"
	logger, syncer := NewZap(name, FormatJSON, LevelDebug, writer)
	defer syncer()

	testCases := []struct {
		name      string
		msg       string
		args      []interface{}
		shouldHas map[string]interface{}
	}{
		{
			name: "simple_msg",
			msg:  "just a simple msg",
			args: nil,
			shouldHas: map[string]interface{}{
				"level": "info",
				"msg":   "just a simple msg",
				"name":  name,
			},
		},
		{
			name: "formatted_msg",
			msg:  "just a simple msg %s",
			args: []interface{}{"the replace"},
			shouldHas: map[string]interface{}{
				"level": "info",
				"msg":   "just a simple msg the replace",
				"name":  name,
			},
		},
		{
			name: "formatted_msg_with_key_value_arg",
			msg:  "just a simple msg %s",
			args: []interface{}{"the replace", map[string]interface{}{
				"key":  "value",
				"key2": "value2",
			}},
			shouldHas: map[string]interface{}{
				"level": "info",
				"msg":   "just a simple msg the replace",
				"name":  name,
				"key":   "value",
				"key2":  "value2",
			},
		},
		{
			name: "formatted_msg_with_error_arg",
			msg:  "just a simple msg %s",
			args: []interface{}{"the replace", errors.New("the new error")},
			shouldHas: map[string]interface{}{
				"level": "info",
				"msg":   "just a simple msg the replace",
				"name":  name,
				"error": "the new error",
			},
		},
		{
			name: "formatted_msg_with_string_arg",
			msg:  "just a simple msg %s",
			args: []interface{}{"the replace", "some random string"},
			shouldHas: map[string]interface{}{
				"level": "info",
				"msg":   "just a simple msg the replace",
				"name":  name,
				"extra": "some random string",
			},
		},
		{
			name: "formatted_msg_with_number_arg",
			msg:  "just a simple msg %s",
			args: []interface{}{"the replace", 15},
			shouldHas: map[string]interface{}{
				"level": "info",
				"msg":   "just a simple msg the replace",
				"name":  name,
				"extra": float64(15),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out.Reset()
			mt := mocktracer.Start()
			defer mt.Stop()
			span := tracer.StartSpan("test")
			ctx := tracer.ContextWithSpan(context.Background(), span)
			logger.InfoWithContext(ctx, tc.msg, tc.args...)
			items := make(map[string]interface{})
			err := json.Unmarshal(out.Bytes(), &items)
			assert.Nil(t, err)
			for k, v := range tc.shouldHas {
				res, has := items[k]
				assert.True(t, has)
				assert.Equal(t, v, res)
			}
			assert.Equal(t, float64(span.Context().SpanID()), items["dd.span_id"])
			assert.Equal(t, float64(span.Context().TraceID()), items["dd.trace_id"])
		})
	}
}

func TestZapLogger_Error(t *testing.T) {

	out := new(bytes.Buffer)
	writer := zapcore.AddSync(out)
	name := "service-name"
	logger, syncer := NewZap(name, FormatJSON, LevelDebug, writer)
	defer syncer()

	testCases := []struct {
		name      string
		msg       string
		args      []interface{}
		shouldHas map[string]interface{}
	}{
		{
			name: "simple_msg",
			msg:  "just a simple msg",
			args: nil,
			shouldHas: map[string]interface{}{
				"level": "error",
				"msg":   "just a simple msg",
				"name":  name,
			},
		},
		{
			name: "formatted_msg",
			msg:  "just a simple msg %s",
			args: []interface{}{"the replace"},
			shouldHas: map[string]interface{}{
				"level": "error",
				"msg":   "just a simple msg the replace",
				"name":  name,
			},
		},
		{
			name: "formatted_msg_with_key_value_arg",
			msg:  "just a simple msg %s",
			args: []interface{}{"the replace", map[string]interface{}{
				"key":  "value",
				"key2": "value2",
			}},
			shouldHas: map[string]interface{}{
				"level": "error",
				"msg":   "just a simple msg the replace",
				"name":  name,
				"key":   "value",
				"key2":  "value2",
			},
		},
		{
			name: "formatted_msg_with_error_arg",
			msg:  "just a simple msg %s",
			args: []interface{}{"the replace", errors.New("the new error")},
			shouldHas: map[string]interface{}{
				"level": "error",
				"msg":   "just a simple msg the replace",
				"name":  name,
				"error": "the new error",
			},
		},
		{
			name: "formatted_msg_with_string_arg",
			msg:  "just a simple msg %s",
			args: []interface{}{"the replace", "some random string"},
			shouldHas: map[string]interface{}{
				"level": "error",
				"msg":   "just a simple msg the replace",
				"name":  name,
				"extra": "some random string",
			},
		},
		{
			name: "formatted_msg_with_number_arg",
			msg:  "just a simple msg %s",
			args: []interface{}{"the replace", 15},
			shouldHas: map[string]interface{}{
				"level": "error",
				"msg":   "just a simple msg the replace",
				"name":  name,
				"extra": float64(15),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out.Reset()
			logger.Error(tc.msg, tc.args...)
			items := make(map[string]interface{})
			err := json.Unmarshal(out.Bytes(), &items)
			assert.Nil(t, err)
			for k, v := range tc.shouldHas {
				res, has := items[k]
				assert.True(t, has)
				assert.Equal(t, v, res)
			}
		})
	}
}

func TestZapLogger_ErrorWithContext(t *testing.T) {

	out := new(bytes.Buffer)
	writer := zapcore.AddSync(out)
	name := "service-name"
	logger, syncer := NewZap(name, FormatJSON, LevelDebug, writer)
	defer syncer()

	testCases := []struct {
		name      string
		msg       string
		args      []interface{}
		shouldHas map[string]interface{}
	}{
		{
			name: "simple_msg",
			msg:  "just a simple msg",
			args: nil,
			shouldHas: map[string]interface{}{
				"level": "error",
				"msg":   "just a simple msg",
				"name":  name,
			},
		},
		{
			name: "formatted_msg",
			msg:  "just a simple msg %s",
			args: []interface{}{"the replace"},
			shouldHas: map[string]interface{}{
				"level": "error",
				"msg":   "just a simple msg the replace",
				"name":  name,
			},
		},
		{
			name: "formatted_msg_with_key_value_arg",
			msg:  "just a simple msg %s",
			args: []interface{}{"the replace", map[string]interface{}{
				"key":  "value",
				"key2": "value2",
			}},
			shouldHas: map[string]interface{}{
				"level": "error",
				"msg":   "just a simple msg the replace",
				"name":  name,
				"key":   "value",
				"key2":  "value2",
			},
		},
		{
			name: "formatted_msg_with_error_arg",
			msg:  "just a simple msg %s",
			args: []interface{}{"the replace", errors.New("the new error")},
			shouldHas: map[string]interface{}{
				"level": "error",
				"msg":   "just a simple msg the replace",
				"name":  name,
				"error": "the new error",
			},
		},
		{
			name: "formatted_msg_with_string_arg",
			msg:  "just a simple msg %s",
			args: []interface{}{"the replace", "some random string"},
			shouldHas: map[string]interface{}{
				"level": "error",
				"msg":   "just a simple msg the replace",
				"name":  name,
				"extra": "some random string",
			},
		},
		{
			name: "formatted_msg_with_number_arg",
			msg:  "just a simple msg %s",
			args: []interface{}{"the replace", 15},
			shouldHas: map[string]interface{}{
				"level": "error",
				"msg":   "just a simple msg the replace",
				"name":  name,
				"extra": float64(15),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			out.Reset()
			mt := mocktracer.Start()
			defer mt.Stop()
			span := tracer.StartSpan("test")
			ctx := tracer.ContextWithSpan(context.Background(), span)
			logger.ErrorWithContext(ctx, tc.msg, tc.args...)
			items := make(map[string]interface{})
			err := json.Unmarshal(out.Bytes(), &items)
			assert.Nil(t, err)
			for k, v := range tc.shouldHas {
				res, has := items[k]
				assert.True(t, has)
				assert.Equal(t, v, res)
			}
			assert.Equal(t, float64(span.Context().SpanID()), items["dd.span_id"])
			assert.Equal(t, float64(span.Context().TraceID()), items["dd.trace_id"])
		})
	}
}
