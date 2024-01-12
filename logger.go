package nebula

import (
	"io"
	"net/http"
	"os"
	"time"
)

// Logger will act as a client to push logs to the nebula server
type Logger struct {
	host        string
	serviceName string
	client      *http.Client
	stdWriter   io.Writer
	httpWriter  io.Writer
}

func NewLogger(serviceName string, opts ...LoggerOption) *Logger {
	logger := &Logger{
		host:        "localhost:8100",
		serviceName: serviceName,
		client: &http.Client{
			Timeout: 15 * time.Second,
		},
		stdWriter:  os.Stdout,
		httpWriter: os.Stdout,
	}

	for _, opt := range opts {
		opt(logger)
	}

	return logger
}

func (l *Logger) Std() *StdLogger {
	return l.buildStdLogger(l.stdWriter)
}

func (l *Logger) Http() *HttpLogger {
	return &HttpLogger{
		StdLogger: l.buildStdLogger(l.httpWriter),
	}
}

func (l *Logger) buildStdLogger(w io.Writer) *StdLogger {
	return &StdLogger{
		nebulaHost:  l.host,
		serviceName: l.serviceName,
		client:      l.client,
		w:           w,
	}
}
