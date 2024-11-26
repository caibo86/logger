// -------------------------------------------
// @file      : redirect_unix.go
// @author    : bo cai
// @contact   : caibo923@gmail.com
// @time      : 2024/11/19 下午6:10
// -------------------------------------------

//go:build !windows

package logger

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

// 重定向标准错误输出到日志文件
func redirectStdErrLog() error {
	panicFile := strings.Replace(global.fileName, ".log", ".panic", -1)
	err := os.MkdirAll(filepath.Dir(panicFile), 0755)
	if err != nil {
		return err
	}
	fd, err := os.OpenFile(panicFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	err = syscall.Dup2(int(fd.Fd()), int(os.Stderr.Fd()))
	if err != nil {
		return err
	}
	Info("successfully redirected std err to panic log file")
	// 保活文件 避免删除
	go func() {
		checkStdErrLogFile()
		_, _ = fmt.Fprintln(os.Stderr, "no panic:"+time.Now().Format(time.RFC3339))
		fmt.Println("什么情况")
		hourTimer := time.NewTicker(1 * time.Hour)
		defer hourTimer.Stop()
		for {
			select {
			case <-hourTimer.C:
				checkStdErrLogFile()
				_, _ = fmt.Fprintln(os.Stderr, "no panic:"+time.Now().Format(time.RFC3339))
				hourTimer.Reset(1 * time.Hour)
			}
		}
	}()
	return nil
}

// 检查panic日志文件是否存在，并且把标准错误输出重定向到该文件
func checkStdErrLogFile() {
	panicFile := strings.Replace(global.fileName, ".log", ".panic", -1)
	_, err := os.Stat(panicFile)
	if !errors.Is(err, fs.ErrNotExist) {
		return
	}
	fd, err := os.OpenFile(panicFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return
	}
	err = syscall.Dup2(int(fd.Fd()), int(os.Stderr.Fd()))
	if err != nil {
		return
	}
}
