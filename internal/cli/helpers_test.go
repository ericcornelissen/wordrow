package cli

import "testing"

// CreateArgs creates a list of arguments in the same fashion as `os.Args`.
func createArgs(args ...string) []string {
	cliArgs := []string{"wordrow"}
	for _, arg := range args {
		cliArgs = append(cliArgs, arg)
	}

	return cliArgs
}

// Test if DryRun has the default value.
func testDefaultDryRun(t *testing.T, arguments Arguments) {
	t.Helper()

	if arguments.DryRun == true {
		t.Error("The default value for the DryRun option should be false")
	}
}

// Test if Invert has the default value.
func testDefaultInvert(t *testing.T, arguments Arguments) {
	t.Helper()

	if arguments.Invert == true {
		t.Error("The default value for the Invert option should be false")
	}
}

// Test if Silent has the default value.
func testDefaultSilent(t *testing.T, arguments Arguments) {
	t.Helper()

	if arguments.Silent == true {
		t.Error("The default value for the Silent option should be false")
	}
}

// Test if Verbose has the default value.
func testDefaultVerbose(t *testing.T, arguments Arguments) {
	t.Helper()

	if arguments.Verbose == true {
		t.Error("The default value for the Silent option should be false")
	}
}

// Test if MapFiles has the default value.
func testDefaultMapFiles(t *testing.T, arguments Arguments) {
	t.Helper()

	if len(arguments.MapFiles) != 0 {
		t.Error("The default list of MapFiles should be empty")
	}
}

// Test if Mappings has the default value.
func testDefaultMappings(t *testing.T, arguments Arguments) {
	t.Helper()

	if len(arguments.Mappings) != 0 {
		t.Error("The default list of Mappings should be empty")
	}
}

// Test if all default values of an Arguments instance except one.
func testDefaultsExcept(t *testing.T, arguments Arguments, exclude string) {
	t.Helper()

	if exclude != "dryrun" {
		testDefaultDryRun(t, arguments)
	}
	if exclude != "invert" {
		testDefaultInvert(t, arguments)
	}
	if exclude != "silent" {
		testDefaultSilent(t, arguments)
	}
	if exclude != "verbose" {
		testDefaultVerbose(t, arguments)
	}
	if exclude != "map files" {
		testDefaultMapFiles(t, arguments)
	}
	if exclude != "mappings" {
		testDefaultMappings(t, arguments)
	}
}
