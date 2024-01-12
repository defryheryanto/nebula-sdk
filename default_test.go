package nebula_test

import (
	"testing"

	"github.com/defryheryanto/nebula-sdk"
	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	t.Parallel()

	logger := nebula.NewLogger("nebula")
	nebula.SetLogger(logger)

	t.Run("Std Logger", func(t *testing.T) {
		assert.Equal(t, logger.Std(), nebula.StdLog())
	})

	t.Run("Http Logger", func(t *testing.T) {
		assert.Equal(t, logger.Http(), nebula.HttpLog())
	})
}
