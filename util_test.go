package goutil

import (
	"testing"
)

func TestFormatBytes(t *testing.T) {
	data := map[ByteSize]string{
		34:               "34 B",
		1024:             "1.00 KB",
		1043:             "1.02 KB",
		2048:             "2.00 KB",
		1024 * 1024:      "1.00 MB",
		1024*1024 + 8243: "1.01 MB",
		37191648:         "35.47 MB",
	}

	for k, v := range data {
		if k.Format() != v {
			t.Error("ERROR", k, v, k.Format())
			break
		}
	}
}
