package replacer

import "testing"


func TestStartsWithCapital(t *testing.T) {
  t.Run("does not start with capital", func(t *testing.T) {
    result := startsWithCapital("foobar")

    if result != false {
      t.Error("Expected result to be false")
    }
  })
  t.Run("starts with capital", func(t *testing.T) {
    result := startsWithCapital("Foobar")

    if result != true {
      t.Error("Expected result to be true")
    }
  })
}

func TestToSentenceCase(t *testing.T) {
  t.Run("one word", func(t *testing.T) {
    result := toSentenceCase("foobar")

    if result != "Foobar" {
      t.Errorf("Unexpected result (got '%s')", result)
    }
  })
  t.Run("uncapatalized sentence", func(t *testing.T) {
    result := toSentenceCase("hello world!")

    if result != "Hello world!" {
      t.Errorf("Unexpected result (got '%s')", result)
    }
  })
  t.Run("capatalized sentence", func(t *testing.T) {
    result := toSentenceCase("Hello world!")

    if result != "Hello world!" {
      t.Errorf("Unexpected result (got '%s')", result)
    }
  })
}

func TestMaintainAllCaps(t *testing.T) {
  t.Run("no caps", func(t *testing.T) {
    result := maintainAllCaps("foo", "bar")

    if result != "bar" {
      t.Errorf("Unexpected result (got '%s')", result)
    }
  })
  t.Run("some caps", func(t *testing.T) {
    result := maintainAllCaps("Foo", "bar")

    if result != "bar" {
      t.Errorf("Unexpected result (got '%s')", result)
    }
  })
  t.Run("all caps", func(t *testing.T) {
    result := maintainAllCaps("FOO", "bar")

    if result != "BAR" {
      t.Errorf("Unexpected result (got '%s')", result)
    }
  })
}

func TestMaintainCapitalization(t *testing.T) {
  t.Run("no capitalization in word", func(t *testing.T) {
    result := maintainCapitalization("foo", "bar")

    if result != "bar" {
      t.Errorf("Unexpected result (got '%s')", result)
    }
  })
  t.Run("capitalization in word", func(t *testing.T) {
    result := maintainCapitalization("Foo", "bar")

    if result != "Bar" {
      t.Errorf("Unexpected result (got '%s')", result)
    }
  })
  t.Run("capitalization in phrase", func(t *testing.T) {
    result := maintainCapitalization("Foo Bar", "bar foo")

    if result != "Bar Foo" {
      t.Errorf("Unexpected result (got '%s')", result)
    }
  })
  t.Run("capitalization in hyphenated phrase", func(t *testing.T) {
    result := maintainCapitalization("So Called", "so-called")

    if result != "So-Called" {
      t.Errorf("Unexpected result (got '%s')", result)
    }
  })
  t.Run("partly capitalized phrase", func(t *testing.T) {
    result := maintainCapitalization("Foo bar", "bar foo")

    if result != "Bar foo" {
      t.Errorf("Unexpected result (got '%s')", result)
    }

    result = maintainCapitalization("foo Bar", "bar foo")

    if result != "bar Foo" {
      t.Errorf("Unexpected result (got '%s')", result)
    }
  })
  t.Run("capitalized phrases, long to short", func(t *testing.T) {
    result := maintainCapitalization("Lorem Ipsum Dolor", "sit amet")

    if result != "Sit Amet" {
      t.Errorf("Unexpected result (got '%s')", result)
    }

    result = maintainCapitalization("Lorem ipsum Dolor", "sit amet")

    if result != "Sit amet" {
      t.Errorf("Unexpected result (got '%s')", result)
    }

    result = maintainCapitalization("lorem Ipsum Dolor", "sit amet")

    if result != "sit Amet" {
      t.Errorf("Unexpected result (got '%s')", result)
    }
  })
  t.Run("capitalized phrases, short to long", func(t *testing.T) {
    result := maintainCapitalization("Lorem Ipsum", "dolor sit amet")

    if result != "Dolor Sit amet" {
      t.Errorf("Unexpected result (got '%s')", result)
    }

    result = maintainCapitalization("Lorem ipsum", "dolor sit amet")

    if result != "Dolor sit amet" {
      t.Errorf("Unexpected result (got '%s')", result)
    }

    result = maintainCapitalization("lorem Ipsum", "dolor sit amet")

    if result != "dolor Sit amet" {
      t.Errorf("Unexpected result (got '%s')", result)
    }
  })
}

func TestMaintainWhitespace(t *testing.T) {
  // TODO
}
