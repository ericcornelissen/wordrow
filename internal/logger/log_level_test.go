package logger

import "testing"

func TestLogLevelDebug(t *testing.T) {
	result := DEBUG.String()

	if result != "Debug" {
		t.Errorf("Unexpected string version of LogLevel DEBUG, got %s", result)
	}

	if DEBUG > INFO {
		t.Error("DEBUG should to be less than INFO")
	}

	if DEBUG > WARNING {
		t.Error("DEBUG should to be less than WARNING")
	}

	if DEBUG > ERROR {
		t.Error("DEBUG should to be less than ERROR")
	}

	if DEBUG > FATAL {
		t.Error("DEBUG should to be less than FATAL")
	}
}

func TestLogLevelInfo(t *testing.T) {
	result := INFO.String()

	if result != "Info" {
		t.Errorf("Unexpected string version of LogLevel DEBUG, got %s", result)
	}

	if INFO < DEBUG {
		t.Error("INFO should to be greater than DEBUG")
	}

	if INFO > WARNING {
		t.Error("INFO should to be less than WARNING")
	}

	if INFO > ERROR {
		t.Error("INFO should to be less than ERROR")
	}

	if INFO > FATAL {
		t.Error("INFO should to be less than FATAL")
	}
}

func TestLogLevelWarning(t *testing.T) {
	result := WARNING.String()

	if result != "Warning" {
		t.Errorf("Unexpected string version of LogLevel DEBUG, got %s", result)
	}

	if WARNING < DEBUG {
		t.Error("WARNING should to be greater than DEBUG")
	}

	if WARNING < INFO {
		t.Error("WARNING should to be greater than INFO")
	}

	if WARNING > ERROR {
		t.Error("WARNING should to be less than ERROR")
	}

	if WARNING > FATAL {
		t.Error("WARNING should to be less than FATAL")
	}
}

func TestLogLevelError(t *testing.T) {
	result := ERROR.String()

	if result != "Error" {
		t.Errorf("Unexpected string version of LogLevel DEBUG, got %s", result)
	}

	if ERROR < DEBUG {
		t.Error("ERROR should to be greater than DEBUG")
	}

	if ERROR < INFO {
		t.Error("ERROR should to be greater than INFO")
	}

	if ERROR < WARNING {
		t.Error("ERROR should to be greater than WARNING")
	}

	if ERROR > FATAL {
		t.Error("ERROR should to be less than FATAL")
	}
}

func TestLogLevelFatal(t *testing.T) {
	result := FATAL.String()

	if result != "Fatal" {
		t.Errorf("Unexpected string version of LogLevel DEBUG, got %s", result)
	}

	if FATAL < DEBUG {
		t.Error("FATAL should to be greater than DEBUG")
	}

	if FATAL < INFO {
		t.Error("FATAL should to be greater than INFO")
	}

	if FATAL < WARNING {
		t.Error("FATAL should to be greater than WARNING")
	}

	if FATAL < ERROR {
		t.Error("FATAL should to be greater than FATAL")
	}
}
