package gpk

import (
	"os"
	"strings"
)

func getDir(fullPath string) (dir string) {
	arr := strings.Split(fullPath, "/")
	for i := 0; i < len(arr)-1; i++ {
		dir += arr[i] + "/"
	}
	return
}

func isExist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return os.IsExist(err)
	}
	return true
}


