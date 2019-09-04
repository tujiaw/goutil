package goutil

import (
	"crypto/rand"
	"crypto/sha1"
	"fmt"
)

func Sub(str string, from int, to int) string {
	array := []rune(str)
	return string(array[from:to])
}

func SubFrom(str string, from int) string {
	array := []rune(str)
	return string(array[from:])
}

func SubTo(str string, to int) string {
	array := []rune(str)
	return string(array[0:to])
}

func Sha1(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func Uuidv4() (string, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		return "", err
	}
	return fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:]), nil
}

func FormatBytes(size float64) string {
	const KB = float64(1024)
	const MB = float64(KB * KB)
	const GB = float64(MB * MB)

	trimZero := func(str string) string {
		for i := len(str) - 1; i >= 0; i-- {
			if str[i] != 48 && str[i] != 46 {
				return str[:i+1]
			}
		}
		return str
	}

	switch {
	case size < KB:
		return "1 KB"
	case size >= KB && size < MB:
		return trimZero(fmt.Sprintf("%.2f", size/KB)) + " KB"
	case size >= MB && size < GB:
		return trimZero(fmt.Sprintf("%.2f", size/MB)) + " MB"
	default:
		return trimZero(fmt.Sprintf("%.2f", size/GB)) + " GB"
	}
}
