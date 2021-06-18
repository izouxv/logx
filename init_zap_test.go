package logx

import (
	"os"
	"testing"
)

// https://github.com/uber-go/zap/blob/master/example_test.go

func init() {
	Stdout = true
}
func Test_zap2(t *testing.T) {
	InitLoggerReset()
	Logger1.Infof("in main args:%v", os.Args)
	Logger1.Errorf("eerror %v", "error")
}
