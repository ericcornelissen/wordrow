package logger


func ExampleDebug() {
  SetLogLevel(DEBUG)

  Debug("hello", "world!")
  // Output: [Debug] Hello world!
}

func ExampleDebugWhenInfo() {
  SetLogLevel(INFO)

  Debug("hello world!")
  // Output:
}

func ExampleDebugWhenWarning() {
  SetLogLevel(WARNING)

  Debug("hello world!")
  // Output:
}

func ExampleDebugWhenError() {
  SetLogLevel(ERROR)

  Debug("hello world!")
  // Output:
}

func ExampleDebugWhenFatal() {
  SetLogLevel(FATAL)

  Debug("hello world!")
  // Output:
}

func ExampleDebugf() {
  SetLogLevel(DEBUG)

  Debugf("%s %s!", "hello", "world")
  // Output: [Debug] Hello world!
}

func ExampleDebugfWhenInfo() {
  SetLogLevel(INFO)

  Debugf("%s", "hello world!")
  // Output:
}

func ExampleDebugfWhenWarning() {
  SetLogLevel(WARNING)

  Debugf("%s", "hello world!")
  // Output:
}

func ExampleDebugfWhenError() {
  SetLogLevel(ERROR)

  Debugf("%s", "hello world!")
  // Output:
}

func ExampleDebugfWhenFatal() {
  SetLogLevel(FATAL)

  Debugf("%s", "hello world!")
  // Output:
}

func ExampleError() {
  SetLogLevel(ERROR)

  Error("hello", "world!")
  // Output: [Error] Hello world!
}

func ExampleErrorWhenDebug() {
  SetLogLevel(DEBUG)

  Error("hello world!")
  // Output: [Error] Hello world!
}

func ExampleErrorWhenInfo() {
  SetLogLevel(INFO)

  Error("hello world!")
  // Output: [Error] Hello world!
}

func ExampleErrorWhenWarning() {
  SetLogLevel(WARNING)

  Error("hello world!")
  // Output: [Error] Hello world!
}

func ExampleErrorWhenFatal() {
  SetLogLevel(FATAL)

  Error("hello world!")
  // Output:
}

func ExampleErrorf() {
  SetLogLevel(ERROR)

  Errorf("%s %s!", "hello", "world")
  // Output: [Error] Hello world!
}

func ExampleErrorfWhenDebug() {
  SetLogLevel(DEBUG)

  Errorf("%s", "hello world!")
  // Output: [Error] Hello world!
}

func ExampleErrorfWhenInfo() {
  SetLogLevel(INFO)

  Errorf("%s", "hello world!")
  // Output: [Error] Hello world!
}

func ExampleErrorfWhenWarning() {
  SetLogLevel(WARNING)

  Errorf("%s", "hello world!")
  // Output: [Error] Hello world!
}

func ExampleErrorfWhenFatal() {
  SetLogLevel(FATAL)

  Errorf("%s", "hello world!")
  // Output:
}

func ExampleFatal() {
  SetLogLevel(FATAL)

  Fatal("hello", "world!")
  // Output: [Fatal] Hello world!
}

func ExampleFatalWhenDebug() {
  SetLogLevel(DEBUG)

  Fatal("hello world!")
  // Output: [Fatal] Hello world!
}

func ExampleFatalWhenInfo() {
  SetLogLevel(INFO)

  Fatal("hello world!")
  // Output: [Fatal] Hello world!
}

func ExampleFatalWhenWarning() {
  SetLogLevel(WARNING)

  Fatal("hello world!")
  // Output: [Fatal] Hello world!
}

func ExampleFatalWhenError() {
  SetLogLevel(ERROR)

  Fatal("hello world!")
  // Output: [Fatal] Hello world!
}

func ExampleFatalf() {
  SetLogLevel(FATAL)

  Fatalf("%s %s!", "hello", "world")
  // Output: [Fatal] Hello world!
}

func ExampleFatalfWhenDebug() {
  SetLogLevel(DEBUG)

  Fatalf("%s", "hello world!")
  // Output: [Fatal] Hello world!
}

func ExampleFatalfWhenInfo() {
  SetLogLevel(INFO)

  Fatalf("%s", "hello world!")
  // Output: [Fatal] Hello world!
}

func ExampleFatalfWhenWarning() {
  SetLogLevel(WARNING)

  Fatalf("%s", "hello world!")
  // Output: [Fatal] Hello world!
}

func ExampleFatalfWhenError() {
  SetLogLevel(ERROR)

  Fatalf("%s", "hello world!")
  // Output: [Fatal] Hello world!
}

func ExampleInfo() {
  SetLogLevel(INFO)

  Info("hello", "world!")
  // Output: [Info] Hello world!
}

func ExampleInfoWhenDebug() {
  SetLogLevel(DEBUG)

  Info("hello world!")
  // Output: [Info] Hello world!
}

func ExampleInfoWhenWarning() {
  SetLogLevel(WARNING)

  Info("hello world!")
  // Output:
}

func ExampleInfoWhenError() {
  SetLogLevel(ERROR)

  Info("hello world!")
  // Output:
}

func ExampleInfoWhenFatal() {
  SetLogLevel(FATAL)

  Info("hello world!")
  // Output:
}

func ExampleInfof() {
  SetLogLevel(INFO)

  Infof("%s %s!", "hello", "world")
  // Output: [Info] Hello world!
}

func ExampleInfofWhenDebug() {
  SetLogLevel(DEBUG)

  Infof("%s", "hello world!")
  // Output: [Info] Hello world!
}

func ExampleInfofWhenWarning() {
  SetLogLevel(WARNING)

  Infof("%s", "hello world!")
  // Output:
}

func ExampleInfofWhenError() {
  SetLogLevel(ERROR)

  Infof("%s", "hello world!")
  // Output:
}

func ExampleInfofWhenFatal() {
  SetLogLevel(FATAL)

  Infof("%s", "hello world!")
  // Output:
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

func ExampleWarningWhenDebug() {
  SetLogLevel(DEBUG)

  Warning("hello world!")
  // Output: [Warning] Hello world!
}

func ExampleWarningWhenInfo() {
  SetLogLevel(INFO)

  Warning("hello world!")
  // Output: [Warning] Hello world!
}

func ExampleWarningWhenError() {
  SetLogLevel(ERROR)

  Warning("hello world!")
  // Output:
}

func ExampleWarningWhenFatal() {
  SetLogLevel(FATAL)

  Warning("hello world!")
  // Output:
}

func ExampleWarningf() {
  SetLogLevel(WARNING)

  Warningf("%s %s!", "hello", "world")
  // Output: [Warning] Hello world!
}

func ExampleWarningfWhenDebug() {
  SetLogLevel(DEBUG)

  Warningf("%s", "hello world!")
  // Output: [Warning] Hello world!
}

func ExampleWarningfWhenInfo() {
  SetLogLevel(INFO)

  Warningf("%s", "hello world!")
  // Output: [Warning] Hello world!
}

func ExampleWarningfWhenError() {
  SetLogLevel(ERROR)

  Warningf("%s", "hello world!")
  // Output:
}

func ExampleWarningfWhenFatal() {
  SetLogLevel(FATAL)

  Warningf("%s", "hello world!")
  // Output:
}
