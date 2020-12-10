package goutil

import (
	"fmt"
	"io"
	"io/ioutil"
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

func FileExists(path string) bool {
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

func DeleteFile(path string) error {
	f, err := os.Stat(path)
	if err != nil {
		return err
	}

	if !f.IsDir() {
		return os.Remove(path)
	}
	return nil
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

func IsPathEmpty(name string) bool {
	f, err := os.Open(name)
	if err != nil {
		return false
	}
	defer f.Close()

	_, err = f.Readdirnames(1)
	if err == io.EOF {
		return true
	}
	return false
}

//计算文件MD5值
func MD5File(file string) (string, error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	return MD5Bytes(data), nil
}
