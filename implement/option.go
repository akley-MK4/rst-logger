package implement

import (
	"bytes"
	"github.com/akley-MK4/rst-logger/define"
	"runtime"
	"strconv"
	"time"
)

const (
	defaultTimeOutputFmt = "2006-01-02 15:04:05"
)

func WithOptionDatetime(_ define.LogLevel, _ int, buf *bytes.Buffer) {
	nowTime := time.Now()
	buf.WriteString(nowTime.Format(defaultTimeOutputFmt))
	buf.WriteString(" ")
}

func WithOptionCallInfo(_ define.LogLevel, callDepth int, buf *bytes.Buffer) {
	_, file, line, ok := runtime.Caller(callDepth)
	if !ok {
		file = "???"
		line = 0
	}
	buf.WriteString(file)
	buf.WriteString(":")
	buf.WriteString(strconv.Itoa(line))
	buf.WriteString(" ")
}
