// -------------------------------------------
// @file      : main.go
// @author    : bo cai
// @contact   : caibo923@gmail.com
// @time      : 2024/11/19 下午5:57
// -------------------------------------------

package main

import (
	"github.com/caibo86/logger"
)

func main() {
	logger.Init(
		logger.SetLevel(logger.DebugLevel),
		logger.SetCallerSkip(2),
		logger.SetFilename("test.log"),
		logger.SetIsRedirectErr(false),
		logger.SetIsOpenFile(true),
		// logger.SetFormatType(logger.LogFormatJson),
		logger.SetStacktrace(logger.ErrorLevel),
	)
	defer func() {
		_ = logger.Close()
	}()
	logger.Debug("hello", "world", 2024)
	logger.Debugf("hello world %v", 2004)
	logger.Debugw("hello world", "foo", "bar")
	logger.Error("we have an error")
	logger.Fatal("hello", "world", 2024)
}
