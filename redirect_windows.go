// -------------------------------------------
// @file      : redirect_windows.go
// @author    : bo cai
// @contact   : caibo923@gmail.com
// @time      : 2024/11/19 下午6:10
// -------------------------------------------

//go:build windows

package logger

import (
	"golang.org/x/sys/windows"
	"os"
	"strings"
)

// 重定向标准错误输出到日志文件
func redirectStdErrLog() error {
	panicFile := strings.Replace(global.fileName, ".log", ".panic", -1)
	file, err := os.OpenFile(panicFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	if err = windows.SetStdHandle(windows.STD_ERROR_HANDLE, windows.Handle(file.Fd())); err != nil {
		return err
	}
	os.Stderr = file
	Debug("successfully redirected std err to panic log file")
	return nil
}

// 检查panic日志文件是否存在，并且把标准错误输出重定向到该文件
func checkStdErrLogFile() {
	panicFile := strings.Replace(global.fileName, ".log", ".panic", -1)
	_, err := os.Stat(panicFile)
	if !os.IsNotExist(err) {
		return
	}
	_ = redirectStdErrLog()
}
