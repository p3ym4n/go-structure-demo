package log

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormatMessageWithArgs(t *testing.T) {
	testCases := []struct {
		name    string
		msgIn   string
		argsIn  []interface{}
		msgOut  string
		argsOut []interface{}
	}{
		{
			name:    "simple_message",
			msgIn:   "this is the simple message",
			argsIn:  nil,
			msgOut:  "this is the simple message",
			argsOut: make([]interface{}, 0),
		},
		{
			name:    "simple_message_with_0_threshold_1_arg",
			msgIn:   "this is the simple message",
			argsIn:  []interface{}{"hard"},
			msgOut:  "this is the simple message",
			argsOut: []interface{}{"hard"},
		},
		{
			name:    "simple_message_with_1_threshold_1_arg",
			msgIn:   "this is the %s simple message",
			argsIn:  []interface{}{"hard"},
			msgOut:  "this is the hard simple message",
			argsOut: make([]interface{}, 0),
		},
		{
			name:    "simple_message_with_1_threshold_2_arg",
			msgIn:   "this is the %s simple message",
			argsIn:  []interface{}{"hard", "soft"},
			msgOut:  "this is the hard simple message",
			argsOut: []interface{}{"soft"},
		},
		{
			name:    "simple_message_with_2_threshold_1_arg",
			msgIn:   "this is the %s simple message %d",
			argsIn:  []interface{}{"hard"},
			msgOut:  "this is the hard simple message %!d(MISSING)",
			argsOut: make([]interface{}, 0),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			outMsg, outArgs := formatMessageWithArgs(tc.msgIn, tc.argsIn)
			assert.Equal(t, tc.msgOut, outMsg)
			assert.Equal(t, tc.argsOut, outArgs)
		})
	}
}
