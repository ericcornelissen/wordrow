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


func ExampleParseArgs() {
  args := []string {"wordrow", "foo.bar"}
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
  args := createArgs(helpOption)
  run, _ := ParseArgs(args)

  if run == true {
    t.Error("The first return value should be false for --help")
  }
}

func TestVersionArg(t *testing.T) {
  t.Run("--version only", func(t *testing.T) {
    args := createArgs(versionOption)
    run, _ := ParseArgs(args)

    if run == true {
      t.Fatal("The first return value should be false if --version is the only argument")
    }
  })
  t.Run("--version and other", func(t *testing.T) {
    args := createArgs(versionOption, "foo.bar")
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

  testDefaultDryRun(t, arguments)
  testDefaultInvert(t, arguments)
  testDefaultSilent(t, arguments)
  testDefaultVerbose(t, arguments)
  testDefaultConfigFile(t, arguments)
  testDefaultMapFiles(t, arguments)

  if len(arguments.InputFiles) != 1 {
    t.Error("The list of InputFiles should contain a single file")
  }
}

func TestDryRunFlag(t *testing.T) {
  args := createArgs(dryRunOption, "foo.bar")
  run, arguments := ParseArgs(args)

  if run != true {
    t.Fatal("The first return value should be true for this test")
  }

  testDefaultInvert(t, arguments)
  testDefaultSilent(t, arguments)
  testDefaultVerbose(t, arguments)
  testDefaultConfigFile(t, arguments)
  testDefaultMapFiles(t, arguments)

  if arguments.DryRun != true {
    t.Errorf("The DryRun value should be true if %s is an argument", dryRunOption)
  }
}

func TestInvertFlag(t *testing.T) {
  t.Run(invertOption, func(t *testing.T) {
    args := createArgs(invertOption, "foo.bar")
    run, arguments := ParseArgs(args)

    if run != true {
      t.Fatal("The first return value should be true for this test")
    }

    testDefaultDryRun(t, arguments)
    testDefaultSilent(t, arguments)
    testDefaultVerbose(t, arguments)
    testDefaultConfigFile(t, arguments)
    testDefaultMapFiles(t, arguments)

    if arguments.Invert != true {
      t.Errorf("The Invert value should be true if %s is an argument", invertOption)
    }
  })
  t.Run(invertOptionAlias, func(t *testing.T) {
    args := createArgs(invertOptionAlias, "foo.bar")
    run, arguments := ParseArgs(args)

    if run != true {
      t.Fatal("The first return value should be true for this test")
    }

    testDefaultDryRun(t, arguments)
    testDefaultSilent(t, arguments)
    testDefaultVerbose(t, arguments)
    testDefaultConfigFile(t, arguments)
    testDefaultMapFiles(t, arguments)

    if arguments.Invert != true {
      t.Errorf("The Invert value should be true if %s is an argument", invertOptionAlias)
    }
  })
}

func TestSilentFlag(t *testing.T) {
  t.Run(silentOption, func(t *testing.T) {
    args := createArgs(silentOption, "foo.bar")
    run, arguments := ParseArgs(args)

    if run != true {
      t.Fatal("The first return value should be true for this test")
    }

    testDefaultDryRun(t, arguments)
    testDefaultInvert(t, arguments)
    testDefaultVerbose(t, arguments)
    testDefaultConfigFile(t, arguments)
    testDefaultMapFiles(t, arguments)

    if arguments.Silent != true {
      t.Errorf("The Silent value should be true if %s is an argument", silentOption)
    }
  })
  t.Run(silentOptionAlias, func(t *testing.T) {
    args := createArgs(silentOptionAlias, "foo.bar")
    run, arguments := ParseArgs(args)

    if run != true {
      t.Fatal("The first return value should be true for this test")
    }

    testDefaultDryRun(t, arguments)
    testDefaultInvert(t, arguments)
    testDefaultVerbose(t, arguments)
    testDefaultConfigFile(t, arguments)
    testDefaultMapFiles(t, arguments)

    if arguments.Silent != true {
      t.Errorf("The Silent value should be true if %s is an argument", silentOptionAlias)
    }
  })
}

func TestVerboseFlag(t *testing.T) {
  t.Run(verboseOption, func(t *testing.T) {
    args := createArgs(verboseOption, "foo.bar")
    run, arguments := ParseArgs(args)

    if run != true {
      t.Fatal("The first return value should be true for this test")
    }

    testDefaultDryRun(t, arguments)
    testDefaultInvert(t, arguments)
    testDefaultSilent(t, arguments)
    testDefaultConfigFile(t, arguments)
    testDefaultMapFiles(t, arguments)

    if arguments.Verbose != true {
      t.Errorf("The Verbose value should be true if %s is an argument", verboseOption)
    }
  })
  t.Run(verboseOptionAlias, func(t *testing.T) {
    args := createArgs(verboseOptionAlias, "foo.bar")
    run, arguments := ParseArgs(args)

    if run != true {
      t.Fatal("The first return value should be true for this test")
    }

    testDefaultDryRun(t, arguments)
    testDefaultInvert(t, arguments)
    testDefaultSilent(t, arguments)
    testDefaultConfigFile(t, arguments)
    testDefaultMapFiles(t, arguments)

    if arguments.Verbose != true {
      t.Errorf("The Verbose value should be true if %s is an argument", verboseOptionAlias)
    }
  })
}

func TestConfigOption(t *testing.T) {
  configFile := "config.json"

  t.Run(configOption, func(t *testing.T) {
    args := createArgs(configOption, configFile, "foo.bar")
    run, arguments := ParseArgs(args)

    if run != true {
      t.Fatal("The first return value should be true for this test")
    }

    testDefaultDryRun(t, arguments)
    testDefaultInvert(t, arguments)
    testDefaultSilent(t, arguments)
    testDefaultVerbose(t, arguments)
    testDefaultMapFiles(t, arguments)

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

    testDefaultDryRun(t, arguments)
    testDefaultInvert(t, arguments)
    testDefaultSilent(t, arguments)
    testDefaultVerbose(t, arguments)
    testDefaultMapFiles(t, arguments)

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

    testDefaultDryRun(t, arguments)
    testDefaultInvert(t, arguments)
    testDefaultSilent(t, arguments)
    testDefaultVerbose(t, arguments)
    testDefaultMapFiles(t, arguments)

    if arguments.ConfigFile != otherConfigFile {
      t.Errorf("The config file was incorrect (was '%s')", arguments.ConfigFile)
    }
  })
}

func TestConfigOptionIncorrect(t *testing.T) {
  t.Run("value missing", func(t *testing.T) {
    args := createArgs(configOption)
    run, _ := ParseArgs(args)

    if run != false {
      t.Error("The first return value should be false if there is an error in the args")
    }
  })
  t.Run("other flag", func(t *testing.T) {
    args := createArgs(configOption, silentOption)
    run, _ := ParseArgs(args)

    if run != false {
      t.Error("The first return value should be false if there is an error in the args")
    }
  })
}

func TestMappingOption(t *testing.T) {
  mapFile := "foo.map"

  t.Run(mapfileOption, func(t *testing.T) {
    args := createArgs(mapfileOption, mapFile, "foo.bar")
    run, arguments := ParseArgs(args)

    if run != true {
      t.Fatal("The first return value should be true for this test")
    }

    testDefaultDryRun(t, arguments)
    testDefaultInvert(t, arguments)
    testDefaultSilent(t, arguments)
    testDefaultVerbose(t, arguments)
    testDefaultConfigFile(t, arguments)

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

    testDefaultDryRun(t, arguments)
    testDefaultInvert(t, arguments)
    testDefaultSilent(t, arguments)
    testDefaultVerbose(t, arguments)
    testDefaultConfigFile(t, arguments)

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

    testDefaultDryRun(t, arguments)
    testDefaultInvert(t, arguments)
    testDefaultSilent(t, arguments)
    testDefaultVerbose(t, arguments)
    testDefaultConfigFile(t, arguments)

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

func TestMappingOptionIncorrect(t *testing.T) {
  t.Run("value missing", func(t *testing.T) {
    args := createArgs(mapfileOption)
    run, _ := ParseArgs(args)

    if run != false {
      t.Error("The first return value should be false if there is an error in the args")
    }
  })
  t.Run("other flag", func(t *testing.T) {
    args := createArgs(mapfileOption, silentOption)
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
