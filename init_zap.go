package logx

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// https://www.jianshu.com/p/d729c7ec9c85

var Logger1 *zap.SugaredLogger
var Logger2 *zap.Logger

func init() {
	items := LogFileDefault()
	items = append(items, LogStdout()...)
	Logger2, Logger1 = LOGX(items)
}
func DiyWriter(w io.Writer, levelFunc zap.LevelEnablerFunc) []zapcore.Core {
	consoleDebugging := zapcore.Lock(zapcore.AddSync(w))
	// consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	var stdoutcore []zapcore.Core = []zapcore.Core{
		zapcore.NewCore(Encoder, consoleDebugging, levelFunc),
	}
	return stdoutcore
}

func LogStdout() []zapcore.Core {
	consoleDebugging := zapcore.Lock(os.Stdout)
	// consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	var stdoutcore []zapcore.Core = []zapcore.Core{
		zapcore.NewCore(Encoder, consoleDebugging, DebugLevel),
	}
	return stdoutcore
}
func LogFileDefault() []zapcore.Core {
	var appName string = "demo"
	var logsFolder string = "./logs"
	return LogFile(appName, logsFolder)
}

func LogFile(appName string, logsFolder string) []zapcore.Core {
	getWriter := func(filename string) io.Writer {
		// 生成rotatelogs的Logger 实际生成的文件名 demo.log.YYmmddHH
		// demo.log是指向最新日志的链接
		// 保存7天内的日志，每1小时(整点)分割一次日志
		hook, err := rotatelogs.New(
			strings.Replace(filename, ".log", "", -1)+"-%Y%m%d%H.log", // 没有使用go风格反人类的format格式
			//rotatelogs.WithLinkName(filename),
			rotatelogs.WithMaxAge(time.Hour*24*2),
			//rotatelogs.WithRotationTime(time.Hour),
		)

		if err != nil {
			panic(err)
		}
		return hook
	}

	// 获取 info、error日志文件的io.Writer 抽象 getWriter() 在下方实现
	infoWriter := getWriter(fmt.Sprintf("%v/%v_info.log", logsFolder, appName))
	warnWriter := getWriter(fmt.Sprintf("%v/%v_warn.log", logsFolder, appName))
	errorWriter := getWriter(fmt.Sprintf("%v/%v_error.log", logsFolder, appName))

	var cores []zapcore.Core = []zapcore.Core{
		zapcore.NewCore(Encoder, zapcore.AddSync(infoWriter), InfoLevel),
		zapcore.NewCore(Encoder, zapcore.AddSync(warnWriter), WarnLevel),
		zapcore.NewCore(Encoder, zapcore.AddSync(errorWriter), ErrorLevel),
	}
	return cores
}

func LOGX(cores []zapcore.Core) (*zap.Logger, *zap.SugaredLogger) {
	core := zapcore.NewTee(cores...)
	log := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1)) // 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数, 有点小坑
	return log, log.Sugar()
}

// 设置一些基本日志格式 具体含义还比较好理解，直接看zap源码也不难懂
var Encoder = zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
	MessageKey:  "msg",
	LevelKey:    "level",
	EncodeLevel: zapcore.CapitalLevelEncoder,
	TimeKey:     "ts",
	EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02 15:04:05.999999"))
	},
	CallerKey:    "file",
	EncodeCaller: zapcore.ShortCallerEncoder,
	EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendInt64(int64(d) / 1000000)
	},
})

// 实现两个判断日志等级的interface
var DebugLevel = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
	return lvl >= zapcore.DebugLevel
})
var InfoLevel = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
	return lvl >= zapcore.InfoLevel
})
var WarnLevel = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
	return lvl >= zapcore.WarnLevel
})
var ErrorLevel = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
	return lvl >= zapcore.ErrorLevel
})
