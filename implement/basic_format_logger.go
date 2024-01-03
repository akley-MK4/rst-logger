package implement

import (
	"bytes"
	"errors"
	"github.com/akley-MK4/rst-logger/define"
	"io"
	"strconv"
)

type KwArgsBasicFormatLogger struct {
	EnableDefaultOpt        bool
	Prefix                  string
	TimeOpt                 define.OutputOption
	CallInfoOpt             define.OutputOption
	LogLevelOpt             define.OutputOption
	AdditionalBeforeFmtOpts []define.OutputOption
	AdditionalAfterFmtOpts  []define.OutputOption
	NewLogLvDescMap         map[define.LogLevel]string
	DisablePrinterLFChar    bool
}

func NewBasicFormatLogger(outWriter io.Writer, lv define.LogLevel, callDepth int, kw KwArgsBasicFormatLogger) (*BasicFormatLogger, error) {
	if outWriter == nil {
		return nil, errors.New("the outWriter is a nil value")
	}

	logLvDescMap := define.BuildDefaultLevelDescMap()
	for repLv, desc := range kw.NewLogLvDescMap {
		logLvDescMap[repLv] = desc
	}

	var newOpts []define.OutputOption
	if kw.Prefix != "" {
		prefixBytes := []byte(kw.Prefix)
		newOpts = append(newOpts, func(_ define.LogLevel, _ int, buf *bytes.Buffer) {
			buf.Write(prefixBytes)
			buf.WriteString(" ")
		})
	}

	if kw.TimeOpt != nil {
		newOpts = append(newOpts, kw.TimeOpt)
	} else if kw.EnableDefaultOpt {
		newOpts = append(newOpts, WithOptionDatetime)
	}

	if kw.CallInfoOpt != nil {
		newOpts = append(newOpts, kw.CallInfoOpt)
	} else if kw.EnableDefaultOpt {
		newOpts = append(newOpts, WithOptionCallInfo)
	}

	if kw.LogLevelOpt != nil {
		newOpts = append(newOpts, kw.LogLevelOpt)
	} else if kw.EnableDefaultOpt {
		newOpts = append(newOpts, func(lv define.LogLevel, _ int, buf *bytes.Buffer) {
			lvDesc := logLvDescMap[lv]
			if lvDesc == "" {
				lvDesc = strconv.Itoa(int(lv))
			}
			buf.WriteString("[")
			buf.WriteString(lvDesc)
			buf.WriteString("] ")
		})
	}

	beforeFmtOpts := append(newOpts, kw.AdditionalBeforeFmtOpts...)
	afterFmtOpts := append([]define.OutputOption{}, kw.AdditionalAfterFmtOpts...)

	printer, printerErr := define.NewPrinter(outWriter, lv, callDepth, beforeFmtOpts, afterFmtOpts, kw.DisablePrinterLFChar)
	if printerErr != nil {
		return nil, printerErr
	}

	retLogger := &BasicFormatLogger{
		printer:      printer,
		logLvDescMap: logLvDescMap,
	}

	return retLogger, nil
}

type BasicFormatLogger struct {
	printer      *define.Printer
	logLvDescMap map[define.LogLevel]string
}

func (t *BasicFormatLogger) SetLevelByDesc(levelDesc string) (retUpdated bool) {
	var chooseLv define.LogLevel
	for lv, desc := range t.logLvDescMap {
		if desc == levelDesc {
			retUpdated = true
			chooseLv = lv
			break
		}
	}

	if !retUpdated {
		return
	}

	t.printer.SetLogLevel(chooseLv)
	return
}

func (t *BasicFormatLogger) All(v ...any) {
	t.printer.OutputFormatContent(define.LogLevelALL, "", v...)
}

func (t *BasicFormatLogger) AllF(format string, v ...any) {
	t.printer.OutputFormatContent(define.LogLevelALL, format, v...)
}

func (t *BasicFormatLogger) Debug(v ...any) {
	t.printer.OutputFormatContent(define.LogLevelDebug, "", v...)
}

func (t *BasicFormatLogger) DebugF(format string, v ...any) {
	t.printer.OutputFormatContent(define.LogLevelDebug, format, v...)
}

func (t *BasicFormatLogger) Info(v ...any) {
	t.printer.OutputFormatContent(define.LogLevelInfo, "", v...)
}

func (t *BasicFormatLogger) InfoF(format string, v ...any) {
	t.printer.OutputFormatContent(define.LogLevelInfo, format, v...)
}

func (t *BasicFormatLogger) Warning(v ...any) {
	t.printer.OutputFormatContent(define.LogLevelWarning, "", v...)
}

func (t *BasicFormatLogger) WarningF(format string, v ...any) {
	t.printer.OutputFormatContent(define.LogLevelWarning, format, v...)
}

func (t *BasicFormatLogger) Error(v ...any) {
	t.printer.OutputFormatContent(define.LogLevelError, "", v...)
}

func (t *BasicFormatLogger) ErrorF(format string, v ...any) {
	t.printer.OutputFormatContent(define.LogLevelError, format, v...)
}
