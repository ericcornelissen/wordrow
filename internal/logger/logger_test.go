package logger

func ExampleDebug() {
	SetLogLevel(DEBUG)

	Debug("hello", "world!")
	// Output: [Debug] Hello world!
}

func ExampleDebugf() {
	SetLogLevel(DEBUG)

	Debugf("%s %s!", "hello", "world")
	// Output: [Debug] Hello world!
}

func ExampleError() {
	SetLogLevel(ERROR)

	Error("hello", "world!")
	// Output: [Error] Hello world!
}

func ExampleErrorf() {
	SetLogLevel(ERROR)

	Errorf("%s %s!", "hello", "world")
	// Output: [Error] Hello world!
}

func ExampleFatal() {
	SetLogLevel(FATAL)

	Fatal("hello", "world!")
	// Output: [Fatal] Hello world!
}

func ExampleFatalf() {
	SetLogLevel(FATAL)

	Fatalf("%s %s!", "hello", "world")
	// Output: [Fatal] Hello world!
}

func ExampleInfo() {
	SetLogLevel(INFO)

	Info("hello", "world!")
	// Output: [Info] Hello world!
}

func ExampleInfof() {
	SetLogLevel(INFO)

	Infof("%s %s!", "hello", "world")
	// Output: [Info] Hello world!
}

func ExamplePrintln() {
	Println("hello", "world!")
	// Output: hello world!
}

func ExamplePrintf() {
	Printf("%s %s!", "hello", "world")
	// Output: hello world!
}

func ExampleWarning() {
	SetLogLevel(WARNING)

	Warning("hello", "world!")
	// Output: [Warning] Hello world!
}

func ExampleWarningf() {
	SetLogLevel(WARNING)

	Warningf("%s %s!", "hello", "world")
	// Output: [Warning] Hello world!
}
