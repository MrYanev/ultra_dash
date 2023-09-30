package internal

import (
	"crypto/rand"
	"math/big"
)

// A func to return a random number from 0 to n
func GetRandomInt(num int) int {
	x, _ := rand.Int(rand.Reader, big.NewInt(int64(num)))
	return int(x.Int64())
}

// GetDiceRoll returns an integer from 1 to the number from the number
func GetDiceRoll(num int) int {
	x, _ := rand.Int(rand.Reader, big.NewInt(int64(num)))
	return int(x.Int64()) + 1
}

// Returns a number between the two numbers inclusive.
func GetRandomBetween(low int, high int) int {
	var randy int = -1
	for {
		randy = GetDiceRoll(high)
		if randy >= low {
			break
		}
	}
	return randy
}
