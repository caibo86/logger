# logger
一个在zap基础上封装的日志库，增加了易用性及功能性

## Features
___
- 参数化配置日志 
- 同时输出多种日志后台：终端日志、文件日志、网络日志
- 高级别日志独立输出
- 重定向panic日志
- 同步异步可配置
- 日志输出格式可选：json、console
- 时区可选：UTC、Local
- 堆栈信息输出
- 日志分割
- 日志压缩
- 日志清理

## Logger Levels
___
- Debug
- Info
- Warn
- Error
- Panic
- Fatal

## Options
___
```golang
package logger

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
```

## Default
___
```golang
package logger

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

```

## API
___
**初始化与关闭**：
- Init 使用选项函数初始化日志服务
- Close 关闭日志服务
___
**选项配置**：
- SetFilename：设置日志文件路径
- SetLevel：设置日志级别
- SetStacktrace：设置记录堆栈的日志级别
- SetMaxFileSize：设置日志分割的尺寸
- SetMaxAge：设置日志保存的时间 单位:天
- SetMaxBackups：设置最大日志数量
- SetFormatType：设置日志格式
- SetCallerSkip：设置堆栈的跳过层数
- SetIsAsync：设置异步日志
- SetIsCompress：设置是否压缩
- SetIsOpenPprof：设置是否打开pprof
- SetIsOpenConsole：设置是否打开终端标准输出
- SetIsOpenFile：设置是否打开文件日志
- SetIsOpenErrorFile：设置是否打开高级别错误文件日志
- SetIsRedirectErr：设置是否重定向标准错误输出
___
**记录日志**：
- Debug：打印Debug级别日志,自动参数
- Debugf：打印Debug级别日志,格式化参数
- Debugw：打印Debug级别日志,KV参数
- Info：打印Info级别日志,自动参数
- Infof：打印Info级别日志,格式化参数
- Infow：打印Info级别日志,KV参数
- Warn：打印Warn级别日志,自动参数
- Warnf：打印Warn级别日志,格式化参数
- Warnw：打印Warn级别日志,KV参数
- Error：打印Error级别日志,自动参数
- Errorf：打印Error级别日志,格式化参数
- Errorw：打印Error级别日志,KV参数
- Panic：打印Panic级别日志,自动参数
- Panicf：打印Panic级别日志,格式化参数
- Panicw：打印Panic级别日志,KV参数
- Fatal：打印Fatal级别日志,自动参数
- Fatalf：打印Fatal级别日志,格式化参数
- Fatalw：打印Fatal级别日志,KV参数
## Example
___
### JSON format

```golang
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
		logger.SetFormatType(logger.LogFormatJson),
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

```
**Output**:
```shell
{"level":"info","ts":"2024-11-20 07:29:16.268","caller":"example/main.go:15","msg":"successfully initialized logger service"}
{"level":"debug","ts":"2024-11-20 07:29:16.268","caller":"runtime/proc.go:267","msg":"helloworld2024"}
{"level":"debug","ts":"2024-11-20 07:29:16.268","caller":"runtime/proc.go:267","msg":"hello world 2004"}
{"level":"debug","ts":"2024-11-20 07:29:16.268","caller":"runtime/proc.go:267","msg":"hello world","foo":"bar"}
{"level":"error","ts":"2024-11-20 07:29:16.268","caller":"runtime/proc.go:267","msg":"we have an error","stacktrace":"runtime.main\n\t/usr/local/go/src/runtime/proc.go:267"}
{"level":"fatal","ts":"2024-11-20 07:29:16.268","caller":"runtime/proc.go:267","msg":"helloworld2024","stacktrace":"runtime.main\n\t/usr/local/go/src/runtime/proc.go:267"}
exit status 1

```
### Console format

```golang
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

```

**Output**:
```shell
2024-11-20 07:31:56.089 test    INFO    example/main.go:15      successfully initialized logger service
2024-11-20 07:31:56.089 test    DEBUG   runtime/proc.go:267     helloworld2024
2024-11-20 07:31:56.089 test    DEBUG   runtime/proc.go:267     hello world 2004
2024-11-20 07:31:56.089 test    DEBUG   runtime/proc.go:267     hello world     {"foo": "bar"}
2024-11-20 07:31:56.090 test    ERROR   runtime/proc.go:267     we have an error
runtime.main
        /usr/local/go/src/runtime/proc.go:267
2024-11-20 07:31:56.090 test    FATAL   runtime/proc.go:267     helloworld2024
runtime.main
        /usr/local/go/src/runtime/proc.go:267
exit status 1

```
## Note
___