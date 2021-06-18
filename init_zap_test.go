package logx

import (
	"os"
	"testing"
)

func init() {
	Stdout = true
}
func Test_zap2(t *testing.T) {
	InitLoggerReset()
	Logger1.Infof("in main args:%v", os.Args)
	Logger1.Errorf("eerror %v", "error")
}
