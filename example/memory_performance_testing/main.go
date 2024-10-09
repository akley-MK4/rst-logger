package main

import (
	"log"
	"os"
	"runtime"
	"sync"
	"time"
)

func main() {
	basicFormatLogger, basicFormatLoggerErr := newBasicFormatLogger()
	if basicFormatLoggerErr != nil {
		log.Println("Failed to create BasicFormatLogger, ", basicFormatLoggerErr.Error())
		os.Exit(1)
	}

	jsonFormatLogger, jsonFormatLoggerErr := newJsonFormatLogger()
	if jsonFormatLoggerErr != nil {
		log.Println("Failed to create JsonFormatLogger, ", jsonFormatLoggerErr.Error())
		os.Exit(1)
	}

	initialMemInfo := collectMemoryStatus("Initial memory")

	maxCoNum := 10
	maxOutputExecWithOneCo := 1000
	var wg sync.WaitGroup
	wg.Add(maxCoNum)

	begTime := time.Now()
	for i := 0; i < maxCoNum; i++ {
		go func(inIdx int, inWg *sync.WaitGroup) {
			defer inWg.Done()

			for j := 0; j < maxOutputExecWithOneCo; j++ {
				jsonFormatLogger.All("jsonFormatLogger Log Test 111111111111111111111111111111111111111111111111111")
				jsonFormatLogger.AllF("jsonFormatLogger Log Test 111111111111111111111111111111111111111111111111, "+
					"%s, %d, %d", "output...", inIdx, j)
				jsonFormatLogger.Debug("jsonFormatLogger Log Test 111111111111111111111111111111111111111111111111111")
				jsonFormatLogger.DebugF("jsonFormatLogger Log Test 111111111111111111111111111111111111111111111111, "+
					"%s, %d, %d", "output...", inIdx, j)
				jsonFormatLogger.Info("jsonFormatLogger Log Test 111111111111111111111111111111111111111111111111111")
				jsonFormatLogger.InfoF("jsonFormatLogger Log Test 111111111111111111111111111111111111111111111111, "+
					"%s, %d, %d", "output...", inIdx, j)
				jsonFormatLogger.Warning("jsonFormatLogger Log Test 111111111111111111111111111111111111111111111111111")
				jsonFormatLogger.WarningF("jsonFormatLogger Log Test 111111111111111111111111111111111111111111111111, "+
					"%s, %d, %d", "output...", inIdx, j)
				jsonFormatLogger.Error("jsonFormatLogger Log Test 111111111111111111111111111111111111111111111111111")
				jsonFormatLogger.ErrorF("jsonFormatLogger Log Test 111111111111111111111111111111111111111111111111, "+
					"%s, %d, %d", "output...", inIdx, j)

				basicFormatLogger.All("BasicFormatLogger Log Test 11111111111111111111111111111111111111111111111111")
				basicFormatLogger.AllF("BasicFormatLogger Log Test 11111111111111111111111111111111111111111111111, "+
					"%s, %d, %d", "output...", inIdx, j)
				basicFormatLogger.Debug("BasicFormatLogger Log Test 11111111111111111111111111111111111111111111111111")
				basicFormatLogger.DebugF("BasicFormatLogger Log Test 11111111111111111111111111111111111111111111111, "+
					"%s, %d, %d", "output...", inIdx, j)
				basicFormatLogger.Info("BasicFormatLogger Log Test 11111111111111111111111111111111111111111111111111")
				basicFormatLogger.InfoF("BasicFormatLogger Log Test 11111111111111111111111111111111111111111111111, "+
					"%s, %d, %d", "output...", inIdx, j)
				basicFormatLogger.Warning("BasicFormatLogger Log Test 11111111111111111111111111111111111111111111111111")
				basicFormatLogger.WarningF("BasicFormatLogger Log Test 11111111111111111111111111111111111111111111111, "+
					"%s, %d, %d", "output...", inIdx, j)
				basicFormatLogger.Error("BasicFormatLogger Log Test 11111111111111111111111111111111111111111111111111")
				basicFormatLogger.ErrorF("BasicFormatLogger Log Test 11111111111111111111111111111111111111111111111, "+
					"%s, %d, %d", "output...", inIdx, j)
			}
		}(i, &wg)
	}

	wg.Wait()

	endTime := time.Now()
	log.Printf("Accumulated output of %d logs, consumed %d seconds %d milliseconds",
		maxCoNum*maxOutputExecWithOneCo*20, endTime.Unix()-begTime.Unix(), endTime.UnixMilli()-begTime.UnixMilli())

	collectMemoryStatus("Current memory")

	if !jsonFormatLogger.SetLevelByDesc("INFO") {
		log.Println("Setting JsonFormatLogger log level failed")
	} else {
		jsonFormatLogger.All("jsonFormatLogger Log Test 2222222222222222222222222")
		jsonFormatLogger.AllF("jsonFormatLogger Log Test 222222222222222222, "+
			"%s, %d, %d", "output...", 0, 0)
		jsonFormatLogger.Debug("jsonFormatLogger Log Test 22222222222222222")
		jsonFormatLogger.DebugF("jsonFormatLogger Log Test 22222222222222222222222, "+
			"%s, %d, %d", "output...", 0, 0)
	}
	if !basicFormatLogger.SetLevelByDesc("INFO") {
		log.Println("Setting BasicFormatLogger log level failed")
	} else {
		basicFormatLogger.All("BasicFormatLogger Log Test 2222222222222222222222222")
		basicFormatLogger.AllF("BasicFormatLogger Log Test 22222222222222222222222222, "+
			"%s, %d, %d", "output...", 0, 0)
		basicFormatLogger.Debug("BasicFormatLogger Log Test 222222222222222222222222222222")
		basicFormatLogger.DebugF("BasicFormatLogger Log Test 22222222222222222222222, "+
			"%s, %d, %d", "output...", 0, 0)
	}

	log.Printf("Clean basicFormatLogger capacitySize %d\n", basicFormatLogger.CleanOutputBuffer())
	log.Printf("Clean jsonFormatLogger capacitySize %d\n", jsonFormatLogger.CleanOutputBuffer())

	time.Sleep(time.Second * 1)
	runtime.GC()
	time.Sleep(time.Second * 1)
	runtime.GC()
	time.Sleep(time.Second * 5)

	finalMemInfo := collectMemoryStatus("Final memory")
	printMemoryStatus(initialMemInfo)

	if (finalMemInfo.AllocMBs - initialMemInfo.AllocMBs) >= 1 {
		log.Println("There is a difference between the final memory and the initial memory size, please check if there is a memory leak")
	} else {
		log.Println("Successfully ran the example without generating any errors")
	}

	os.Exit(0)
}
