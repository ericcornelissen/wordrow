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
	if arguments.DryRun == true {
		t.Error("The default value for the DryRun option should be false")
	}
}

func testDefaultInvert(t *testing.T, arguments Arguments) {
	if arguments.Invert == true {
		t.Error("The default value for the Invert option should be false")
	}
}

func testDefaultSilent(t *testing.T, arguments Arguments) {
	if arguments.Silent == true {
		t.Error("The default value for the Silent option should be false")
	}
}

func testDefaultVerbose(t *testing.T, arguments Arguments) {
	if arguments.Verbose == true {
		t.Error("The default value for the Silent option should be false")
	}
}

func testDefaultConfigFile(t *testing.T, arguments Arguments) {
	if arguments.ConfigFile != "" {
		t.Error("The default ConfigFile should not be set")
	}
}

func testDefaultMapFiles(t *testing.T, arguments Arguments) {
	if len(arguments.MapFiles) != 0 {
		t.Error("The default list of MapFiles should be empty")
	}
}

func testDefaultMappings(t *testing.T, arguments Arguments) {
	if len(arguments.Mappings) != 0 {
		t.Error("The default list of Mappings should be empty")
	}
}

func testDefaultsExcept(t *testing.T, arguments Arguments, exclude string) {
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

func TestHelpArg(t *testing.T) {
	args := createArgs(helpFlag)
	run, _ := ParseArgs(args)

	if run == true {
		t.Error("The first return value should be false for --help")
	}
}

func TestVersionArg(t *testing.T) {
	t.Run("--version only", func(t *testing.T) {
		args := createArgs(versionFlag)
		run, _ := ParseArgs(args)

		if run == true {
			t.Fatal("The first return value should be false if --version is the only argument")
		}
	})
	t.Run("--version and other", func(t *testing.T) {
		args := createArgs(versionFlag, "foo.bar")
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
	args := createArgs(dryRunFlag, "foo.bar")
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
	t.Run(invertFlag, func(t *testing.T) {
		args := createArgs(invertFlag, "foo.bar")
		run, arguments := ParseArgs(args)

		if run != true {
			t.Fatal("The first return value should be true for this test")
		}

		testDefaultsExcept(t, arguments, "invert")

		if arguments.Invert != true {
			t.Errorf("The Invert value should be true if %s is an argument", invertFlag)
		}
	})
	t.Run(invertFlagAlias, func(t *testing.T) {
		args := createArgs(invertFlagAlias, "foo.bar")
		run, arguments := ParseArgs(args)

		if run != true {
			t.Fatal("The first return value should be true for this test")
		}

		testDefaultsExcept(t, arguments, "invert")

		if arguments.Invert != true {
			t.Errorf("The Invert value should be true if %s is an argument", invertFlagAlias)
		}
	})
}

func TestSilentFlag(t *testing.T) {
	t.Run(silentFlag, func(t *testing.T) {
		args := createArgs(silentFlag, "foo.bar")
		run, arguments := ParseArgs(args)

		if run != true {
			t.Fatal("The first return value should be true for this test")
		}

		testDefaultsExcept(t, arguments, "silent")

		if arguments.Silent != true {
			t.Errorf("The Silent value should be true if %s is an argument", silentFlag)
		}
	})
	t.Run(silentFlagAlias, func(t *testing.T) {
		args := createArgs(silentFlagAlias, "foo.bar")
		run, arguments := ParseArgs(args)

		if run != true {
			t.Fatal("The first return value should be true for this test")
		}

		testDefaultsExcept(t, arguments, "silent")

		if arguments.Silent != true {
			t.Errorf("The Silent value should be true if %s is an argument", silentFlagAlias)
		}
	})
}

func TestVerboseFlag(t *testing.T) {
	t.Run(verboseFlag, func(t *testing.T) {
		args := createArgs(verboseFlag, "foo.bar")
		run, arguments := ParseArgs(args)

		if run != true {
			t.Fatal("The first return value should be true for this test")
		}

		testDefaultsExcept(t, arguments, "verbose")

		if arguments.Verbose != true {
			t.Errorf("The Verbose value should be true if %s is an argument", verboseFlag)
		}
	})
	t.Run(verboseFlagAlias, func(t *testing.T) {
		args := createArgs(verboseFlagAlias, "foo.bar")
		run, arguments := ParseArgs(args)

		if run != true {
			t.Fatal("The first return value should be true for this test")
		}

		testDefaultsExcept(t, arguments, "verbose")

		if arguments.Verbose != true {
			t.Errorf("The Verbose value should be true if %s is an argument", verboseFlagAlias)
		}
	})
}

func TestConfigFileOption(t *testing.T) {
	configFile := "config.json"

	t.Run(configOption, func(t *testing.T) {
		args := createArgs(configOption, configFile, "foo.bar")
		run, arguments := ParseArgs(args)

		if run != true {
			t.Fatal("The first return value should be true for this test")
		}

		testDefaultsExcept(t, arguments, "config file")

		if arguments.ConfigFile != configFile {
			t.Errorf("The config file was incorrect (was '%s')", arguments.ConfigFile)
		}
	})
	t.Run(configOptionAlias, func(t *testing.T) {
		args := createArgs(configOptionAlias, configFile, "foo.bar")
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

		args := createArgs(configOption, configFile, configOption, otherConfigFile, "foo.bar")
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
		args := createArgs(configOption)
		run, _ := ParseArgs(args)

		if run != false {
			t.Error("The first return value should be false if there is an error in the args")
		}
	})
	t.Run("other flag", func(t *testing.T) {
		args := createArgs(configOption, silentFlag)
		run, _ := ParseArgs(args)

		if run != false {
			t.Error("The first return value should be false if there is an error in the args")
		}
	})
}

func TestMapFileOption(t *testing.T) {
	mapFile := "foo.map"

	t.Run(mapfileOption, func(t *testing.T) {
		args := createArgs(mapfileOption, mapFile, "foo.bar")
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
	t.Run(mapfileOptionAlias, func(t *testing.T) {
		args := createArgs(mapfileOptionAlias, mapFile, "foo.bar")
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

		args := createArgs(mapfileOption, mapFile, mapfileOption, otherMapFile, "foo.bar")
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
		args := createArgs(mapfileOption)
		run, _ := ParseArgs(args)

		if run != false {
			t.Error("The first return value should be false if there is an error in the args")
		}
	})
	t.Run("other flag", func(t *testing.T) {
		args := createArgs(mapfileOption, silentFlag)
		run, _ := ParseArgs(args)

		if run != false {
			t.Error("The first return value should be false if there is an error in the args")
		}
	})
}

func TestMappingOption(t *testing.T) {
	mapping := "foo,bar"

	t.Run(mappingOption, func(t *testing.T) {
		args := createArgs(mappingOption, mapping, "foo.bar")
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
	t.Run(mappingOptionAlias, func(t *testing.T) {
		args := createArgs(mappingOptionAlias, mapping, "foo.bar")
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

		args := createArgs(mappingOption, mapping, mappingOptionAlias, otherMapping, "foo.bar")
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
		args := createArgs(mappingOption)
		run, _ := ParseArgs(args)

		if run != false {
			t.Error("The first return value should be false if there is an error in the args")
		}
	})
	t.Run("other flag", func(t *testing.T) {
		args := createArgs(mappingOption, silentFlag)
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
