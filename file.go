package goutil

import (
	"fmt"
	"os"
)

func WriteFileAppend(filename string, fileBytes []byte) error {
	fi, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("open file error, %v", err)
	}
	defer fi.Close()
	_, err = fi.Write(fileBytes)
	return err
}

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func GetFileSize(filename string) int64 {
	fi, err := os.Stat(filename)
	if err != nil {
		return -1
	}
	return fi.Size()
}
