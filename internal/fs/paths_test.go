package fs

import "runtime"
import "testing"


func getAnAbsolutePathFor(file string) string {
  if runtime.GOOS == "windows" {
    return "C:\\" + file
  }

  return "/usr/aang/" + file
}


func TestGetExt(t *testing.T) {
  path := "foo.bar"

  result := getExt(path)

  if result != ".bar" {
    t.Errorf("Unexpected file extension (got '%s')", result)
  }
}

func TestResolveGlobsNoGlobs(t *testing.T) {
  resolvedPaths, err := ResolveGlobs()

  if err != nil {
    t.Error("Resolving zero globs should not set the error")
  }

  if len(resolvedPaths) != 0 {
    t.Error("Resolving zero globs should result in an empty slice")
  }
}

func TestResolveGlobsWithoutGlobs(t *testing.T) {
  path0 := getAnAbsolutePathFor("foo.bar")
  path1 := "./hello.world"
  resolvedPaths, err := ResolveGlobs(path0, path1)

  if err != nil {
    t.Error("Resolving legitimate paths should not set the error")
  }

  if len(resolvedPaths) != 2 {
    t.Errorf("Resolving two non-glob paths should return two paths (was %d)", len(resolvedPaths))
  }
}

func TestResolveGlobsWithGlobs(t *testing.T) {
  resolvedPaths, err := ResolveGlobs("./*")

  if err != nil {
    t.Error("Resolving legitimate globs should not set the error")
  }

  if len(resolvedPaths) < 1 {
    t.Errorf("Resolving a glob should return at least one path (was %d)", len(resolvedPaths))
  }
}

func TestResolveGlobsWithMalformedPattern(t *testing.T) {
  _, err := ResolveGlobs("[")

  if err == nil {
    t.Error("The error should be set for a malformed glob")
  }
}

func TestResolveNoPaths(t *testing.T) {
  resolvedPaths := ResolvePaths()

  if len(resolvedPaths) != 0 {
    t.Error("Resolving no paths should result in an empty slice")
  }
}

func TestResolveRelativePaths(t *testing.T) {
  path := "./foo.bar"
  resolvedPaths := ResolvePaths(path)

  if len(resolvedPaths) != 1 {
    t.Error("Resolving one path should result in a single path")
  }

  if resolvedPaths[0] == path {
    t.Errorf("The resolved path should be different from the original ('%s')", resolvedPaths[0])
  }
}

func TestResolveAbsolutePaths(t *testing.T) {
  path := getAnAbsolutePathFor("foo.bar")
  resolvedPaths := ResolvePaths(path)

  if len(resolvedPaths) != 1 {
    t.Error("Resolving one path should result in a single path")
  }

  if resolvedPaths[0] != path {
    t.Errorf("The resolved path should be the same as the original ('%s')", resolvedPaths[0])
  }
}

func TestResolveAbsoluteAndRelativePaths(t *testing.T) {
  path0 := getAnAbsolutePathFor("foo.bar")
  path1 := "./hello.world"
  resolvedPaths := ResolvePaths(path0, path1)

  if len(resolvedPaths) != 2 {
    t.Error("Resolving one path should result in a single path")
  }

  if resolvedPaths[0] != path0 {
    t.Errorf("The resolved path should be the same as the original ('%s')", resolvedPaths[0])
  }

  if resolvedPaths[1] == path1 {
    t.Errorf("The resolved path should be different from the original ('%s')", resolvedPaths[1])
  }
}
