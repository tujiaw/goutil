package goutil

import (
	"reflect"
	"strings"
)

func SliceDelete(slice interface{}, index int) (interface{}, bool) {
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil, false
	}
	if v.Len() == 0 || index < 0 || index > v.Len()-1 {
		return nil, false
	}
	return reflect.AppendSlice(v.Slice(0, index), v.Slice(index+1, v.Len())).Interface(), true
}

func UniqueSliceString(s []string, caseSensitive bool) []string {
	result := make([]string, 0, len(s))
	mp := make(map[string]int)
	for _, v := range s {
		key := ""
		if caseSensitive {
			key = v
		} else {
			key = strings.ToLower(v)
		}
		if mp[key] == 0 {
			result = append(result, v)
		}
		mp[key] += 1
	}
	return result
}

func IndexOf(arr []string, str string) int {
	for i, v := range arr {
		if v == str {
			return i
		}
	}
	return -1
}

func Contain(target interface{}, search interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == search {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(search)).IsValid() {
			return true
		}
	}
	return false
}
