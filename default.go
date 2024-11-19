// -------------------------------------------
// @file      : default.go
// @author    : bo cai
// @contact   : caibo923@gmail.com
// @time      : 2024/11/19 上午10:16
// -------------------------------------------

package logger

import "go.uber.org/zap"

const (
	LogFormatJson    = "json"    // json日志格式
	LogFormatConsole = "console" // 控制台日志格式
)

const (
	AsyncChanSize = 4096 // 异步日志队列大小
)

// DefaultOptions 默认的日志配置
var DefaultOptions = Options{
	Filename:        "./log/default.log",
	Level:           zap.DebugLevel,
	MaxFileSize:     128,
	MaxAge:          60,
	MaxBackups:      1024,
	Stacktrace:      zap.ErrorLevel,
	FormatType:      LogFormatConsole,
	CallerSkip:      1,
	IsAsync:         false,
	IsCompress:      true,
	IsOpenPprof:     false,
	IsOpenConsole:   true,
	IsOpenFile:      false,
	IsOpenErrorFile: false,
	IsRedirectErr:   true,
}
