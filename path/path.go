package path

import (
	"os"
)

// IsExist 判断路径是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}

func CreatePath(path ...string) error {
	for _, p := range path {
		if !IsExist(p) {
			return os.MkdirAll(p, os.ModePerm)
		}
	}
	return nil
}
