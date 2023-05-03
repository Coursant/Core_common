package Core_common

import "math/rand"

// ReadNumber create random number
func ReadNumber() int {
	// random number range
	rnr := 10
	return rand.Intn(rnr)
	//
}
