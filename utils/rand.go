package utils

import (
	"math/rand"
	"strconv"
	"time"
)

// GenOrderNo ...
func GenOrderNo() string {
	timestamp := time.Now().Unix()

	return strconv.Itoa(int(timestamp)) + randSeq(4)
}

// Timestamp ...
func Timestamp() string {
	return strconv.Itoa(int(time.Now().Unix()))
}

// GenSMSCode ...
func GenSMSCode() string {
	return randSeq(4)
}

func randSeq(n int) string {
	var letters = []rune("1234567890")
	b := make([]rune, n)
	for i := range b {
		rand.Seed(time.Now().UnixNano())
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
