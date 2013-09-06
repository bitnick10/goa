package char

import (
	"math/rand"
	"time"
)

func IsAlphanumeric(b byte) bool {
	if '0' <= b && b <= '9' {
		return true
	}
	if 'A' <= b && b <= 'Z' {
		return true
	}
	if 'a' <= b && b <= 'z' {
		return true
	}
	return false
}

func RandNum() byte {
	rand.Seed(time.Now().UTC().UnixNano())
	// [0,9)
	b := byte(rand.Intn(9)) + 48
	return b
}
func RandUpperCase() byte {
	rand.Seed(time.Now().UTC().UnixNano())
	// [0,26)
	b := byte(rand.Intn(26)) + 65
	return b
}
func RandLowerCase() byte {
	rand.Seed(time.Now().UTC().UnixNano())
	// [0,26)
	b := byte(rand.Intn(26)) + 97
	return b
}
