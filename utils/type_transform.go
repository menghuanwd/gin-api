package utils

import (
	"fmt"
	"strconv"
)

// Float64ToString ...
func Float64ToString(input float64) string {
	return strconv.FormatFloat(input, 'f', 6, 64)
}

// UintToString ...
func UintToString(input uint) string {
	return fmt.Sprintf("%v", input)
}

// StringToUint ...
func StringToUint(input string) (uint, error) {
	uint64ID, err := strconv.ParseUint(input, 10, 64)

	return uint(uint64ID), err
}

// Int32ToString ...
func Int32ToString(input int32) string {
	return fmt.Sprintf("%v", input)
}

// CamelString string, xx_yy to XxYy
func CamelString(s string) string {
	data := make([]byte, 0, len(s))
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			d = d - 32
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}
