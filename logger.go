// -------------------------------------------
// @file      : logger.go
// @author    : bo cai
// @contact   : caibo923@gmail.com
// @time      : 2024/11/19 上午10:18
// -------------------------------------------

package logger

import (
	"go.uber.org/zap"
	"sync"
)

// 全局日志管理器
var global = &Logger{}

// 全局只能初始化一次
var once sync.Once

// Logger 日志封装
type Logger struct {
	zapLogger *zap.SugaredLogger
	atom      zap.AtomicLevel
	fileName  string
	options   *Options
}

// Close 关闭日志
func (logger *Logger) Close() error {
	if logger.zapLogger == nil {
		return nil
	}
	return logger.zapLogger.Sync()
}

// Init 日志初始化
func (logger *Logger) Init(option ...Option) {
	options := DefaultOptions
	for _, o := range option {
		o(&options)
	}
	logger.options = &options
	logger.fileName = options.Filename
	logger.atom = zap.NewAtomicLevel()
	logger.atom.SetLevel(options.Level)
	logger.zapLogger = options.GetZapLogger()
}
