package main

import (
	"github.com/akley-MK4/rst-logger/define"
	"github.com/akley-MK4/rst-logger/implement"
	"os"
)

func newBasicFormatLogger() (*implement.BasicFormatLogger, error) {
	logLvDescMap := define.BuildDefaultLevelDescMap()
	//logLvDescMap[define.LogLevelALL] = "trace"
	kw := implement.KwArgsBasicFormatLogger{
		EnableDefaultOpt:     true,
		Prefix:               "[App]",
		DisablePrinterLFChar: false,
		NewLogLvDescMap:      logLvDescMap,
	}

	return implement.NewBasicFormatLogger(os.Stdout, define.LogLevelALL, 3, kw)
}
