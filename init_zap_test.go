package logx

import (
	"log"
	"os"
	"testing"

	"go.uber.org/zap/zapcore"
)

// https://github.com/uber-go/zap/blob/master/example_test.go

func Test_zap2(t *testing.T) {
	// InitLoggerReset()
	Infof("in main args:%v", os.Args)
	Errorf("eerror %v", "error")
}

func Test_diy(t *testing.T) {

	items := LogFileDefault()
	items = append(items, LogStdout()...)
	items = append(items, DiyWriter()...)
	Logger2, Logger1 = LOGX(items)

	Infof("in main args:%v", os.Args)
	Errorf("eerror %v", "error")
}

type diywirter struct {
}

func (w *diywirter) Write(p []byte) (n int, err error) {
	log.Printf("sys log: %v", string(p))
	return len(p), nil
}
func DiyWriter() []zapcore.Core {
	consoleDebugging := zapcore.Lock(zapcore.AddSync(&diywirter{}))
	// consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	var stdoutcore []zapcore.Core = []zapcore.Core{
		zapcore.NewCore(encoder, consoleDebugging, debugLevel),
	}
	return stdoutcore
}
