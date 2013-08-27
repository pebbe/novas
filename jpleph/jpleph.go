/*
Location of file JPLEPH.
*/
package jpleph

import (
	"runtime"
	"strings"
)

/*
This assumes the file JPLEPH was installed in the same directory as the current sourcefile.
*/
func JPLEPH() string {
	_, filename, _, _ := runtime.Caller(0)
	i := strings.LastIndex(filename, "/")
	if i > 0 {
		filename = filename[:i+1]
	} else {
		filename = ""
	}
	return filename + "JPLEPH"
}
