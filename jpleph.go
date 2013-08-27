package novas

import (
	"runtime"
	"strings"
)

// This function must be in a source file that doesn't include package "C"
// or else it will give the wrong directory name.
func jpleph() string {
	_, filename, _, _ := runtime.Caller(0)
	i := strings.LastIndex(filename, "/")
	if i > 0 {
		filename = filename[:i+1]
	} else {
		filename = ""
	}
	return filename + "jpleph/JPLEPH"
}

