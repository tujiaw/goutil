package goutil

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	_           = iota             // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota) // 1 << (10*1)
	MB                             // 1 << (10*2)
	GB                             // 1 << (10*3)
	TB                             // 1 << (10*4)
	PB                             // 1 << (10*5)
	EB                             // 1 << (10*6)
	ZB                             // 1 << (10*7)
	YB                             // 1 << (10*8)
)

type ByteSize float64

func (b ByteSize) Format() string {
	if b >= GB {
		return fmt.Sprintf("%.2f GB", b/GB)
	} else if b >= MB {
		return fmt.Sprintf("%.2f MB", b/MB)
	} else if b >= KB {
		return fmt.Sprintf("%.2f KB", b/KB)
	} else {
		return fmt.Sprintf("%v B", b)
	}
}

func If(yes bool, left interface{}, right interface{}) interface{} {
	if yes {
		return left
	}
	return right
}

func IfBool(yes bool, left bool, right bool) bool {
	if yes {
		return left
	}
	return right
}

func RandPick(total int, pick int) []int {
	if total < pick {
		return []int{}
	}

	start := time.Now().UnixNano()
	mp := map[int]int{}
	for {
		start = start + 1
		rand.Seed(start)
		n := rand.Intn(total)
		mp[n]++
		if len(mp) == pick {
			break
		}
	}

	result := make([]int, 0, len(mp))
	for k, _ := range mp {
		result = append(result, k)
	}

	return result
}
