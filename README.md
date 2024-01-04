# rst-logger
## Introduction
This is a simple log library that can output the log format you want through string concatenation operations. Meanwhile, this library reuses memory for outputting logs.

## Example of log output
This is a basic example of formatting logs
```go
package main

import (
	"github.com/akley-MK4/rst-logger/define"
	"github.com/akley-MK4/rst-logger/implement"
	"log"
	"os"
)

func main() {
	kw := implement.KwArgsBasicFormatLogger{
		EnableDefaultOpt:     true,
		Prefix:               "[App]",
		DisablePrinterLFChar: false,
	}

	callDepth := 3
	logger, loggerErr := implement.NewBasicFormatLogger(os.Stdout, define.LogLevelALL, callDepth, kw)
	if loggerErr != nil {
		log.Println("Failed to create a logger, ", loggerErr.Error())
		os.Exit(1)
	}

	logger.All("Using the All function to output logs with level ALL")
	logger.AllF("Using the AllF function to output logs with level ALL, "+
		"%s, %d, %v", "output...", 1, true)
	logger.Debug("Using the Debug function to output logs with level DEBUG")
	logger.DebugF("Using the DebugF function to output logs with level DEBUG, "+
		"%s, %d, %v", "output...", 1, true)
	logger.Info("Using the Info function to output logs with level INFO")
	logger.InfoF("Using the InfoF function to output logs with level INFO, "+
		"%s, %d, %v", "output...", 1, true)
	logger.Warning("Using the Warning function to output logs with level WARNING")
	logger.WarningF("Using the WarningF function to output logs with level WARNING, "+
		"%s, %d, %v", "output...", 1, true)
	logger.Error("Using the Error function to output logs with level ERROR")
	logger.ErrorF("Using the ErrorF function to output logs with level ERROR, "+
		"%s, %d, %v", "output...", 1, true)

	os.Exit(0)
}
```
The above example will output the following logs
```text
[App] 2024-01-04 01:53:47 /home/op1/go_project_local/rst-logger/example/basic_format_logger/main.go:24 [ALL] Using the All function to output logs with level ALL
[App] 2024-01-04 01:53:47 /home/op1/go_project_local/rst-logger/example/basic_format_logger/main.go:25 [ALL] Using the AllF function to output logs with level ALL, output..., 1, true
[App] 2024-01-04 01:53:47 /home/op1/go_project_local/rst-logger/example/basic_format_logger/main.go:27 [DEBUG] Using the Debug function to output logs with level DEBUG
[App] 2024-01-04 01:53:47 /home/op1/go_project_local/rst-logger/example/basic_format_logger/main.go:28 [DEBUG] Using the DebugF function to output logs with level DEBUG, output..., 1, true
[App] 2024-01-04 01:53:47 /home/op1/go_project_local/rst-logger/example/basic_format_logger/main.go:30 [INFO] Using the Info function to output logs with level INFO
[App] 2024-01-04 01:53:47 /home/op1/go_project_local/rst-logger/example/basic_format_logger/main.go:31 [INFO] Using the InfoF function to output logs with level INFO, output..., 1, true
[App] 2024-01-04 01:53:47 /home/op1/go_project_local/rst-logger/example/basic_format_logger/main.go:33 [WARNING] Using the Warning function to output logs with level WARNING
[App] 2024-01-04 01:53:47 /home/op1/go_project_local/rst-logger/example/basic_format_logger/main.go:34 [WARNING] Using the WarningF function to output logs with level WARNING, output..., 1, true
[App] 2024-01-04 01:53:47 /home/op1/go_project_local/rst-logger/example/basic_format_logger/main.go:36 [ERROR] Using the Error function to output logs with level ERROR
[App] 2024-01-04 01:53:47 /home/op1/go_project_local/rst-logger/example/basic_format_logger/main.go:37 [ERROR] Using the ErrorF function to output logs with level ERROR, output..., 1, true
```

// The following example outputs a log in Json format. In the future, template based Json format log output will be supported.
```go
package main

import (
	"bytes"
	"github.com/akley-MK4/rst-logger/define"
	"github.com/akley-MK4/rst-logger/implement"
	"log"
	"os"
	"runtime"
	"strconv"
)

func main() {

	logLvDescMap := define.BuildDefaultLevelDescMap()
	// Use custom log level descriptions
	logLvDescMap[define.LogLevelALL] = "Tracing"
	logLvDescMap[define.LogLevelDebug] = "Debug"
	logLvDescMap[define.LogLevelInfo] = "Info"
	logLvDescMap[define.LogLevelWarning] = "Warning"
	logLvDescMap[define.LogLevelError] = "Error"
	kw := implement.KwArgsBasicFormatLogger{
		EnableDefaultOpt:     false,
		DisablePrinterLFChar: true,
		NewLogLvDescMap:      logLvDescMap,
	}

	kw.AdditionalBeforeFmtOpts = append(kw.AdditionalBeforeFmtOpts, func(lv define.LogLevel, callDepth int, buf *bytes.Buffer) {
		// level
		buf.WriteString(`{"level":"`)
		lvDesc := logLvDescMap[lv]
		if lvDesc == "" {
			lvDesc = strconv.Itoa(int(lv))
		}
		buf.WriteString(lvDesc)
		buf.WriteString(`",`)

		// call info
		var funcName string
		pc, file, line, ok := runtime.Caller(callDepth)
		if !ok {
			file = "???"
			line = 0
		} else {
			f := runtime.FuncForPC(pc)
			if f != nil {
				funcName = f.Name()
			}
		}

		buf.WriteString(`"profile":{"src":{"file":"`)
		buf.WriteString(file)
		buf.WriteString(`",`)
		buf.WriteString(`"line":`)
		buf.WriteString(strconv.Itoa(line))
		buf.WriteString(`,`)
		buf.WriteString(`"function":"`)
		buf.WriteString(funcName)
		buf.WriteString(`"}}`)

		// message
		buf.WriteString(`,"message":"`)
	})

	kw.AdditionalAfterFmtOpts = append(kw.AdditionalAfterFmtOpts, func(_ define.LogLevel, _ int, buf *bytes.Buffer) {
		buf.WriteString(`"}`)
		buf.WriteString("\n")
	})

	logger, loggerErr := implement.NewBasicFormatLogger(os.Stdout, define.LogLevelALL, 3, kw)
	if loggerErr != nil {
		log.Println("Failed to create a logger, ", loggerErr.Error())
		os.Exit(1)
	}

	logger.All("Using the All function to output logs with level ALL")
	logger.AllF("Using the AllF function to output logs with level ALL, "+
		"%s, %d, %v", "output...", 1, true)
	logger.Debug("Using the Debug function to output logs with level DEBUG")
	logger.DebugF("Using the DebugF function to output logs with level DEBUG, "+
		"%s, %d, %v", "output...", 1, true)
	logger.Info("Using the Info function to output logs with level INFO")
	logger.InfoF("Using the InfoF function to output logs with level INFO, "+
		"%s, %d, %v", "output...", 1, true)
	logger.Warning("Using the Warning function to output logs with level WARNING")
	logger.WarningF("Using the WarningF function to output logs with level WARNING, "+
		"%s, %d, %v", "output...", 1, true)
	logger.Error("Using the Error function to output logs with level ERROR")
	logger.ErrorF("Using the ErrorF function to output logs with level ERROR, "+
		"%s, %d, %v", "output...", 1, true)

	os.Exit(0)
}
```
```
The above example will output the following logs
```text
{"level":"Tracing","profile":{"src":{"file":"/home/op1/go_project_local/rst-logger/example/json_format_logger/main.go","line":76,"function":"main.main"}},"message":"Using the All function to output logs with level ALL"}
{"level":"Tracing","profile":{"src":{"file":"/home/op1/go_project_local/rst-logger/example/json_format_logger/main.go","line":77,"function":"main.main"}},"message":"Using the AllF function to output logs with level ALL, output..., 1, true"}
{"level":"Debug","profile":{"src":{"file":"/home/op1/go_project_local/rst-logger/example/json_format_logger/main.go","line":79,"function":"main.main"}},"message":"Using the Debug function to output logs with level DEBUG"}
{"level":"Debug","profile":{"src":{"file":"/home/op1/go_project_local/rst-logger/example/json_format_logger/main.go","line":80,"function":"main.main"}},"message":"Using the DebugF function to output logs with level DEBUG, output..., 1, true"}
{"level":"Info","profile":{"src":{"file":"/home/op1/go_project_local/rst-logger/example/json_format_logger/main.go","line":82,"function":"main.main"}},"message":"Using the Info function to output logs with level INFO"}
{"level":"Info","profile":{"src":{"file":"/home/op1/go_project_local/rst-logger/example/json_format_logger/main.go","line":83,"function":"main.main"}},"message":"Using the InfoF function to output logs with level INFO, output..., 1, true"}
{"level":"Warning","profile":{"src":{"file":"/home/op1/go_project_local/rst-logger/example/json_format_logger/main.go","line":85,"function":"main.main"}},"message":"Using the Warning function to output logs with level WARNING"}
{"level":"Warning","profile":{"src":{"file":"/home/op1/go_project_local/rst-logger/example/json_format_logger/main.go","line":86,"function":"main.main"}},"message":"Using the WarningF function to output logs with level WARNING, output..., 1, true"}
{"level":"Error","profile":{"src":{"file":"/home/op1/go_project_local/rst-logger/example/json_format_logger/main.go","line":88,"function":"main.main"}},"message":"Using the Error function to output logs with level ERROR"}
{"level":"Error","profile":{"src":{"file":"/home/op1/go_project_local/rst-logger/example/json_format_logger/main.go","line":89,"function":"main.main"}},"message":"Using the ErrorF function to output logs with level ERROR, output..., 1, true"}
```