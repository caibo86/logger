// -------------------------------------------
// @file      : option.go
// @author    : bo cai
// @contact   : caibo923@gmail.com
// @time      : 2024/11/19 上午10:17
// -------------------------------------------

package logger

import (
	"github.com/caibo86/cberrors"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
	"time"
)

// Option 日志配置项
type Option func(options *Options)

// Options 日志配置
type Options struct {
	Filename        string        // 日志文件路径
	Level           zapcore.Level // 日志级别
	Stacktrace      zapcore.Level // 记录堆栈的日志级别
	MaxFileSize     int           // 日志分割的尺寸
	MaxAge          int           // 日志保存的时间 单位:天
	MaxBackups      int           // 最大日志数量
	FormatType      string        // 日志格式
	CallerSkip      int           // 堆栈的跳过层数
	IsAsync         bool          // 异步日志
	IsCompress      bool          // 是否压缩
	IsOpenPprof     bool          // 是否打开pprof
	IsOpenConsole   bool          // 是否打开终端标准输出
	IsOpenFile      bool          // 是否打开文件日志
	IsOpenErrorFile bool          // 是否打开高级别错误文件日志
	IsRedirectErr   bool          // 是否重定向标准错误输出
}

// 获取日志文件名,不包含文件后缀名,默认为Unknown
func (o *Options) getLogFilename() string {
	ret := "Unknown"
	arr := strings.Split(o.Filename, "/")
	if len(arr) <= 0 {
		return ret
	}
	ret = arr[len(arr)-1]
	if tmp := strings.Split(ret, "."); len(tmp) > 0 {
		return tmp[0]
	}
	return ret
}

// 终端日志核心
func (o *Options) getConsoleCore() zapcore.Core {
	var consoleWS zapcore.WriteSyncer
	if o.IsAsync {
		consoleWS = &zapcore.BufferedWriteSyncer{
			WS:   zapcore.AddSync(os.Stdout),
			Size: AsyncChanSize,
		}
	} else {
		consoleWS = zapcore.AddSync(os.Stdout)
	}
	var enabler zap.LevelEnablerFunc
	enabler = func(level zapcore.Level) bool {
		return level >= global.atom.Level()
	}
	return zapcore.NewCore(o.getEncoder(), consoleWS, enabler)
}

// 文件日志核心
func (o *Options) getFileCore() zapcore.Core {
	var fileLogger zapcore.WriteSyncer
	fileLogger = &FileLogger{
		Logger: &lumberjack.Logger{
			Filename:   o.Filename,
			MaxSize:    o.MaxFileSize,
			MaxBackups: o.MaxBackups,
			MaxAge:     o.MaxAge,
			Compress:   o.IsCompress,
		},
	}
	if o.IsAsync {
		fileLogger = &zapcore.BufferedWriteSyncer{
			WS:   fileLogger,
			Size: AsyncChanSize,
		}
	}
	var enabler zap.LevelEnablerFunc
	enabler = func(level zapcore.Level) bool {
		return level >= global.atom.Level() // && level <= zap.ErrorLevel
	}
	return zapcore.NewCore(o.getEncoder(), fileLogger, enabler)
}

// 高级别错误文件日志核心
func (o *Options) getErrorFileCore() zapcore.Core {
	// 高级别错误文件日志直接采用同步写
	errFilename := strings.Replace(o.Filename, ".log", ".err", 1)
	fileLogger := &FileLogger{
		Logger: &lumberjack.Logger{
			Filename:   errFilename,
			MaxSize:    o.MaxFileSize,
			MaxBackups: o.MaxBackups,
			MaxAge:     o.MaxAge,
			Compress:   o.IsCompress,
		},
	}
	var enabler zap.LevelEnablerFunc
	enabler = func(level zapcore.Level) bool {
		return level >= zap.ErrorLevel
	}
	return zapcore.NewCore(o.getEncoder(), fileLogger, enabler)
}

// 获取核心
func (o *Options) getCore() zapcore.Core {
	var cores []zapcore.Core
	if o.IsOpenConsole {
		cores = append(cores, o.getConsoleCore())
	}
	if o.IsOpenFile {
		cores = append(cores, o.getFileCore())
	}
	if o.IsOpenErrorFile {
		cores = append(cores, o.getErrorFileCore())
	}
	if len(cores) == 0 {
		cberrors.Panic("At least one log output needs to be opened")
	}
	return zapcore.NewTee(cores...)
}

// GetZapLogger 根据配置选项，获取zap logger对象
func (o *Options) GetZapLogger() *zap.SugaredLogger {
	var options []zap.Option
	options = append(options, zap.AddStacktrace(o.Stacktrace))
	options = append(options, zap.AddCaller())
	options = append(options, zap.AddCallerSkip(o.CallerSkip))

	logger := zap.New(o.getCore(), options...).Sugar()
	if logger == nil {
		cberrors.Panic("get zap logger failed.")
	}
	return logger
}

// 根据配置项获取encoder
func (o *Options) getEncoder() zapcore.Encoder {
	rfc339UTC := func(t time.Time, pe zapcore.PrimitiveArrayEncoder) {
		// 取UTC时间
		t = t.UTC()
		format := "2006-01-02 15:04:05.000"
		type appendTimeEncoder interface {
			AppendTimeLayout(time.Time, string)
		}
		if encoder, ok := pe.(appendTimeEncoder); ok {
			encoder.AppendTimeLayout(t, format)
			return
		}
		filename := o.getLogFilename()
		if filename != "" {
			pe.AppendString(t.Format(format) + "\t" + filename)
			return
		}
		pe.AppendString(t.Format(format))
	}
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = rfc339UTC
	if o.FormatType == LogFormatJson {
		// json格式encoder
		return zapcore.NewJSONEncoder(encoderConfig)
	}
	// 控制台格式encoder
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}
