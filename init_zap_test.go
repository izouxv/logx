package logx

import (
	"log"
	"os"
	"testing"
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
	items = append(items, DiyWriter(&diywirter{}, DebugLevel)...)
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
