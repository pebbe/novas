package novas

import (
	"runtime"
	"strings"
)

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

