package log

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/mocktracer"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
	"testing"
)

func TestNewMock(t *testing.T) {
	logger := NewMock("service-name")
	assert.NotNil(t, logger)
}

func TestMockLogger_Debug(t *testing.T) {

	name := "service-name"
	logger := NewMock(name).(*MockLogger)

	testCases := []struct {
		name          string
		msg           string
		args          []interface{}
		shouldLevel   Level
		shouldLastMsg string
		shouldItems   map[string]interface{}
	}{
		{
			name:          "simple_msg",
			msg:           "just a simple msg",
			args:          nil,
			shouldLevel:   LevelDebug,
			shouldLastMsg: "just a simple msg",
			shouldItems:   map[string]interface{}{},
		},
		{
			name:          "formatted_msg",
			msg:           "just a simple msg %s",
			args:          []interface{}{"the replace"},
			shouldLevel:   LevelDebug,
			shouldLastMsg: "just a simple msg the replace",
			shouldItems:   map[string]interface{}{},
		},
		{
			name: "formatted_msg_with_key_value_arg",
			msg:  "just a simple msg %s",
			args: []interface{}{"the replace", map[string]interface{}{
				"key":  "value",
				"key2": "value2",
			}},
			shouldLevel:   LevelDebug,
			shouldLastMsg: "just a simple msg the replace",
			shouldItems: map[string]interface{}{
				"key":  "value",
				"key2": "value2",
			},
		},
		{
			name:          "formatted_msg_with_error_arg",
			msg:           "just a simple msg %s",
			args:          []interface{}{"the replace", errors.New("the new error")},
			shouldLevel:   LevelDebug,
			shouldLastMsg: "just a simple msg the replace",
			shouldItems: map[string]interface{}{
				"error": "the new error",
			},
		},
		{
			name:          "formatted_msg_with_string_arg",
			msg:           "just a simple msg %s",
			args:          []interface{}{"the replace", "some random string"},
			shouldLevel:   LevelDebug,
			shouldLastMsg: "just a simple msg the replace",
			shouldItems: map[string]interface{}{
				"extra": "some random string",
			},
		},
		{
			name:          "formatted_msg_with_number_arg",
			msg:           "just a simple msg %s",
			args:          []interface{}{"the replace", 15},
			shouldLevel:   LevelDebug,
			shouldLastMsg: "just a simple msg the replace",
			shouldItems: map[string]interface{}{
				"extra": 15,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			logger.Debug(tc.msg, tc.args...)
			assert.False(t, logger.WasFatal)
			assert.Equal(t, logger.Level, tc.shouldLevel)
			assert.Equal(t, logger.LastMessage, tc.shouldLastMsg)
			assert.Equal(t, logger.LastItems, tc.shouldItems)
		})
	}
}

func TestMockLogger_DebugWithContext(t *testing.T) {

	name := "service-name"
	logger := NewMock(name).(*MockLogger)

	testCases := []struct {
		name          string
		msg           string
		args          []interface{}
		shouldLevel   Level
		shouldLastMsg string
		shouldItems   map[string]interface{}
	}{
		{
			name:          "simple_msg",
			msg:           "just a simple msg",
			args:          nil,
			shouldLevel:   LevelDebug,
			shouldLastMsg: "just a simple msg",
			shouldItems:   map[string]interface{}{},
		},
		{
			name:          "formatted_msg",
			msg:           "just a simple msg %s",
			args:          []interface{}{"the replace"},
			shouldLevel:   LevelDebug,
			shouldLastMsg: "just a simple msg the replace",
			shouldItems:   map[string]interface{}{},
		},
		{
			name: "formatted_msg_with_key_value_arg",
			msg:  "just a simple msg %s",
			args: []interface{}{"the replace", map[string]interface{}{
				"key":  "value",
				"key2": "value2",
			}},
			shouldLevel:   LevelDebug,
			shouldLastMsg: "just a simple msg the replace",
			shouldItems: map[string]interface{}{
				"key":  "value",
				"key2": "value2",
			},
		},
		{
			name:          "formatted_msg_with_error_arg",
			msg:           "just a simple msg %s",
			args:          []interface{}{"the replace", errors.New("the new error")},
			shouldLevel:   LevelDebug,
			shouldLastMsg: "just a simple msg the replace",
			shouldItems: map[string]interface{}{
				"error": "the new error",
			},
		},
		{
			name:          "formatted_msg_with_string_arg",
			msg:           "just a simple msg %s",
			args:          []interface{}{"the replace", "some random string"},
			shouldLevel:   LevelDebug,
			shouldLastMsg: "just a simple msg the replace",
			shouldItems: map[string]interface{}{
				"extra": "some random string",
			},
		},
		{
			name:          "formatted_msg_with_number_arg",
			msg:           "just a simple msg %s",
			args:          []interface{}{"the replace", 15},
			shouldLevel:   LevelDebug,
			shouldLastMsg: "just a simple msg the replace",
			shouldItems: map[string]interface{}{
				"extra": 15,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mt := mocktracer.Start()
			defer mt.Stop()
			span := tracer.StartSpan("test")
			ctx := tracer.ContextWithSpan(context.Background(), span)
			logger.DebugWithContext(ctx, tc.msg, tc.args...)
			assert.False(t, logger.WasFatal)
			assert.Equal(t, logger.Level, tc.shouldLevel)
			assert.Equal(t, logger.LastMessage, tc.shouldLastMsg)
			assert.Equal(t, span.Context().SpanID(), logger.LastItems["dd.span_id"])
			assert.Equal(t, span.Context().TraceID(), logger.LastItems["dd.trace_id"])
			for k, v := range tc.shouldItems {
				assert.Equal(t, logger.LastItems[k], v)
			}
		})
	}
}

func TestMockLogger_Info(t *testing.T) {

	name := "service-name"
	logger := NewMock(name).(*MockLogger)

	testCases := []struct {
		name          string
		msg           string
		args          []interface{}
		shouldLevel   Level
		shouldLastMsg string
		shouldItems   map[string]interface{}
	}{
		{
			name:          "simple_msg",
			msg:           "just a simple msg",
			args:          nil,
			shouldLevel:   LevelInfo,
			shouldLastMsg: "just a simple msg",
			shouldItems:   map[string]interface{}{},
		},
		{
			name:          "formatted_msg",
			msg:           "just a simple msg %s",
			args:          []interface{}{"the replace"},
			shouldLevel:   LevelInfo,
			shouldLastMsg: "just a simple msg the replace",
			shouldItems:   map[string]interface{}{},
		},
		{
			name: "formatted_msg_with_key_value_arg",
			msg:  "just a simple msg %s",
			args: []interface{}{"the replace", map[string]interface{}{
				"key":  "value",
				"key2": "value2",
			}},
			shouldLevel:   LevelInfo,
			shouldLastMsg: "just a simple msg the replace",
			shouldItems: map[string]interface{}{
				"key":  "value",
				"key2": "value2",
			},
		},
		{
			name:          "formatted_msg_with_error_arg",
			msg:           "just a simple msg %s",
			args:          []interface{}{"the replace", errors.New("the new error")},
			shouldLevel:   LevelInfo,
			shouldLastMsg: "just a simple msg the replace",
			shouldItems: map[string]interface{}{
				"error": "the new error",
			},
		},
		{
			name:          "formatted_msg_with_string_arg",
			msg:           "just a simple msg %s",
			args:          []interface{}{"the replace", "some random string"},
			shouldLevel:   LevelInfo,
			shouldLastMsg: "just a simple msg the replace",
			shouldItems: map[string]interface{}{
				"extra": "some random string",
			},
		},
		{
			name:          "formatted_msg_with_number_arg",
			msg:           "just a simple msg %s",
			args:          []interface{}{"the replace", 15},
			shouldLevel:   LevelInfo,
			shouldLastMsg: "just a simple msg the replace",
			shouldItems: map[string]interface{}{
				"extra": 15,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			logger.Info(tc.msg, tc.args...)
			assert.False(t, logger.WasFatal)
			assert.Equal(t, logger.Level, tc.shouldLevel)
			assert.Equal(t, logger.LastMessage, tc.shouldLastMsg)
			assert.Equal(t, logger.LastItems, tc.shouldItems)
		})
	}
}

func TestMockLogger_InfoWithContext(t *testing.T) {

	name := "service-name"
	logger := NewMock(name).(*MockLogger)

	testCases := []struct {
		name          string
		msg           string
		args          []interface{}
		shouldLevel   Level
		shouldLastMsg string
		shouldItems   map[string]interface{}
	}{
		{
			name:          "simple_msg",
			msg:           "just a simple msg",
			args:          nil,
			shouldLevel:   LevelInfo,
			shouldLastMsg: "just a simple msg",
			shouldItems:   map[string]interface{}{},
		},
		{
			name:          "formatted_msg",
			msg:           "just a simple msg %s",
			args:          []interface{}{"the replace"},
			shouldLevel:   LevelInfo,
			shouldLastMsg: "just a simple msg the replace",
			shouldItems:   map[string]interface{}{},
		},
		{
			name: "formatted_msg_with_key_value_arg",
			msg:  "just a simple msg %s",
			args: []interface{}{"the replace", map[string]interface{}{
				"key":  "value",
				"key2": "value2",
			}},
			shouldLevel:   LevelInfo,
			shouldLastMsg: "just a simple msg the replace",
			shouldItems: map[string]interface{}{
				"key":  "value",
				"key2": "value2",
			},
		},
		{
			name:          "formatted_msg_with_error_arg",
			msg:           "just a simple msg %s",
			args:          []interface{}{"the replace", errors.New("the new error")},
			shouldLevel:   LevelInfo,
			shouldLastMsg: "just a simple msg the replace",
			shouldItems: map[string]interface{}{
				"error": "the new error",
			},
		},
		{
			name:          "formatted_msg_with_string_arg",
			msg:           "just a simple msg %s",
			args:          []interface{}{"the replace", "some random string"},
			shouldLevel:   LevelInfo,
			shouldLastMsg: "just a simple msg the replace",
			shouldItems: map[string]interface{}{
				"extra": "some random string",
			},
		},
		{
			name:          "formatted_msg_with_number_arg",
			msg:           "just a simple msg %s",
			args:          []interface{}{"the replace", 15},
			shouldLevel:   LevelInfo,
			shouldLastMsg: "just a simple msg the replace",
			shouldItems: map[string]interface{}{
				"extra": 15,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mt := mocktracer.Start()
			defer mt.Stop()
			span := tracer.StartSpan("test")
			ctx := tracer.ContextWithSpan(context.Background(), span)
			logger.InfoWithContext(ctx, tc.msg, tc.args...)
			assert.False(t, logger.WasFatal)
			assert.Equal(t, logger.Level, tc.shouldLevel)
			assert.Equal(t, logger.LastMessage, tc.shouldLastMsg)
			assert.Equal(t, span.Context().SpanID(), logger.LastItems["dd.span_id"])
			assert.Equal(t, span.Context().TraceID(), logger.LastItems["dd.trace_id"])
			for k, v := range tc.shouldItems {
				assert.Equal(t, logger.LastItems[k], v)
			}
		})
	}
}

func TestMockLogger_Error(t *testing.T) {

	name := "service-name"
	logger := NewMock(name).(*MockLogger)

	testCases := []struct {
		name          string
		msg           string
		args          []interface{}
		shouldLevel   Level
		shouldLastMsg string
		shouldItems   map[string]interface{}
	}{
		{
			name:          "simple_msg",
			msg:           "just a simple msg",
			args:          nil,
			shouldLevel:   LevelError,
			shouldLastMsg: "just a simple msg",
			shouldItems:   map[string]interface{}{},
		},
		{
			name:          "formatted_msg",
			msg:           "just a simple msg %s",
			args:          []interface{}{"the replace"},
			shouldLevel:   LevelError,
			shouldLastMsg: "just a simple msg the replace",
			shouldItems:   map[string]interface{}{},
		},
		{
			name: "formatted_msg_with_key_value_arg",
			msg:  "just a simple msg %s",
			args: []interface{}{"the replace", map[string]interface{}{
				"key":  "value",
				"key2": "value2",
			}},
			shouldLevel:   LevelError,
			shouldLastMsg: "just a simple msg the replace",
			shouldItems: map[string]interface{}{
				"key":  "value",
				"key2": "value2",
			},
		},
		{
			name:          "formatted_msg_with_error_arg",
			msg:           "just a simple msg %s",
			args:          []interface{}{"the replace", errors.New("the new error")},
			shouldLevel:   LevelError,
			shouldLastMsg: "just a simple msg the replace",
			shouldItems: map[string]interface{}{
				"error": "the new error",
			},
		},
		{
			name:          "formatted_msg_with_string_arg",
			msg:           "just a simple msg %s",
			args:          []interface{}{"the replace", "some random string"},
			shouldLevel:   LevelError,
			shouldLastMsg: "just a simple msg the replace",
			shouldItems: map[string]interface{}{
				"extra": "some random string",
			},
		},
		{
			name:          "formatted_msg_with_number_arg",
			msg:           "just a simple msg %s",
			args:          []interface{}{"the replace", 15},
			shouldLevel:   LevelError,
			shouldLastMsg: "just a simple msg the replace",
			shouldItems: map[string]interface{}{
				"extra": 15,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			logger.Error(tc.msg, tc.args...)
			assert.False(t, logger.WasFatal)
			assert.Equal(t, logger.Level, tc.shouldLevel)
			assert.Equal(t, logger.LastMessage, tc.shouldLastMsg)
			assert.Equal(t, logger.LastItems, tc.shouldItems)
		})
	}
}

func TestMockLogger_ErrorWithContext(t *testing.T) {

	name := "service-name"
	logger := NewMock(name).(*MockLogger)

	testCases := []struct {
		name          string
		msg           string
		args          []interface{}
		shouldLevel   Level
		shouldLastMsg string
		shouldItems   map[string]interface{}
	}{
		{
			name:          "simple_msg",
			msg:           "just a simple msg",
			args:          nil,
			shouldLevel:   LevelError,
			shouldLastMsg: "just a simple msg",
			shouldItems:   map[string]interface{}{},
		},
		{
			name:          "formatted_msg",
			msg:           "just a simple msg %s",
			args:          []interface{}{"the replace"},
			shouldLevel:   LevelError,
			shouldLastMsg: "just a simple msg the replace",
			shouldItems:   map[string]interface{}{},
		},
		{
			name: "formatted_msg_with_key_value_arg",
			msg:  "just a simple msg %s",
			args: []interface{}{"the replace", map[string]interface{}{
				"key":  "value",
				"key2": "value2",
			}},
			shouldLevel:   LevelError,
			shouldLastMsg: "just a simple msg the replace",
			shouldItems: map[string]interface{}{
				"key":  "value",
				"key2": "value2",
			},
		},
		{
			name:          "formatted_msg_with_error_arg",
			msg:           "just a simple msg %s",
			args:          []interface{}{"the replace", errors.New("the new error")},
			shouldLevel:   LevelError,
			shouldLastMsg: "just a simple msg the replace",
			shouldItems: map[string]interface{}{
				"error": "the new error",
			},
		},
		{
			name:          "formatted_msg_with_string_arg",
			msg:           "just a simple msg %s",
			args:          []interface{}{"the replace", "some random string"},
			shouldLevel:   LevelError,
			shouldLastMsg: "just a simple msg the replace",
			shouldItems: map[string]interface{}{
				"extra": "some random string",
			},
		},
		{
			name:          "formatted_msg_with_number_arg",
			msg:           "just a simple msg %s",
			args:          []interface{}{"the replace", 15},
			shouldLevel:   LevelError,
			shouldLastMsg: "just a simple msg the replace",
			shouldItems: map[string]interface{}{
				"extra": 15,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mt := mocktracer.Start()
			defer mt.Stop()
			span := tracer.StartSpan("test")
			ctx := tracer.ContextWithSpan(context.Background(), span)
			logger.ErrorWithContext(ctx, tc.msg, tc.args...)
			assert.False(t, logger.WasFatal)
			assert.Equal(t, logger.Level, tc.shouldLevel)
			assert.Equal(t, logger.LastMessage, tc.shouldLastMsg)
			assert.Equal(t, span.Context().SpanID(), logger.LastItems["dd.span_id"])
			assert.Equal(t, span.Context().TraceID(), logger.LastItems["dd.trace_id"])
			for k, v := range tc.shouldItems {
				assert.Equal(t, logger.LastItems[k], v)
			}
		})
	}
}
