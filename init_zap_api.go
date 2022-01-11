package logx

func Debug(args ...interface{}) {
	Logger1.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	Logger1.Debugf(template, args...)
}

func Info(args ...interface{}) {
	Logger1.Info(args...)
}

func Infof(template string, args ...interface{}) {
	Logger1.Infof(template, args...)
}

func Print(args ...interface{}) {
	Logger1.Info(args...)
}

func Printf(template string, args ...interface{}) {
	Logger1.Infof(template, args...)
}

func Warn(args ...interface{}) {
	Logger1.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	Logger1.Warnf(template, args...)
}

func Error(args ...interface{}) {
	Logger1.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	Logger1.Errorf(template, args...)
}

func DPanic(args ...interface{}) {
	Logger1.DPanic(args...)
}

func DPanicf(template string, args ...interface{}) {
	Logger1.DPanicf(template, args...)
}

func Panic(args ...interface{}) {
	Logger1.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	Logger1.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	Logger1.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	Logger1.Fatalf(template, args...)
}

// func testst() {
// 	// createAndCheckToken()
// 	errorLogger.Infof("in main args:%v", os.Args)

// 	errorLogger.Errorf("eerror %v", "error")
// 	// flag.Parse()
// 	// errorLogger.Infof("env is %v", *env)
// 	// config := util.InitConfig("./config/" + *env + ".conf")
// 	// ip := config["ip"]
// 	// port := config["port"]
// 	// envConfig := config["env"]
// 	// errorLogger.Infof("ip=%v, port=%v, env=%v", ip, port, envConfig)
// }
