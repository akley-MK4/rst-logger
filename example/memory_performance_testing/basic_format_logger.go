package main

import (
	"os"

	"github.com/akley-MK4/rst-logger/define"
	"github.com/akley-MK4/rst-logger/implement"
)

func newBasicFormatLogger() (*implement.BasicFormatLogger, error) {
	logLvDescMap := define.BuildDefaultLevelDescMap()
	//logLvDescMap[define.LogLevelALL] = "trace"
	kw := implement.KwArgsBasicFormatLogger{
		EnableDefaultOpt:     true,
		Prefix:               "[App]",
		DisablePrinterLFChar: false,
		NewLogLvDescMap:      logLvDescMap,
		EnabledBufferPool:    true,
	}

	return implement.NewBasicFormatLogger(os.Stdout, define.LogLevelALL, 3, kw)
}
