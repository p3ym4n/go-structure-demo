package log

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestConfigFromEnv(t *testing.T) {

	testCases := []struct {
		name      string
		envLevel  string
		envFormat string
		outLevel  Level
		outFormat Format
	}{
		{
			name:      "defaults",
			envLevel:  "",
			envFormat: "",
			outLevel:  LevelInfo,
			outFormat: FormatJSON,
		},
		{
			name:      "recognized_case_insensitive",
			envLevel:  "ErRor",
			envFormat: "CoNsoLe",
			outLevel:  LevelError,
			outFormat: FormatConsole,
		},
		{
			name:      "non_recognized_case_insensitive",
			envLevel:  "not-correct",
			envFormat: "non-related",
			outLevel:  LevelInfo,
			outFormat: FormatJSON,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_ = os.Setenv("LOG_LEVEL", tc.envLevel)
			_ = os.Setenv("LOG_FORMAT", tc.envFormat)
			format, level := configFromEnv()
			assert.Equal(t, tc.outFormat, format)
			assert.Equal(t, tc.outLevel, level)
			_ = os.Setenv("LOG_LEVEL", "")
			_ = os.Setenv("LOG_FORMAT", "")
		})
	}
}
