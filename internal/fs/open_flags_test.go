package fs

import "testing"

func TestFlagToString(t *testing.T) {
	t.Run("ReadOnly", func(t *testing.T) {
		result := OReadOnly.String()
		if result != "ReadOnly" {
			t.Errorf("Unexpected string for OReadOnly (got '%s')", result)
		}
	})
	t.Run("ReadWrite", func(t *testing.T) {
		result := OReadWrite.String()
		if result != "ReadWrite" {
			t.Errorf("Unexpected string for OReadWrite (got '%s')", result)
		}
	})
}

func TestFlagToFlag(t *testing.T) {
	checkForPanic := func(t *testing.T) {
		t.Helper()
		if r := recover(); r != nil {
			t.Fatal("Recovery should never be necessary")
		}
	}

	t.Run("ReadOnly", func(t *testing.T) {
		defer checkForPanic(t)
		flagToFlag(OReadOnly)
	})
	t.Run("ReadWrite", func(t *testing.T) {
		defer checkForPanic(t)
		flagToFlag(OReadWrite)
	})
	t.Run("ReadWrite", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Fatal("Panic required for invalid Flag value")
			}
		}()

		flagToFlag(-1)
	})
}
