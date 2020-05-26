package fs

import (
	"fmt"
	"runtime"
)

// The value of `runtime.GOOS` on Windows.
const windows = `windows`

// GetAnAbsolutePathFor returns an arbitrary OS-dependent absolute file path for
// a given file name.
func getAnAbsolutePathFor(filename string) string {
	if runtime.GOOS == windows {
		return fmt.Sprintf(`c:\%s`, filename)
	}

	return fmt.Sprintf(`/usr/aang/%s`, filename)
}
