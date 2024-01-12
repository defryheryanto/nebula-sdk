package nebula

import (
	"net/http"
	"net/http/httptest"
)

type httpResponseRecorder struct {
	http.ResponseWriter
	recorder *httptest.ResponseRecorder
}

func (w *httpResponseRecorder) Header() http.Header {
	return w.ResponseWriter.Header()
}

func (w *httpResponseRecorder) Write(b []byte) (int, error) {
	w.recorder.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w *httpResponseRecorder) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.recorder.WriteHeader(statusCode)
}

func (w *httpResponseRecorder) Response() *http.Response {
	w.recorder.Result().Header = w.ResponseWriter.Header()
	return w.recorder.Result()
}

// LogHttpMiddleware will push logs as an info level http log
func LogHttpMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		responseRecorder := &httpResponseRecorder{
			ResponseWriter: w,
			recorder:       httptest.NewRecorder(),
		}

		httpLogger := HttpLog()
		httpLogger.SetRequest(r)

		next.ServeHTTP(responseRecorder, r)
		httpLogger.SetResponse(responseRecorder.Response())
		httpLogger.Info()
	})
}
