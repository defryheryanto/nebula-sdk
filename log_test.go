package nebula

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLog(t *testing.T) {
	t.Parallel()

	t.Run("without message", func(t *testing.T) {
		l := &log{
			level:   infoLogLevel,
			message: "",
			data: map[string]any{
				"id":       1,
				"username": "hehe",
			},
		}

		res := l.getData()
		assert.Equal(t, map[string]any{
			"level":    l.level,
			"id":       1,
			"username": "hehe",
		}, res)
	})

	t.Run("without error", func(t *testing.T) {
		l := &log{
			level:   infoLogLevel,
			message: "hello world!",
			data: map[string]any{
				"id":       1,
				"username": "hehe",
			},
		}

		res := l.getData()
		assert.Equal(t, map[string]any{
			"level":    l.level,
			"message":  l.message,
			"id":       1,
			"username": "hehe",
		}, res)
	})

	t.Run("with error", func(t *testing.T) {
		l := &log{
			level:   errorLogLevel,
			message: "hello world!",
			err:     fmt.Errorf("test"),
			data: map[string]any{
				"id":       1,
				"username": "hehe",
			},
		}

		res := l.getData()
		assert.Equal(t, map[string]any{
			"level":    l.level,
			"message":  l.message,
			"error":    l.err.Error(),
			"id":       1,
			"username": "hehe",
		}, res)
	})
}
