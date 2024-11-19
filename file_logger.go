// -------------------------------------------
// @file      : file_logger.go
// @author    : bo cai
// @contact   : caibo923@gmail.com
// @time      : 2024/11/19 下午5:34
// -------------------------------------------

package logger

import (
	"github.com/natefinch/lumberjack"
)

// FileLogger 封装lumberjack.Logger，实现WriteSyncer接口
type FileLogger struct {
	*lumberjack.Logger
}

// Sync 实现WriteSyncer接口
func (fl *FileLogger) Sync() error {
	if fl.Logger != nil {
		return fl.Logger.Close()
	}
	return nil
}
