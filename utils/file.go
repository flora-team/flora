package utils

import "os"

func IsPathExist(path string) bool {
	_, err := os.Lstat(path)
	return !os.IsNotExist(err)
}
