package dir

import (
	"os"
	"strings"
)

func GetAllFilesInDirWithSuffix(dirPath string, suffix string, fileList []string) ([]string, error) {
	files, err := os.ReadDir(dirPath)
	for _, file := range files {
		if file.IsDir() {
			fileList, err = GetAllFilesInDirWithSuffix(dirPath+"/"+file.Name(), suffix, fileList)
			if err != nil {
				return nil, err
			}
		} else {
			if strings.HasSuffix(file.Name(), suffix) {
				fileList = append(fileList, dirPath+"/"+file.Name())
			}
		}
	}
	return fileList, err
}
