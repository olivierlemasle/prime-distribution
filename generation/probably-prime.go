package generation

import (
	"math/big"
)

// NewProbablyPrime creates a prime generator using
// Int.ProbablyPrime to test each integer independently.
func NewProbablyPrime() PrimeGenerator {
	return goProbablyPrime{}
}

type goProbablyPrime struct{}

func (goProbablyPrime) GeneratePrimes(n int) []uint {
	result := make([]uint, n)
	c := 0
	for i := uint(0); c < n; i++ {
		if big.NewInt(int64(i)).ProbablyPrime(0) {
			result[c] = i
			c++
		}
	}
	return result
}
