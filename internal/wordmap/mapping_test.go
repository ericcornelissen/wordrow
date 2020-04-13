package wordmap

import "fmt"
import "testing"


func ExampleMappingMatch() {
  s := "Hello world!"
  m := Mapping{
    newFromValue("hello"),
    newToValue("hey"),
  }

  for match := range m.Match(s) {
    fmt.Println(match.Word)
    // Output: Hello
  }
}


func TestFromValue(t *testing.T)  {
  baseValue := "foo"
  t.Run("without prefix or suffix", func(t *testing.T) {
    rawValue := baseValue
    from := newFromValue(rawValue)

    if from.Value != baseValue {
      t.Errorf("Unexpected value (got '%s')", from.Value)
    }

    if from.String() != rawValue {
      t.Errorf("Unexpected String value (got '%s')", from.String())
    }

    if from.IncludePrefix != false {
      t.Error("Prefix should not be included")
    }

    if from.IncludeSuffix != false {
      t.Error("Suffix should not be included")
    }
  })
  t.Run("with prefix and without suffix", func(t *testing.T) {
    rawValue := fmt.Sprintf("-%s", baseValue)
    from := newFromValue(rawValue)

    if from.Value != baseValue {
      t.Errorf("Unexpected value (got '%s')", from.Value)
    }

    if from.String() != rawValue {
      t.Errorf("Unexpected String value (got '%s')", from.String())
    }

    if from.IncludePrefix != true {
      t.Error("Prefix should be included")
    }

    if from.IncludeSuffix != false {
      t.Error("Suffix should not be included")
    }
  })
  t.Run("without prefix and with suffix", func(t *testing.T) {
    rawValue := fmt.Sprintf("%s-", baseValue)
    from := newFromValue(rawValue)

    if from.Value != baseValue {
      t.Errorf("Unexpected value (got '%s')", from.Value)
    }

    if from.String() != rawValue {
      t.Errorf("Unexpected String value (got '%s')", from.String())
    }

    if from.IncludePrefix != false {
      t.Error("Prefix should not be included")
    }

    if from.IncludeSuffix != true {
      t.Error("Suffix should be included")
    }
  })
  t.Run("with prefix and suffix", func(t *testing.T) {
    rawValue := fmt.Sprintf("-%s-", baseValue)
    from := newFromValue(rawValue)

    if from.Value != baseValue {
      t.Errorf("Unexpected value (got '%s')", from.Value)
    }

    if from.String() != rawValue {
      t.Errorf("Unexpected String value (got '%s')", from.String())
    }

    if from.IncludePrefix != true {
      t.Error("Prefix should be included")
    }

    if from.IncludeSuffix != true {
      t.Error("Suffix should be included")
    }
  })
}

func TestToValue(t *testing.T)  {
  baseValue := "bar"
  t.Run("without prefix or suffix", func(t *testing.T) {
    rawValue := baseValue
    from := newToValue(rawValue)

    if from.Value != baseValue {
      t.Errorf("Unexpected value (got '%s')", from.Value)
    }

    if from.String() != rawValue {
      t.Errorf("Unexpected String value (got '%s')", from.String())
    }

    if from.KeepPrefix != false {
      t.Error("Prefix should not be kept")
    }

    if from.KeepSuffix != false {
      t.Error("Suffix should not be kept")
    }
  })
  t.Run("with prefix and without suffix", func(t *testing.T) {
    rawValue := fmt.Sprintf("-%s", baseValue)
    from := newToValue(rawValue)

    if from.Value != baseValue {
      t.Errorf("Unexpected value (got '%s')", from.Value)
    }

    if from.String() != rawValue {
      t.Errorf("Unexpected String value (got '%s')", from.String())
    }

    if from.KeepPrefix != true {
      t.Error("Prefix should be kept")
    }

    if from.KeepSuffix != false {
      t.Error("Suffix should not be kept")
    }
  })
  t.Run("without prefix and with suffix", func(t *testing.T) {
    rawValue := fmt.Sprintf("%s-", baseValue)
    from := newToValue(rawValue)

    if from.Value != baseValue {
      t.Errorf("Unexpected value (got '%s')", from.Value)
    }

    if from.String() != rawValue {
      t.Errorf("Unexpected String value (got '%s')", from.String())
    }

    if from.KeepPrefix != false {
      t.Error("Prefix should not be kept")
    }

    if from.KeepSuffix != true {
      t.Error("Suffix should be kept")
    }
  })
  t.Run("with prefix and suffix", func(t *testing.T) {
    rawValue := fmt.Sprintf("-%s-", baseValue)
    from := newToValue(rawValue)

    if from.Value != baseValue {
      t.Errorf("Unexpected value (got '%s')", from.Value)
    }

    if from.String() != rawValue {
      t.Errorf("Unexpected String value (got '%s')", from.String())
    }

    if from.KeepPrefix != true {
      t.Error("Prefix should be kept")
    }

    if from.KeepSuffix != true {
      t.Error("Suffix should be kept")
    }
  })
}

func TestMappingNoPrefixNoSuffix(t *testing.T) {
  from, to := "hello", "hey"
  m := Mapping{
    newFromValue(from),
    newToValue(to),
  }

  t.Run("GetReplacement", func(t *testing.T) {
    t.Run("Empty prefix, empty suffix", func(t *testing.T) {
      result := m.GetReplacement("", "")

      if result != to {
        t.Errorf("Unexpected replacement (got '%s')", result)
      }
    })
    t.Run("Non-empty prefix, empty suffix", func(t *testing.T) {
      result := m.GetReplacement("howdy", "")

      if result != to {
        t.Errorf("Unexpected replacement (got '%s')", result)
      }
    })
    t.Run("Empty prefix, non-empty suffix", func(t *testing.T) {
      result := m.GetReplacement("", "yo")

      if result != to {
        t.Errorf("Unexpected replacement (got '%s')", result)
      }
    })
    t.Run("Non-mpty prefix, non-empty suffix", func(t *testing.T) {
      result := m.GetReplacement("howdy", "yo")

      if result != to {
        t.Errorf("Unexpected replacement (got '%s')", result)
      }
    })
  })
  t.Run("Match", func(t *testing.T) {
    t.Run("Empty input string", func(t *testing.T) {
      for match := range m.Match("") {
        t.Errorf("There shouldn't be any matches (got %+v)", match)
      }
    })
    t.Run("No matches", func(t *testing.T) {
      source := "This string should not contain the from"
      for match := range m.Match(source) {
        t.Errorf("There shouldn't be any matches (got %+v)", match)
      }
    })
    t.Run("No prefix and no suffix", func(t *testing.T) {
      rawSource := "%s there! %s, how are you?"
      source := fmt.Sprintf(rawSource, from, from)

      expectedMatches := []Match{
        Match{
          Full: from,
          Word: from,
          Prefix: "",
          Suffix: "",
          Start: 0,
          End: len(from),
        },
        Match{
          Full: from,
          Word: from,
          Prefix: "",
          Suffix: "",
          Start: len(from) + 8,
          End: len(from) + 8 + len(from),
        },
      }

      i := 0
      for match := range m.Match(source) {
        if i >= len(expectedMatches) {
          t.Fatal("Too many matches found")
        }

        if match != expectedMatches[i] {
          t.Errorf("Unexpected match at index %d (was %+v)", i, match)
        }

        i++
      }

      if i != len(expectedMatches) {
        t.Errorf("not enough matches (got %d)", i)
      }
    })
    t.Run("With prefix and no suffix", func(t *testing.T) {
      rawSource := "%s there! foo%s, how are you?"
      source := fmt.Sprintf(rawSource, from, from)

      expectedMatches := []Match{
        Match{
          Full: from,
          Word: from,
          Prefix: "",
          Suffix: "",
          Start: 0,
          End: len(from),
        },
      }

      i := 0
      for match := range m.Match(source) {
        if i >= len(expectedMatches) {
          t.Fatal("Too many matches found")
        }

        if match != expectedMatches[i] {
          t.Errorf("Unexpected match at index %d (was %+v)", i, match)
        }

        i++
      }

      if i != len(expectedMatches) {
        t.Errorf("not enough matches (got %d)", i)
      }
    })
    t.Run("No prefix and with suffix", func(t *testing.T) {
      rawSource := "%sbar there! %s, how are you?"
      source := fmt.Sprintf(rawSource, from, from)

      expectedMatches := []Match{
        Match{
          Full: from,
          Word: from,
          Prefix: "",
          Suffix: "",
          Start: len(from) + 11,
          End: len(from) + 11 + len(from),
        },
      }

      i := 0
      for match := range m.Match(source) {
        if i >= len(expectedMatches) {
          t.Fatal("Too many matches found")
        }

        if match != expectedMatches[i] {
          t.Errorf("Unexpected match at index %d (was %+v)", i, match)
        }

        i++
      }

      if i != len(expectedMatches) {
        t.Errorf("not enough matches (got %d)", i)
      }
    })
    t.Run("With prefix and with suffix", func(t *testing.T) {
      rawSource := "foo%s there! %sbar, how are you?"
      source := fmt.Sprintf(rawSource, from, from)

      for match := range m.Match(source) {
        t.Errorf("There shouldn't be any matches (got %+v)", match)
      }
    })
  })
  t.Run("String", func(t *testing.T) {
    result := m.String()

    if result != "[hello -> hey]" {
      t.Errorf("Unexpected String value (was '%s')", result)
    }
  })
}

func TestMappingWithPrefixNoSuffix(t *testing.T) {
  from, to := "bar", "foo"
  m := Mapping{
    newFromValue("-" + from),
    newToValue("-" + to),
  }

  t.Run("GetReplacement", func(t *testing.T) {
    t.Run("Empty prefix, empty suffix", func(t *testing.T) {
      result := m.GetReplacement("", "")

      if result != to {
        t.Errorf("Unexpected replacement (got '%s')", result)
      }
    })
    t.Run("Non-empty prefix, empty suffix", func(t *testing.T) {
      result := m.GetReplacement("dead", "")

      if result != "dead" + to {
        t.Errorf("Unexpected replacement (got '%s')", result)
      }
    })
    t.Run("Empty prefix, non-empty suffix", func(t *testing.T) {
      result := m.GetReplacement("", "beef")

      if result != to {
        t.Errorf("Unexpected replacement (got '%s')", result)
      }
    })
    t.Run("Non-mpty prefix, non-empty suffix", func(t *testing.T) {
      result := m.GetReplacement("dead", "beef")

      if result != "dead" + to {
        t.Errorf("Unexpected replacement (got '%s')", result)
      }
    })
  })
  t.Run("Match", func(t *testing.T) {
    t.Run("Empty input string", func(t *testing.T) {
      for match := range m.Match("") {
        t.Errorf("There shouldn't be any matches (got %+v)", match)
      }
    })
    t.Run("No matches", func(t *testing.T) {
      source := "This string should not contain the from"
      for match := range m.Match(source) {
        t.Errorf("There shouldn't be any matches (got %+v)", match)
      }
    })
    t.Run("No prefix and no suffix", func(t *testing.T) {
      rawSource := "Here is a %s and there is another %s"
      source := fmt.Sprintf(rawSource, from, from)

      expectedMatches := []Match{
        Match{
          Full: from,
          Word: from,
          Prefix: "",
          Suffix: "",
          Start: 10,
          End: 10 + len(from),
        },
        Match{
          Full: from,
          Word: from,
          Prefix: "",
          Suffix: "",
          Start: len(from) + 32,
          End: len(from) + 32 + len(from),
        },
      }

      i := 0
      for match := range m.Match(source) {
        if i >= len(expectedMatches) {
          t.Fatal("Too many matches found")
        }

        if match != expectedMatches[i] {
          t.Errorf("Unexpected match at index %d (was %+v)", i, match)
        }

        i++
      }

      if i != len(expectedMatches) {
        t.Errorf("not enough matches (got %d)", i)
      }
    })
    t.Run("With prefix and no suffix", func(t *testing.T) {
      rawSource := "Here is a %s and there is another pre%s"
      source := fmt.Sprintf(rawSource, from, from)

      expectedMatches := []Match{
        Match{
          Full: from,
          Word: from,
          Prefix: "",
          Suffix: "",
          Start: 10,
          End: 10 + len(from),
        },
        Match{
          Full: "pre" + from,
          Word: from,
          Prefix: "pre",
          Suffix: "",
          Start: len(from) + 32,
          End: len(from) + 35 + len(from),
        },
      }

      i := 0
      for match := range m.Match(source) {
        if i >= len(expectedMatches) {
          t.Fatal("Too many matches found")
        }

        if match != expectedMatches[i] {
          t.Errorf("Unexpected match at index %d (was %+v)", i, match)
        }

        i++
      }

      if i != len(expectedMatches) {
        t.Errorf("not enough matches (got %d)", i)
      }
    })
    t.Run("No prefix and with suffix", func(t *testing.T) {
      rawSource := "Here is a %ssuf and there is another %s"
      source := fmt.Sprintf(rawSource, from, from)

      expectedMatches := []Match{
        Match{
          Full: from,
          Word: from,
          Prefix: "",
          Suffix: "",
          Start: len(from) + 35,
          End: len(from) + 35 + len(from),
        },
      }

      i := 0
      for match := range m.Match(source) {
        if i >= len(expectedMatches) {
          t.Fatal("Too many matches found")
        }

        if match != expectedMatches[i] {
          t.Errorf("Unexpected match at index %d (was %+v)", i, match)
        }

        i++
      }

      if i != len(expectedMatches) {
        t.Errorf("not enough matches (got %d)", i)
      }
    })
    t.Run("With prefix and with suffix", func(t *testing.T) {
      rawSource := "Here is a pre%s and there is another %ssuf"
      source := fmt.Sprintf(rawSource, from, from)

      expectedMatches := []Match{
        Match{
          Full: "pre" + from,
          Word: from,
          Prefix: "pre",
          Suffix: "",
          Start: 10,
          End: 13 + len(from),
        },
      }

      i := 0
      for match := range m.Match(source) {
        if i >= len(expectedMatches) {
          t.Fatal("Too many matches found")
        }

        if match != expectedMatches[i] {
          t.Errorf("Unexpected match at index %d (was %+v)", i, match)
        }

        i++
      }

      if i != len(expectedMatches) {
        t.Errorf("not enough matches (got %d)", i)
      }
    })
  })
  t.Run("String", func(t *testing.T) {
    result := m.String()

    if result != "[-bar -> -foo]" {
      t.Errorf("Unexpected String value (was '%s')", result)
    }
  })
}

func TestMappingNoPrefixWithSuffix(t *testing.T) {
  from, to := "foo", "bar"
  m := Mapping{
    newFromValue(from + "-"),
    newToValue(to + "-"),
  }

  t.Run("GetReplacement", func(t *testing.T) {
    t.Run("Empty prefix, empty suffix", func(t *testing.T) {
      result := m.GetReplacement("", "")

      if result != to {
        t.Errorf("Unexpected replacement (got '%s')", result)
      }
    })
    t.Run("Non-empty prefix, empty suffix", func(t *testing.T) {
      result := m.GetReplacement("high", "")

      if result != to {
        t.Errorf("Unexpected replacement (got '%s')", result)
      }
    })
    t.Run("Empty prefix, non-empty suffix", func(t *testing.T) {
      result := m.GetReplacement("", "dragun")

      if result != to + "dragun" {
        t.Errorf("Unexpected replacement (got '%s')", result)
      }
    })
    t.Run("Non-mpty prefix, non-empty suffix", func(t *testing.T) {
      result := m.GetReplacement("high", "dragun")

      if result != to + "dragun" {
        t.Errorf("Unexpected replacement (got '%s')", result)
      }
    })
  })
  t.Run("Match", func(t *testing.T) {
    t.Run("Empty input string", func(t *testing.T) {
      for match := range m.Match("") {
        t.Errorf("There shouldn't be any matches (got %+v)", match)
      }
    })
    t.Run("No matches", func(t *testing.T) {
      source := "This string should not contain the from"
      for match := range m.Match(source) {
        t.Errorf("There shouldn't be any matches (got %+v)", match)
      }
    })
    t.Run("No prefix and no suffix", func(t *testing.T) {
      rawSource := "Here is a %s and there is another %s"
      source := fmt.Sprintf(rawSource, from, from)

      expectedMatches := []Match{
        Match{
          Full: from,
          Word: from,
          Prefix: "",
          Suffix: "",
          Start: 10,
          End: 10 + len(from),
        },
        Match{
          Full: from,
          Word: from,
          Prefix: "",
          Suffix: "",
          Start: len(from) + 32,
          End: len(from) + 32 + len(from),
        },
      }

      i := 0
      for match := range m.Match(source) {
        if i >= len(expectedMatches) {
          t.Fatal("Too many matches found")
        }

        if match != expectedMatches[i] {
          t.Errorf("Unexpected match at index %d (was %+v)", i, match)
        }

        i++
      }

      if i != len(expectedMatches) {
        t.Errorf("not enough matches (got %d)", i)
      }
    })
    t.Run("With prefix and no suffix", func(t *testing.T) {
      rawSource := "Here is a %s and there is another pre%s"
      source := fmt.Sprintf(rawSource, from, from)

      expectedMatches := []Match{
        Match{
          Full: from,
          Word: from,
          Prefix: "",
          Suffix: "",
          Start: 10,
          End: 10 + len(from),
        },
      }

      i := 0
      for match := range m.Match(source) {
        if i >= len(expectedMatches) {
          t.Fatal("Too many matches found")
        }

        if match != expectedMatches[i] {
          t.Errorf("Unexpected match at index %d (was %+v)", i, match)
        }

        i++
      }

      if i != len(expectedMatches) {
        t.Errorf("not enough matches (got %d)", i)
      }
    })
    t.Run("No prefix and with suffix", func(t *testing.T) {
      rawSource := "Here is a %ssuf and there is another %s"
      source := fmt.Sprintf(rawSource, from, from)

      expectedMatches := []Match{
        Match{
          Full: from + "suf",
          Word: from,
          Prefix: "",
          Suffix: "suf",
          Start: 10,
          End: 13 + len(from),
        },
        Match{
          Full: from,
          Word: from,
          Prefix: "",
          Suffix: "",
          Start: len(from) + 35,
          End: len(from) + 35 + len(from),
        },
      }

      i := 0
      for match := range m.Match(source) {
        if i >= len(expectedMatches) {
          t.Fatal("Too many matches found")
        }

        if match != expectedMatches[i] {
          t.Errorf("Unexpected match at index %d (was %+v)", i, match)
        }

        i++
      }

      if i != len(expectedMatches) {
        t.Errorf("not enough matches (got %d)", i)
      }
    })
    t.Run("With prefix and with suffix", func(t *testing.T) {
      rawSource := "Here is a pre%s and there is another %ssuf"
      source := fmt.Sprintf(rawSource, from, from)

      expectedMatches := []Match{
        Match{
          Full: from + "suf",
          Word: from,
          Prefix: "",
          Suffix: "suf",
          Start: 35 + len(from),
          End: 38 + len(from) + len(from),
        },
      }

      i := 0
      for match := range m.Match(source) {
        if i >= len(expectedMatches) {
          t.Fatal("Too many matches found")
        }

        if match != expectedMatches[i] {
          t.Errorf("Unexpected match at index %d (was %+v)", i, match)
        }

        i++
      }

      if i != len(expectedMatches) {
        t.Errorf("not enough matches (got %d)", i)
      }
    })
  })
  t.Run("String", func(t *testing.T) {
    result := m.String()

    if result != "[foo- -> bar-]" {
      t.Errorf("Unexpected String value (was '%s')", result)
    }
  })
}

func TestMappingWithPrefixAndSuffix(t *testing.T) {
  from, to := "foobar", "lorem"
  m := Mapping{
    newFromValue("-" + from + "-"),
    newToValue("-" + to + "-"),
  }

  t.Run("GetReplacement", func(t *testing.T) {
    t.Run("Empty prefix, empty suffix", func(t *testing.T) {
      result := m.GetReplacement("", "")

      if result != to {
        t.Errorf("Unexpected replacement (got '%s')", result)
      }
    })
    t.Run("Non-empty prefix, empty suffix", func(t *testing.T) {
      result := m.GetReplacement("praise", "")

      if result != "praise" + to {
        t.Errorf("Unexpected replacement (got '%s')", result)
      }
    })
    t.Run("Empty prefix, non-empty suffix", func(t *testing.T) {
      result := m.GetReplacement("", "sun")

      if result != to + "sun" {
        t.Errorf("Unexpected replacement (got '%s')", result)
      }
    })
    t.Run("Non-mpty prefix, non-empty suffix", func(t *testing.T) {
      result := m.GetReplacement("praise", "sun")

      if result != "praise" + to + "sun" {
        t.Errorf("Unexpected replacement (got '%s')", result)
      }
    })
  })
  t.Run("Match", func(t *testing.T) {
    t.Run("Empty input string", func(t *testing.T) {
      for match := range m.Match("") {
        t.Errorf("There shouldn't be any matches (got %+v)", match)
      }
    })
    t.Run("No matches", func(t *testing.T) {
      source := "This string should not contain the from"
      for match := range m.Match(source) {
        t.Errorf("There shouldn't be any matches (got %+v)", match)
      }
    })
    t.Run("No prefix and no suffix", func(t *testing.T) {
      rawSource := "Here is a %s and there is another %s"
      source := fmt.Sprintf(rawSource, from, from)

      expectedMatches := []Match{
        Match{
          Full: from,
          Word: from,
          Prefix: "",
          Suffix: "",
          Start: 10,
          End: 10 + len(from),
        },
        Match{
          Full: from,
          Word: from,
          Prefix: "",
          Suffix: "",
          Start: len(from) + 32,
          End: len(from) + 32 + len(from),
        },
      }

      i := 0
      for match := range m.Match(source) {
        if i >= len(expectedMatches) {
          t.Fatal("Too many matches found")
        }

        if match != expectedMatches[i] {
          t.Errorf("Unexpected match at index %d (was %+v)", i, match)
        }

        i++
      }

      if i != len(expectedMatches) {
        t.Errorf("not enough matches (got %d)", i)
      }
    })
    t.Run("With prefix and no suffix", func(t *testing.T) {
      rawSource := "Here is a %s and there is another pre%s"
      source := fmt.Sprintf(rawSource, from, from)

      expectedMatches := []Match{
        Match{
          Full: from,
          Word: from,
          Prefix: "",
          Suffix: "",
          Start: 10,
          End: 10 + len(from),
        },
        Match{
          Full: "pre" + from,
          Word: from,
          Prefix: "pre",
          Suffix: "",
          Start: len(from) + 32,
          End: len(from) + 35 + len(from),
        },
      }

      i := 0
      for match := range m.Match(source) {
        if i >= len(expectedMatches) {
          t.Fatal("Too many matches found")
        }

        if match != expectedMatches[i] {
          t.Errorf("Unexpected match at index %d (was %+v)", i, match)
        }

        i++
      }

      if i != len(expectedMatches) {
        t.Errorf("not enough matches (got %d)", i)
      }
    })
    t.Run("No prefix and with suffix", func(t *testing.T) {
      rawSource := "Here is a %ssuf and there is another %s"
      source := fmt.Sprintf(rawSource, from, from)

      expectedMatches := []Match{
        Match{
          Full: from + "suf",
          Word: from,
          Prefix: "",
          Suffix: "suf",
          Start: 10,
          End: 13 + len(from),
        },
        Match{
          Full: from,
          Word: from,
          Prefix: "",
          Suffix: "",
          Start: len(from) + 35,
          End: len(from) + 35 + len(from),
        },
      }

      i := 0
      for match := range m.Match(source) {
        if i >= len(expectedMatches) {
          t.Fatal("Too many matches found")
        }

        if match != expectedMatches[i] {
          t.Errorf("Unexpected match at index %d (was %+v)", i, match)
        }

        i++
      }

      if i != len(expectedMatches) {
        t.Errorf("not enough matches (got %d)", i)
      }
    })
    t.Run("With prefix and with suffix", func(t *testing.T) {
      rawSource := "Here is a pre%s and there is another %ssuf"
      source := fmt.Sprintf(rawSource, from, from)

      expectedMatches := []Match{
        Match{
          Full: "pre" + from,
          Word: from,
          Prefix: "pre",
          Suffix: "",
          Start: 10,
          End: 13 + len(from),
        },
        Match{
          Full: from + "suf",
          Word: from,
          Prefix: "",
          Suffix: "suf",
          Start: 35 + len(from),
          End: 38 + len(from) + len(from),
        },
      }

      i := 0
      for match := range m.Match(source) {
        if i >= len(expectedMatches) {
          t.Fatal("Too many matches found")
        }

        if match != expectedMatches[i] {
          t.Errorf("Unexpected match at index %d (was %+v)", i, match)
        }

        i++
      }

      if i != len(expectedMatches) {
        t.Errorf("not enough matches (got %d)", i)
      }
    })
  })
  t.Run("String", func(t *testing.T) {
    result := m.String()

    if result != "[-foobar- -> -lorem-]" {
      t.Errorf("Unexpected String value (was '%s')", result)
    }
  })
}
