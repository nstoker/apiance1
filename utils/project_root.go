package utils

import (
	"path/filepath"
	"runtime"
	"strings"
)

var (
	_, b, _, _ = runtime.Caller(0)
	basepath   = filepath.Dir(b)
)

// GetProjectRoot will return the root path of the project
func GetProjectRoot() string {

	return strings.Replace(basepath, "/utils", "", -1)
}
