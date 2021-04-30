package logx

import (
	"os"
	"testing"
)

func Test_zap2(t *testing.T) {
	Logger1.Infof("in main args:%v", os.Args)
	Logger1.Errorf("eerror %v", "error")
}
