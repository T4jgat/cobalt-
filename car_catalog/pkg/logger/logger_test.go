package logger

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("valid log level", func(t *testing.T) {
		logger := New("info")

		assert.NotNil(t, logger)
		assert.NotNil(t, logger.logger)
	})

	t.Run("invalid log level", func(t *testing.T) {
		logger := New("invalid")

		assert.NotNil(t, logger)
		assert.NotNil(t, logger.logger)
	})
}
