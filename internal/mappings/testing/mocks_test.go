package testing

// MockT is a basic mock of `testing.T`.
type mockT struct {
	// Counts the number of calls to `t.Error()`.
	errorCallCount int

	// Counts the number of calls to `t.Helper()`.
	helperCallCount int
}

func (mockT *mockT) Errorf(format string, args ...interface{}) {
	mockT.errorCallCount++
}

func (mockT *mockT) Fatalf(format string, args ...interface{}) {
	panic("fatal was called")
}

func (mockT *mockT) Helper() {
	mockT.helperCallCount++
}

// Create a new `testing.T` mock.
func newMockT() mockT {
	return mockT{
		errorCallCount:  0,
		helperCallCount: 0,
	}
}
