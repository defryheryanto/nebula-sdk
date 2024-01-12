# Go Nebula SDK
Go SDK for [Nebula Dashboard](https://github.com/defryheryanto/nebula). Import and use for easier integration between your app and Nebula.

## Instruction
### Setting Up Logger
Install `go get github.com/defryheryanto/nebula-sdk` in your application.<br><br>
Set up a new logger instance
```
logger := nebula.NewLogger("your-service-name")
```
To set up a new logger that will point to a specific host to push the log.<br>
```
logger := nebula.NewLogger("your-service-name", nebula.LoggerHostOption("https://your-nebula-host"))
```
If `nebula.LoggerHostOption` is not provided, it will point to the default host which is `http://localhost:8100`

### Std Logging
Std Logging is a type of logs in nebula, which is all regular logs should use this type of log.<br>
To push log to the std log:
```
logger := nebula.NewLogger("your-service-name")
nebula.SetLogger(logger)
nebula.StdLog().Info("Hello World!", map[string]any{
  "additional_info": "This will push the log to std-log type",
})
```
This code will output the following log:
```
{"level":"INFO","message":"Hello World!","additional_info":"This will push the log to std-log type"}
```
There are 3 level of the log, `Info`, `Warning`, and `Error`. `Info` and `Warning` have the same effect, the only difference is the level. For `Error` level, you can specify the `error` instance and it will be included in the log automatically.
To push an error log:
```
logger := nebula.NewLogger("your-service-name")
nebula.SetLogger(logger)
nebula.Stdlog().Error("error test", errors.New("this is a test error"), map[string]any{
  "additional_info": "This will be push to the std log with error level",
})
```
This code will output the following log:
```
{"level":"ERROR", "message":"error test","error":"this is a test error","additional_info":"This will be push to the std log with error level"}
```

### Http Logging
Http Logging will provide functionalities to log HTTP Information
To set request to be included in the log:
```
httpLogger := nebula.HttpLog()
httpLogger.SetRequest(req)
```
To set response to be included in the log:
```
httpLogger := nebula.HttpLog()
httpLogger.SetResponse(resp)
```
To push the log to nebula:
```
httpLogger.Info()
```
Same with the Std() log, there are 3 level of the log, you can use it by doing `httpLogger.Warning()` or `httpLogger.Error(err)`

Complete Http Logging example:
```
package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/defryheryanto/nebula-sdk"
)

func main() {
	httpLogger := nebula.HttpLog()

	type testPayload struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	p := &testPayload{
		Username: "admin",
		Password: "admin123",
	}
	b, _ := json.Marshal(p)
	req, _ := http.NewRequest(http.MethodPost, "http://localhost:8100/api/v1/users/auth", bytes.NewReader(b))
	req.Header.Add("Authorization", "Token your-token-here")

	httpLogger.SetRequest(req)

	type testResponsePayload struct {
		Token string `json:"token"`
	}
	rp := &testResponsePayload{
		Token: "some token",
	}
	b, _ = json.Marshal(rp)

	resp := &http.Response{}
	resp.Body = io.NopCloser(bytes.NewReader(b))

	httpLogger.SetResponse(resp)

	httpLogger.Info()
}
```
This code will output the following log:
```
{"endpoint":"/api/v1/users/auth","headers":{"Authorization":["Token your-token-here"]},"host":"localhost:8100","level":"INFO","method":"POST","requestBody":"{\"username\":\"admin\",\"password\":\"admin123\"}","responseBody":"{\"token\":\"some token\"}"}
```

