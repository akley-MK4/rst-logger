package define

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"sync"
)

type OutputOption func(lv LogLevel, callDepth int, buf *bytes.Buffer)

var (
	bufferPool = sync.Pool{
		New: func() any {
			return &bytes.Buffer{}
		},
	}
)

func NewPrinter(outWriter io.Writer, lv LogLevel, callDepth int, beforeFmtOpts, afterFmtOpts []OutputOption, disableLFChar, enabledBufferPool bool) (*Printer, error) {
	if outWriter == nil {
		return nil, errors.New("the outWriter is a nil value")
	}

	return &Printer{
		outWriter:         outWriter,
		lv:                lv,
		callDepth:         callDepth,
		beforeFmtOpts:     beforeFmtOpts,
		afterFmtOpts:      afterFmtOpts,
		disabledLFChar:    disableLFChar,
		outputBuffer:      &bytes.Buffer{},
		enabledBufferPool: enabledBufferPool,
	}, nil
}

type Printer struct {
	lv                LogLevel
	callDepth         int
	beforeFmtOpts     []OutputOption
	afterFmtOpts      []OutputOption
	outWriter         io.Writer
	disabledLFChar    bool
	outputMutex       sync.Mutex
	outputBuffer      *bytes.Buffer
	enabledBufferPool bool
}

func (t *Printer) OutputFormatContent(lv LogLevel, format string, v ...any) (retErr error) {
	if lv < t.lv {
		return
	}
	if len(v) <= 0 {
		return errors.New("no output values")
	}

	var buf *bytes.Buffer
	if t.enabledBufferPool {
		buf = bufferPool.Get().(*bytes.Buffer)
	} else {
		t.outputMutex.Lock()
		buf = t.outputBuffer
	}

	defer func() {
		buf.Reset()
		if t.enabledBufferPool {
			bufferPool.Put(buf)
		} else {
			t.outputMutex.Unlock()
		}
	}()

	for _, opt := range t.beforeFmtOpts {
		opt(lv, t.callDepth, buf)
	}

	if format != "" {
		if _, err := fmt.Fprintf(buf, format, v...); err != nil {
			retErr = err
			return
		}
	} else {
		if _, err := fmt.Fprint(buf, v...); err != nil {
			retErr = err
			return
		}
	}
	if !t.disabledLFChar {
		buf.WriteString("\n")
	}

	for _, opt := range t.afterFmtOpts {
		opt(lv, t.callDepth, buf)
	}

	_, writeErr := t.outWriter.Write(buf.Bytes())
	retErr = writeErr
	return
}

func (t *Printer) OutputBytesContent(lv LogLevel, v ...[]byte) (retErr error) {
	if lv < t.lv {
		return
	}
	if len(v) <= 0 {
		return errors.New("no output values")
	}

	var buf *bytes.Buffer
	if t.enabledBufferPool {
		buf = bufferPool.Get().(*bytes.Buffer)
	} else {
		t.outputMutex.Lock()
		buf = t.outputBuffer
	}

	defer func() {
		buf.Reset()
		if t.enabledBufferPool {
			bufferPool.Put(buf)
		} else {
			t.outputMutex.Unlock()
		}
	}()

	for _, opt := range t.beforeFmtOpts {
		opt(lv, t.callDepth, buf)
	}

	var capSize int
	for _, d := range v {
		capSize += len(d)
	}
	if buf.Cap() < (capSize + LfCharsLen) {
		buf.Grow(capSize + LfCharsLen)
	}

	for _, d := range v {
		buf.Write(d)
	}

	for _, opt := range t.afterFmtOpts {
		opt(lv, t.callDepth, buf)
	}

	_, writeErr := t.outWriter.Write(buf.Bytes())
	retErr = writeErr
	return
}

func (t *Printer) SetLogLevel(lv LogLevel) {
	t.lv = lv
}

func (t *Printer) GetLevel() LogLevel {
	return t.lv
}

func (t *Printer) CleanOutputBuffer() (capacitySize int) {
	if t.enabledBufferPool {
		return 0
	}

	t.outputMutex.Lock()
	t.outputBuffer.Reset()
	capacitySize = t.outputBuffer.Cap()
	t.outputBuffer = &bytes.Buffer{}
	t.outputMutex.Unlock()

	return
}
