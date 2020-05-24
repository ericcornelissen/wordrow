package cli

import "fmt"
import "testing"

func createArgs(args ...string) []string {
	cliArgs := []string{"wordrow"}
	for _, arg := range args {
		cliArgs = append(cliArgs, arg)
	}

	return cliArgs
}

func testDefaultDryRun(t *testing.T, arguments Arguments) {
	t.Helper()

	if arguments.DryRun == true {
		t.Error("The default value for the DryRun option should be false")
	}
}

func testDefaultInvert(t *testing.T, arguments Arguments) {
	t.Helper()

	if arguments.Invert == true {
		t.Error("The default value for the Invert option should be false")
	}
}

func testDefaultSilent(t *testing.T, arguments Arguments) {
	t.Helper()

	if arguments.Silent == true {
		t.Error("The default value for the Silent option should be false")
	}
}

func testDefaultVerbose(t *testing.T, arguments Arguments) {
	t.Helper()

	if arguments.Verbose == true {
		t.Error("The default value for the Silent option should be false")
	}
}

func testDefaultConfigFile(t *testing.T, arguments Arguments) {
	t.Helper()

	if arguments.ConfigFile != "" {
		t.Error("The default ConfigFile should not be set")
	}
}

func testDefaultMapFiles(t *testing.T, arguments Arguments) {
	t.Helper()

	if len(arguments.MapFiles) != 0 {
		t.Error("The default list of MapFiles should be empty")
	}
}

func testDefaultMappings(t *testing.T, arguments Arguments) {
	t.Helper()

	if len(arguments.Mappings) != 0 {
		t.Error("The default list of Mappings should be empty")
	}
}

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
	if exclude != "config file" {
		testDefaultConfigFile(t, arguments)
	}
	if exclude != "map files" {
		testDefaultMapFiles(t, arguments)
	}
	if exclude != "mappings" {
		testDefaultMappings(t, arguments)
	}
}

func ExampleParseArgs() {
	args := []string{"wordrow", "foo.bar"}
	run, _ := ParseArgs(args)
	fmt.Print(run)
	// Output: true
}

func TestNoArgs(t *testing.T) {
	args := createArgs()
	run, _ := ParseArgs(args)

	if run == true {
		t.Error("The first return value should be false if no args are given")
	}
}

func TestEmptyArgument(t *testing.T) {
	args := createArgs("")
	ParseArgs(args)
	// Expect no panic
}

func TestHelpArg(t *testing.T) {
	args := createArgs(helpFlag.name)
	run, _ := ParseArgs(args)

	if run == true {
		t.Error("The first return value should be false for --help")
	}
}

func TestVersionArg(t *testing.T) {
	t.Run("--version only", func(t *testing.T) {
		args := createArgs(versionFlag.name)
		run, _ := ParseArgs(args)

		if run == true {
			t.Fatal("The first return value should be false if --version is the only argument")
		}
	})
	t.Run("--version and other", func(t *testing.T) {
		args := createArgs(versionFlag.name, "foo.bar")
		run, _ := ParseArgs(args)

		if run != true {
			t.Fatal("The first return value should be true for this test")
		}
	})
}

func TestDefaultOptions(t *testing.T) {
	args := createArgs("foo.bar")
	run, arguments := ParseArgs(args)

	if run != true {
		t.Fatal("The first return value should be true for this test")
	}

	testDefaultsExcept(t, arguments, "")

	if len(arguments.InputFiles) != 1 {
		t.Error("The list of InputFiles should contain a single file")
	}
}

func TestDryRunFlag(t *testing.T) {
	args := createArgs(dryRunFlag.name, "foo.bar")
	run, arguments := ParseArgs(args)

	if run != true {
		t.Fatal("The first return value should be true for this test")
	}

	testDefaultsExcept(t, arguments, "dryrun")

	if arguments.DryRun != true {
		t.Errorf("The DryRun value should be true if %s is an argument", dryRunFlag)
	}
}

func TestInvertFlag(t *testing.T) {
	t.Run(invertFlag.name, func(t *testing.T) {
		args := createArgs(invertFlag.name, "foo.bar")
		run, arguments := ParseArgs(args)

		if run != true {
			t.Fatal("The first return value should be true for this test")
		}

		testDefaultsExcept(t, arguments, "invert")

		if arguments.Invert != true {
			t.Errorf("The Invert value should be true if %s is an argument", invertFlag)
		}
	})
	t.Run(invertFlag.alias, func(t *testing.T) {
		args := createArgs(invertFlag.alias, "foo.bar")
		run, arguments := ParseArgs(args)

		if run != true {
			t.Fatal("The first return value should be true for this test")
		}

		testDefaultsExcept(t, arguments, "invert")

		if arguments.Invert != true {
			t.Errorf("The Invert value should be true if %s is an argument", invertFlag.alias)
		}
	})
}

func TestSilentFlag(t *testing.T) {
	t.Run(silentFlag.name, func(t *testing.T) {
		args := createArgs(silentFlag.name, "foo.bar")
		run, arguments := ParseArgs(args)

		if run != true {
			t.Fatal("The first return value should be true for this test")
		}

		testDefaultsExcept(t, arguments, "silent")

		if arguments.Silent != true {
			t.Errorf("The Silent value should be true if %s is an argument", silentFlag)
		}
	})
	t.Run(silentFlag.alias, func(t *testing.T) {
		args := createArgs(silentFlag.alias, "foo.bar")
		run, arguments := ParseArgs(args)

		if run != true {
			t.Fatal("The first return value should be true for this test")
		}

		testDefaultsExcept(t, arguments, "silent")

		if arguments.Silent != true {
			t.Errorf("The Silent value should be true if %s is an argument", silentFlag.alias)
		}
	})
}

func TestVerboseFlag(t *testing.T) {
	t.Run(verboseFlag.name, func(t *testing.T) {
		args := createArgs(verboseFlag.name, "foo.bar")
		run, arguments := ParseArgs(args)

		if run != true {
			t.Fatal("The first return value should be true for this test")
		}

		testDefaultsExcept(t, arguments, "verbose")

		if arguments.Verbose != true {
			t.Errorf("The Verbose value should be true if %s is an argument", verboseFlag)
		}
	})
	t.Run(verboseFlag.alias, func(t *testing.T) {
		args := createArgs(verboseFlag.alias, "foo.bar")
		run, arguments := ParseArgs(args)

		if run != true {
			t.Fatal("The first return value should be true for this test")
		}

		testDefaultsExcept(t, arguments, "verbose")

		if arguments.Verbose != true {
			t.Errorf("The Verbose value should be true if %s is an argument", verboseFlag.alias)
		}
	})
}

func TestConfigFileOption(t *testing.T) {
	configFile := "config.json"

	t.Run(configOption.name, func(t *testing.T) {
		args := createArgs(configOption.name, configFile, "foo.bar")
		run, arguments := ParseArgs(args)

		if run != true {
			t.Fatal("The first return value should be true for this test")
		}

		testDefaultsExcept(t, arguments, "config file")

		if arguments.ConfigFile != configFile {
			t.Errorf("The config file was incorrect (was '%s')", arguments.ConfigFile)
		}
	})
	t.Run(configOption.alias, func(t *testing.T) {
		args := createArgs(configOption.alias, configFile, "foo.bar")
		run, arguments := ParseArgs(args)

		if run != true {
			t.Fatal("The first return value should be true for this test")
		}

		testDefaultsExcept(t, arguments, "config file")

		if arguments.ConfigFile != configFile {
			t.Errorf("The config file was incorrect (was '%s')", arguments.ConfigFile)
		}
	})
	t.Run("multiple configuration files (overrides)", func(t *testing.T) {
		otherConfigFile := "foobar.json"

		args := createArgs(configOption.name, configFile, configOption.name, otherConfigFile, "foo.bar")
		run, arguments := ParseArgs(args)

		if run != true {
			t.Fatal("The first return value should be true for this test")
		}

		testDefaultsExcept(t, arguments, "config file")

		if arguments.ConfigFile != otherConfigFile {
			t.Errorf("The config file was incorrect (was '%s')", arguments.ConfigFile)
		}
	})
}

func TestConfigFileOptionIncorrect(t *testing.T) {
	t.Run("value missing", func(t *testing.T) {
		args := createArgs(configOption.name)
		run, _ := ParseArgs(args)

		if run != false {
			t.Error("The first return value should be false if there is an error in the args")
		}
	})
	t.Run("other flag", func(t *testing.T) {
		args := createArgs(configOption.name, silentFlag.name)
		run, _ := ParseArgs(args)

		if run != false {
			t.Error("The first return value should be false if there is an error in the args")
		}
	})
}

func TestMapFileOption(t *testing.T) {
	mapFile := "foo.map"

	t.Run(mapfileOption.name, func(t *testing.T) {
		args := createArgs(mapfileOption.name, mapFile, "foo.bar")
		run, arguments := ParseArgs(args)

		if run != true {
			t.Fatal("The first return value should be true for this test")
		}

		testDefaultsExcept(t, arguments, "map files")

		if len(arguments.MapFiles) != 1 {
			t.Fatalf("The MapFiles list have length 1 (was %d)", len(arguments.MapFiles))
		}

		if arguments.MapFiles[0] != mapFile {
			t.Errorf("First map file was incorrect (was '%s')", arguments.MapFiles[0])
		}
	})
	t.Run(mapfileOption.alias, func(t *testing.T) {
		args := createArgs(mapfileOption.alias, mapFile, "foo.bar")
		run, arguments := ParseArgs(args)

		if run != true {
			t.Fatal("The first return value should be true for this test")
		}

		testDefaultsExcept(t, arguments, "map files")

		if len(arguments.MapFiles) != 1 {
			t.Fatalf("The MapFiles list have length 1 (was %d)", len(arguments.MapFiles))
		}

		if arguments.MapFiles[0] != mapFile {
			t.Errorf("First map file was incorrect (was '%s')", arguments.MapFiles[0])
		}
	})
	t.Run("multiple map files", func(t *testing.T) {
		otherMapFile := "bar.map"

		args := createArgs(mapfileOption.name, mapFile, mapfileOption.name, otherMapFile, "foo.bar")
		run, arguments := ParseArgs(args)

		if run != true {
			t.Fatal("The first return value should be true for this test")
		}

		testDefaultsExcept(t, arguments, "map files")

		if len(arguments.MapFiles) != 2 {
			t.Fatalf("The MapFiles list have length 2 (was %d)", len(arguments.MapFiles))
		}

		if arguments.MapFiles[0] != mapFile {
			t.Errorf("First map file was incorrect (was '%s')", arguments.MapFiles[0])
		}

		if arguments.MapFiles[1] != otherMapFile {
			t.Errorf("First map file was incorrect (was '%s')", arguments.MapFiles[1])
		}
	})
}

func TestMapFileOptionIncorrect(t *testing.T) {
	t.Run("value missing", func(t *testing.T) {
		args := createArgs(mapfileOption.name)
		run, _ := ParseArgs(args)

		if run != false {
			t.Error("The first return value should be false if there is an error in the args")
		}
	})
	t.Run("other flag", func(t *testing.T) {
		args := createArgs(mapfileOption.name, silentFlag.name)
		run, _ := ParseArgs(args)

		if run != false {
			t.Error("The first return value should be false if there is an error in the args")
		}
	})
}

func TestMappingOption(t *testing.T) {
	mapping := "foo,bar"

	t.Run(mappingOption.name, func(t *testing.T) {
		args := createArgs(mappingOption.name, mapping, "foo.bar")
		run, arguments := ParseArgs(args)

		if run != true {
			t.Fatal("The first return value should be true for this test")
		}

		testDefaultsExcept(t, arguments, "mappings")

		if len(arguments.Mappings) != 1 {
			t.Fatalf("The Mappings list have length 1 (was %d)", len(arguments.Mappings))
		}

		if arguments.Mappings[0] != mapping {
			t.Errorf("First mapping was incorrect (was '%s')", arguments.Mappings[0])
		}
	})
	t.Run(mappingOption.alias, func(t *testing.T) {
		args := createArgs(mappingOption.alias, mapping, "foo.bar")
		run, arguments := ParseArgs(args)

		if run != true {
			t.Fatal("The first return value should be true for this test")
		}

		testDefaultsExcept(t, arguments, "mappings")

		if len(arguments.Mappings) != 1 {
			t.Fatalf("The Mappings list have length 1 (was %d)", len(arguments.Mappings))
		}

		if arguments.Mappings[0] != mapping {
			t.Errorf("First mapping was incorrect (was '%s')", arguments.Mappings[0])
		}
	})
	t.Run("multiple mappings", func(t *testing.T) {
		otherMapping := "cat,dog"

		args := createArgs(mappingOption.name, mapping, mappingOption.alias, otherMapping, "foo.bar")
		run, arguments := ParseArgs(args)

		if run != true {
			t.Fatal("The first return value should be true for this test")
		}

		testDefaultsExcept(t, arguments, "mappings")

		if len(arguments.Mappings) != 2 {
			t.Fatalf("The Mappings list have length 1 (was %d)", len(arguments.Mappings))
		}

		if arguments.Mappings[0] != mapping {
			t.Errorf("First mapping was incorrect (was '%s')", arguments.Mappings[0])
		}

		if arguments.Mappings[1] != otherMapping {
			t.Errorf("First mapping was incorrect (was '%s')", arguments.Mappings[0])
		}
	})
}

func TestMappingOptionIncorrect(t *testing.T) {
	t.Run("value missing", func(t *testing.T) {
		args := createArgs(mappingOption.name)
		run, _ := ParseArgs(args)

		if run != false {
			t.Error("The first return value should be false if there is an error in the args")
		}
	})
	t.Run("other flag", func(t *testing.T) {
		args := createArgs(mappingOption.name, silentFlag.name)
		run, _ := ParseArgs(args)

		if run != false {
			t.Error("The first return value should be false if there is an error in the args")
		}
	})
}

func TestUnknownOption(t *testing.T) {
	args := createArgs("--this-is-definitely-not-a-valid-option")
	run, _ := ParseArgs(args)

	if run == true {
		t.Error("The first return value should be false if there is an unknown arg")
	}
}
