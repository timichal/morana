/*
helper functions
*/

package main

import (
	"math/rand"
	"time"
)

// generating a random int [0,n)
func randomInt(upper int) int {
	randseed := rand.NewSource(time.Now().UnixNano())
	randsrc := rand.New(randseed)

	return randsrc.Intn(upper)
}
