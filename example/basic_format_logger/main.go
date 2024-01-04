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
