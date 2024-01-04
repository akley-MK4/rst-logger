package main

import (
	"bytes"
	"github.com/akley-MK4/rst-logger/define"
	"github.com/akley-MK4/rst-logger/implement"
	"os"
	"runtime"
	"strconv"
)

func newJsonFormatLogger() (*implement.BasicFormatLogger, error) {
	logLvDescMap := define.BuildDefaultLevelDescMap()
	logLvDescMap[define.LogLevelALL] = "trace"
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

	return implement.NewBasicFormatLogger(os.Stdout, define.LogLevelALL, 3, kw)
}
