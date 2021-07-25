package utils

import (
	"errors"
	"reflect"
	"sort"
	"strings"
)

// Contains 判断obj是否在target中，target支持的类型arrary,slice,map
func Contains(obj interface{}, target interface{}) (bool, error) {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true, nil
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true, nil
		}
	}

	return false, errors.New("not in array")
}

// AddSingleQuoteToArray ...
func AddSingleQuoteToArray(array []string) (newArray []string) {
	for _, item := range array {
		newArray = append(newArray, "E'"+item+"'")
	}
	return
}

// HasOne ...
func HasOne(array []string, tmp string) bool {
	for _, item := range array {
		if item == tmp {
			return true
		}
	}

	return false
}

// String2Array ...
func String2Array(str string) []string {
	return strings.Split(str, ",")
}

// SliceRemoveDuplicates ...
func SliceRemoveDuplicates(slice []string) []string {
	sort.Strings(slice)
	i := 0
	var j int
	for {
		if i >= len(slice)-1 {
			break
		}
		for j = i + 1; j < len(slice) && slice[i] == slice[j]; j++ {
		}
		slice = append(slice[:i+1], slice[j:]...)
		i++
	}
	return slice
}

// RemoveSameFromAnotherArray ...
func RemoveSameFromAnotherArray(arrayA, arrayB []string) (diff []string) {
	for _, i := range arrayB {
		if !HasOne(arrayA, i) {
			diff = append(diff, i)
		}
	}
	return
}

// DeleteArrayEmpty ...
func DeleteArrayEmpty(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}

	return RemoveRepByMap(r)
}

// RemoveRepByMap ...
func RemoveRepByMap(slc []string) []string {
	result := []string{}
	tempMap := map[string]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result
}
