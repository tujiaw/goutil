package goutil

import (
	"testing"
)

func TestFormatBytes(t *testing.T) {
	data := map[float64]string{
		34:               "1 KB",
		1024:             "1 KB",
		1043:             "1.02 KB",
		2048:             "2 KB",
		1024 * 1024:      "1 MB",
		1024*1024 + 8243: "1.01 MB",
		37191648:         "35.47 MB",
	}

	for k, v := range data {
		if FormatBytes(k) != v {
			t.Error("ERROR", k, v, FormatBytes(k))
			break
		}
	}
}
