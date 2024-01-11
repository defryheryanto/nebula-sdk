package nebula

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHttpLogger(t *testing.T) {
	t.Parallel()
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	t.Run("Failed server", func(t *testing.T) {
		failServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))

		logger := &HttpLogger{
			StdLogger: &StdLogger{
				nebulaHost:  failServer.URL,
				serviceName: "nebula",
				client:      &http.Client{},
			},
		}

		req, err := http.NewRequest(http.MethodPost, "localhost/api", bytes.NewReader([]byte("request body")))
		assert.NoError(t, err)

		err = logger.SetRequest(req)
		assert.NoError(t, err)

		resp := &http.Response{}
		resp.Body = io.NopCloser(bytes.NewReader([]byte("response")))

		err = logger.SetResponse(resp)
		assert.NoError(t, err)

		err = logger.Error(fmt.Errorf("test"))
		assert.Error(t, err)
	})

	t.Run("Without request", func(t *testing.T) {
		logger := &HttpLogger{
			StdLogger: &StdLogger{
				nebulaHost:  server.URL,
				serviceName: "nebula",
				client:      &http.Client{},
			},
		}

		resp := &http.Response{}
		resp.Body = io.NopCloser(bytes.NewReader([]byte("response")))

		err := logger.SetResponse(resp)
		assert.NoError(t, err)

		err = logger.Info()
		assert.NoError(t, err)
	})

	t.Run("Without response", func(t *testing.T) {
		logger := &HttpLogger{
			StdLogger: &StdLogger{
				nebulaHost:  server.URL,
				serviceName: "nebula",
				client:      &http.Client{},
			},
		}

		req, err := http.NewRequest(http.MethodPost, "localhost/api", bytes.NewReader([]byte("request body")))
		assert.NoError(t, err)

		err = logger.SetRequest(req)
		assert.NoError(t, err)

		err = logger.Warning()
		assert.NoError(t, err)
	})

	t.Run("With Request & Response", func(t *testing.T) {
		logger := &HttpLogger{
			StdLogger: &StdLogger{
				nebulaHost:  server.URL,
				serviceName: "nebula",
				client:      &http.Client{},
			},
		}

		req, err := http.NewRequest(http.MethodPost, "localhost/api", bytes.NewReader([]byte("request body")))
		assert.NoError(t, err)

		err = logger.SetRequest(req)
		assert.NoError(t, err)

		resp := &http.Response{}
		resp.Body = io.NopCloser(bytes.NewReader([]byte("response")))

		err = logger.SetResponse(resp)
		assert.NoError(t, err)

		err = logger.Error(fmt.Errorf("test"))
		assert.NoError(t, err)
	})
}
