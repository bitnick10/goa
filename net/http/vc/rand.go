package vc

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/Bitnick2002/goa/char"
)

func rd() int {
	rand.Seed(time.Now().UTC().UnixNano())
	// returns a non-negative pseudo-random int
	return rand.Int()
}
func randString() string {
	return strconv.Itoa(rd())
}
func randCode() string {
	rand.Seed(time.Now().UTC().UnixNano())
	var bytes [4]byte
	for i := 0; i < 4; {
		b := byte(rand.Intn(75))
		b += 48
		if char.IsAlphanumeric(b) {
			i++
		}
	}
	return string(bytes[:])
}

// func randAlphaNumric() string {
// 	for {
// 		rand.Seed(time.Now().UTC().UnixNano())
// 		// [0,75)
// 		b := byte(rand.Intn(75))
// 		if isAlphaNumeric(b) {
// 			return string(b)
// 		}
// 	}
// }
func randNum() string {
	rand.Seed(time.Now().UTC().UnixNano())
	// [0,9)
	b := byte(rand.Intn(9)) + 48
	return string(b)
}
func randUpperCase() string {
	rand.Seed(time.Now().UTC().UnixNano())
	// [0,26)
	b := byte(rand.Intn(26)) + 65
	return string(b)
}
func randLowerCase() string {
	rand.Seed(time.Now().UTC().UnixNano())
	// [0,26)
	b := byte(rand.Intn(26)) + 97
	return string(b)
}
