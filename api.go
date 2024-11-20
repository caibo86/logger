// -------------------------------------------
// @file      : api.go
// @author    : bo cai
// @contact   : caibo923@gmail.com
// @time      : 2024/11/18 下午6:34
// -------------------------------------------

package logger

import (
	"go.uber.org/zap/zapcore"
)

// Init 全局日志初始化 每个app必须调用一次
func Init(options ...Option) {
	once.Do(func() {
		global.Init(options...)
	})
	// 触发创建目录
	Info("successfully initialized logger service")
	if global.options.IsRedirectErr {
		err := redirectStdErrLog()
		if err != nil {
			Errorf("redirect panic log err: %s", err)
		}
	}
}

// Close 关闭日志服务时调用
func Close() error {
	if global != nil {
		err := global.Close()
		if err != nil && err.Error() == "sync /dev/stdout: invalid argument" ||
			err.Error() == "sync /dev/stdout: The handle is invalid." {
			return nil
		}
		return err
	}
	return nil
}

// Debug 打印Debug级别日志,自动参数
func Debug(args ...interface{}) {
	if global.zapLogger != nil {
		global.zapLogger.Debug(args...)
	}
}

// Debugf 打印Debug级别日志,格式化参数
func Debugf(template string, args ...interface{}) {
	if global.zapLogger != nil {
		global.zapLogger.Debugf(template, args...)
	}
}

// Debugw 打印Debug级别日志,键值对参数
func Debugw(msg string, keysAndValues ...interface{}) {
	if global.zapLogger != nil {
		global.zapLogger.Debugw(msg, keysAndValues...)
	}
}

// Info 打印Info级别日志,自动参数
func Info(args ...interface{}) {
	if global.zapLogger != nil {
		global.zapLogger.Info(args...)
	}
}

// Infof 打印Info级别日志,格式化参数
func Infof(template string, args ...interface{}) {
	if global.zapLogger != nil {
		global.zapLogger.Infof(template, args...)
	}
}

// Infow 打印Info级别日志,键值对参数
func Infow(msg string, keysAndValues ...interface{}) {
	if global.zapLogger != nil {
		global.zapLogger.Infow(msg, keysAndValues...)
	}
}

// Warn 打印Warn级别日志,自动参数
func Warn(args ...interface{}) {
	if global.zapLogger != nil {
		global.zapLogger.Warn(args...)
	}
}

// Warnf 打印Warn级别日志,格式化参数
func Warnf(template string, args ...interface{}) {
	if global.zapLogger != nil {
		global.zapLogger.Warnf(template, args...)
	}
}

// Warnw 打印Warn级别日志,键值对参数
func Warnw(msg string, keysAndValues ...interface{}) {
	if global.zapLogger != nil {
		global.zapLogger.Warnw(msg, keysAndValues...)
	}
}

// Error 打印Error级别日志,自动参数
func Error(args ...interface{}) {
	if global.zapLogger != nil {
		global.zapLogger.Error(args...)
	}
}

// Errorf 打印Error级别日志,格式化参数
func Errorf(template string, args ...interface{}) {
	if global.zapLogger != nil {
		global.zapLogger.Errorf(template, args...)
	}
}

// Errorw 打印Error级别日志,键值对参数
func Errorw(msg string, keysAndValues ...interface{}) {
	if global.zapLogger != nil {
		global.zapLogger.Errorw(msg, keysAndValues...)
	}
}

// Panic 打印Panic级别日志,自动参数
func Panic(args ...interface{}) {
	if global.zapLogger != nil {
		global.zapLogger.Panic(args...)
	}
}

// Panicf 打印Panic级别日志,格式化参数
func Panicf(template string, args ...interface{}) {
	if global.zapLogger != nil {
		global.zapLogger.Panicf(template, args...)
	}
}

// Panicw 打印Panic级别日志,键值对参数
func Panicw(msg string, keysAndValues ...interface{}) {
	if global.zapLogger != nil {
		global.zapLogger.Panicw(msg, keysAndValues...)
	}
}

// Fatal 打印Fatal级别日志,自动参数
func Fatal(args ...interface{}) {
	if global.zapLogger != nil {
		global.zapLogger.Fatal(args...)
	}
}

// Fatalf 打印Fatal级别日志,格式化参数
func Fatalf(template string, args ...interface{}) {
	if global.zapLogger != nil {
		global.zapLogger.Fatalf(template, args...)
	}
}

// Fatalw 打印Fatal级别日志,键值对参数
func Fatalw(msg string, keysAndValues ...interface{}) {
	if global.zapLogger != nil {
		global.zapLogger.Fatalw(msg, keysAndValues...)
	}
}

// SetFilename 设置日志文件路径
func SetFilename(filename string) Option {
	return func(options *Options) {
		if filename != "" {
			options.Filename = filename
		}
	}
}

// SetLevel 设置日志级别
func SetLevel(level zapcore.Level) Option {
	return func(options *Options) {
		options.Level = level
	}
}

// SetMaxFileSize 设置日志分割的尺寸
func SetMaxFileSize(size int) Option {
	return func(options *Options) {
		if size > 0 {
			options.MaxFileSize = size
		}
	}
}

// SetMaxAge 设置日志保存的时间
func SetMaxAge(age int) Option {
	return func(options *Options) {
		options.MaxAge = age
	}
}

// SetMaxBackups 设置最大日志数量
func SetMaxBackups(backups int) Option {
	return func(options *Options) {
		options.MaxBackups = backups
	}
}

// SetStacktrace 设置记录堆栈的日志级别
func SetStacktrace(level zapcore.Level) Option {
	return func(options *Options) {
		options.Stacktrace = level
	}
}

// SetIsOpenConsole 设置是否打开终端标准输出
func SetIsOpenConsole(console bool) Option {
	return func(options *Options) {
		options.IsOpenConsole = console
	}
}

// SetFormatType 设置日志格式
func SetFormatType(format string) Option {
	return func(options *Options) {
		options.FormatType = format
	}
}

// SetCallerSkip 设置堆栈的跳过层数
func SetCallerSkip(callerSkip int) Option {
	return func(options *Options) {
		options.CallerSkip = callerSkip
	}
}

// SetIsAsync 设置是否异步日志
func SetIsAsync(async bool) Option {
	return func(options *Options) {
		options.IsAsync = async
	}
}

// SetIsCompress 设置是否压缩
func SetIsCompress(compress bool) Option {
	return func(options *Options) {
		options.IsCompress = compress
	}
}

// SetIsOpenPprof 设置是否打开pprof
func SetIsOpenPprof(pprof bool) Option {
	return func(options *Options) {
		options.IsOpenPprof = pprof
	}
}

// SetIsOpenFile 设置是否打开文件日志
func SetIsOpenFile(file bool) Option {
	return func(options *Options) {
		options.IsOpenFile = file
	}
}

// SetIsOpenErrorFile 设置是否打开高级别错误文件日志
func SetIsOpenErrorFile(errorFile bool) Option {
	return func(options *Options) {
		options.IsOpenErrorFile = errorFile
	}
}

// SetIsRedirectErr 设置是否重定向标准错误输出
func SetIsRedirectErr(redirect bool) Option {
	return func(options *Options) {
		options.IsRedirectErr = redirect
	}
}
