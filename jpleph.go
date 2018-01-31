package novas

import (
	"os"
	"path/filepath"
	"runtime"
)

var (
	JPLephFile = ""
)

// This function must be in a source file that doesn't include package "C"
// or else it will give the wrong directory name.
func jpleph() string {

	if jpl := os.Getenv("JPLEPH"); jpl != "" {
		JPLephFile = jpl
	}

	if JPLephFile != "" {
		return JPLephFile
	}

	_, filename, _, _ := runtime.Caller(0)
	JPLephFile = filepath.Join(filepath.Dir(filename), "jpleph", "JPLEPH")
	return JPLephFile

}
