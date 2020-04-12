package logger


func ExampleDebug() {
  Debug("hello", "world!")
  // Output: [D] Hello world!
}

func ExampleDebugf() {
  Debugf("%s %s!", "hello", "world")
  // Output: [D] Hello world!
}

func ExampleError() {
  Error("hello", "world!")
  // Output: [E] Hello world!
}

func ExampleErrorf() {
  Errorf("%s %s!", "hello", "world")
  // Output: [E] Hello world!
}

func ExampleFatal() {
  Fatal("hello", "world!")
  // Output: [F] Hello world!
}

func ExampleFatalf() {
  Fatalf("%s %s!", "hello", "world")
  // Output: [F] Hello world!
}

func ExampleInfo() {
  Info("hello", "world!")
  // Output: [I] Hello world!
}

func ExampleInfof() {
  Infof("%s %s!", "hello", "world")
  // Output: [I] Hello world!
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
  Warning("hello", "world!")
  // Output: [W] Hello world!
}

func ExampleWarningf() {
  Warningf("%s %s!", "hello", "world")
  // Output: [W] Hello world!
}
