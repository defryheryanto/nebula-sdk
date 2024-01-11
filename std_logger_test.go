package nebula

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStdLogger(t *testing.T) {
	t.Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	t.Run("Failed server", func(t *testing.T) {
		failServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))

		logger := &StdLogger{
			nebulaHost:  failServer.URL,
			serviceName: "nebula",
			client:      &http.Client{},
		}

		err := logger.Info("hehe", nil)
		assert.Error(t, err)
	})

	t.Run("Info", func(t *testing.T) {
		logger := &StdLogger{
			nebulaHost:  server.URL,
			serviceName: "nebula",
			client:      &http.Client{},
		}

		err := logger.Info("hehe", nil)
		assert.NoError(t, err)
	})

	t.Run("Warning", func(t *testing.T) {
		logger := &StdLogger{
			nebulaHost:  server.URL,
			serviceName: "nebula",
			client:      &http.Client{},
		}

		err := logger.Warning("hehe", map[string]any{})
		assert.NoError(t, err)
	})

	t.Run("Error", func(t *testing.T) {
		logger := &StdLogger{
			nebulaHost:  server.URL,
			serviceName: "nebula",
			client:      &http.Client{},
		}

		err := logger.Error("hehe", fmt.Errorf("test"), map[string]any{
			"id":       1,
			"username": "hehe",
			"session": map[string]any{
				"id":    "hehe",
				"token": "hoho",
			},
		})
		assert.NoError(t, err)
	})
}
