package goutil

import "testing"

type item struct {
	target interface{}
	search []interface{}
	result []bool
}

func TestContain(t *testing.T) {
	sliceData := item{
		target: []int{1, 3, 45, 32, 4, 356, 436},
		search: []interface{}{1, 2, 3, 4},
		result: []bool{true, false, true, true},
	}
	for i, s := range sliceData.search {
		if Contain(sliceData.target, s) != sliceData.result[i] {
			t.Error("index", i, "error")
		}
	}
}
