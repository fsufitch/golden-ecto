package util

import (
	"path/filepath"
)

var resourceDir string

func SetResourceDir(newdir string) {
	resourceDir = newdir
}

func GetResourcePath(path ...string) string {
	fullPath := append([]string{resourceDir}, path...)
	return filepath.Join(fullPath...)
}
