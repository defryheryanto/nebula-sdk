package nebula_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/defryheryanto/nebula-sdk"
	"github.com/stretchr/testify/assert"
)

func TestLogHttpMiddleware(t *testing.T) {
	t.Parallel()

	realHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World!"))
	})
	middleware := nebula.LogHttpMiddleware(realHandler)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/test", nil)
	tt := httptest.NewRecorder()
	middleware.ServeHTTP(tt, req)

	assert.Equal(t, http.StatusOK, tt.Result().StatusCode)

	b, err := io.ReadAll(tt.Result().Body)
	assert.NoError(t, err)

	assert.Equal(t, "Hello World!", string(b))
}
