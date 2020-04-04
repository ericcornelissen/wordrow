package dicts

import "errors"
import "reflect"
import "testing"


func TestGetParserForMarkDownFile(t *testing.T)  {
  parserFn, err := getParserForFile("foobar.md")

  if err != nil {
    t.Fatalf("The error should be nil for this test (Error: %s)", err)
  }

  actual, expected := reflect.ValueOf(parserFn), reflect.ValueOf(parseMarkDownFile)
  if actual.Pointer() != expected.Pointer() {
    t.Error("The parser function should be the MarkDown parse function")
  }
}

func TestGetParserForCSVFile(t *testing.T)  {
  parserFn, err := getParserForFile("foobar.csv")

  if err != nil {
    t.Fatalf("The error should be nil for this test (Error: %s)", err)
  }

  actual, expected := reflect.ValueOf(parserFn), reflect.ValueOf(parseCsvFile)
  if actual.Pointer() != expected.Pointer() {
    t.Error("The parser function should be the CSV parse function")
  }
}

func TestGetParserForUnknownFileType(t *testing.T)  {
  _, err := getParserForFile("foobar")

  if err == nil {
    t.Error("The error should be set for unknown file types")
  }
}

func TestParseFileCallsParseFunction(t *testing.T) {
  var fileContent string
  var wordmap WordMap

  calledFlag := false
  parseFn := func(fc *string) (wm WordMap, err error) {
    if fc != &fileContent {
      t.Error("The argument given to the parse function should be fileContent")
    }

    calledFlag = true
    return
  }

  err := parseFile(&fileContent, parseFn, &wordmap)

  if err != nil {
    t.Fatalf("The error should be nil for this test (Error: %s)", err)
  }

  if calledFlag == false {
    t.Error("The parse function should have been called")
  }
}

func TestParseFileUpdatesWordMap(t *testing.T) {
  var fileContent string
  var wordmap WordMap

  parseFn := func(_ *string) (wm WordMap, err error) {
    wm.AddOne("a", "b")
    return
  }

  if wordmap.Size() != 0 {
    t.Fatalf("The WordMap should have size 0 before parsing (was %d)", wordmap.Size())
  }

  err := parseFile(&fileContent, parseFn, &wordmap)

  if err != nil {
    t.Fatalf("The error should be nil for this test (Error: %s)", err)
  }

  if wordmap.Size() != 1 {
    t.Error("The WordMap should have size 1 after parsing")
  }
}

func TestParseFileParseFunctionError(t *testing.T) {
  var fileContent string
  var wordmap WordMap

  expectedError := errors.New("dummy error")
  parseFn := func(_ *string) (WordMap, error) {
    return wordmap, expectedError
  }

  err := parseFile(&fileContent, parseFn, &wordmap)

  if err == nil {
    t.Fatal("The error should be set for this test")
  }

  if err != expectedError {
    t.Error("The wrong error was returned")
  }
}

func TestWordMapFromNoInput(t *testing.T) {
  wordmap, err := WordMapFrom()

  if err != nil {
    t.Error("The error should be set if there are no inputs")
  }

  if wordmap.Size() != 0 {
    t.Error("The WordMap should have size 0 if there are no inputs")
  }
}
