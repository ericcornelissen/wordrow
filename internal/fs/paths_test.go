package fs

import "testing"

func TestGetExt(t *testing.T) {
	path := "foo.bar"

	result := GetExt(path)

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
