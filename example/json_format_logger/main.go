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
