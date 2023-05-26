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

// CreatePath 创建路径
func CreatePath(path ...string) error {
	for _, p := range path {
		if !IsExist(p) {
			if err := os.MkdirAll(p, os.ModePerm); err != nil {
				return err
			}
		}
	}
	return nil
}
