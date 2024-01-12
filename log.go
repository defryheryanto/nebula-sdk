package nebula

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
)

type LogLevel string

const (
	infoLogLevel    LogLevel = "INFO"
	warningLogLevel LogLevel = "WARNING"
	errorLogLevel   LogLevel = "ERROR"

	stdLog  = "std-log"
	httpLog = "http-log"
)

type log struct {
	level   LogLevel
	message string
	err     error
	data    map[string]any
}

func (l *log) getData() map[string]any {
	out := map[string]any{
		"level": l.level,
	}
	if l.message != "" {
		out["message"] = l.message
	}
	if l.err != nil {
		out["error"] = l.err.Error()
	}

	if l.data != nil {
		for key, val := range l.data {
			if key != "level" && key != "message" && key != "error" {
				out[key] = val
			}
		}
	}

	return out
}

type pushLogRequestBody struct {
	ServiceName string         `json:"service_name"`
	LogType     string         `json:"log_type"`
	Log         map[string]any `json:"log"`
}

func printLog(w io.Writer, logBody *pushLogRequestBody) {
	output := logBody.Log
	output["service_name"] = logBody.ServiceName
	output["log_type"] = logBody.LogType

	b, _ := json.Marshal(output)

	fmt.Fprintf(w, "%s\n", string(b))
}

func push(w io.Writer, host string, client *http.Client, serviceName string, logType string, data *log) error {
	url := fmt.Sprintf("%s/api/logs", host)

	body := &pushLogRequestBody{
		ServiceName: serviceName,
		LogType:     logType,
		Log:         data.getData(),
	}
	b, err := json.Marshal(body)
	if err != nil {
		return err
	}

	printLog(w, body)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(b))
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode >= http.StatusMultipleChoices {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return errors.Wrap(err, "error reading body")
		}

		return errors.New(fmt.Sprintf("http error %d: %s", resp.StatusCode, string(b)))
	}

	return nil
}
