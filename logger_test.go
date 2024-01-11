package nebula_test

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/defryheryanto/nebula-sdk"
	"github.com/stretchr/testify/assert"
)

func TestLogger_StdLogger(t *testing.T) {
	t.Parallel()

	t.Run("Fail Server", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("hello world"))
		}))

		logger := nebula.NewLogger("nebula", nebula.LoggerHostOption(server.URL))
		err := logger.Std().Info("test", nil)
		assert.Equal(t, "http error 500: hello world", err.Error())
	})

	t.Run("Success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		logger := nebula.NewLogger("nebula", nebula.LoggerHostOption(server.URL))
		err := logger.Std().Info("test", nil)
		assert.NoError(t, err)
	})
}

func TestLogger_HttpLogger(t *testing.T) {
	t.Parallel()

	t.Run("Fail Server", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("hello world"))
		}))

		logger := nebula.NewLogger("nebula", nebula.LoggerHostOption(server.URL))
		httpLogger := logger.Http()

		req, err := http.NewRequest(http.MethodPost, "localhost/api", bytes.NewReader([]byte("request body")))
		assert.NoError(t, err)

		err = httpLogger.SetRequest(req)
		assert.NoError(t, err)

		resp := &http.Response{}
		resp.Body = io.NopCloser(bytes.NewReader([]byte("response")))

		err = httpLogger.SetResponse(resp)
		assert.NoError(t, err)

		err = httpLogger.Info()
		assert.Equal(t, "http error 500: hello world", err.Error())
	})

	t.Run("Success", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
		}))

		logger := nebula.NewLogger("nebula", nebula.LoggerHostOption(server.URL))
		httpLogger := logger.Http()

		req, err := http.NewRequest(http.MethodPost, "localhost/api", bytes.NewReader([]byte("request body")))
		assert.NoError(t, err)

		err = httpLogger.SetRequest(req)
		assert.NoError(t, err)

		resp := &http.Response{}
		resp.Body = io.NopCloser(bytes.NewReader([]byte("response")))

		err = httpLogger.SetResponse(resp)
		assert.NoError(t, err)

		err = httpLogger.Info()
		assert.NoError(t, err)
	})
}
