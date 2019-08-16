package main

import (
	"math"
)

// First naive implementation of the Sieve of Eratosthenes

type eratos1 struct {
	max uint
}

// iToN maps an index to an integer
func (eratos1) iToN(i uint) uint {
	return i + 2
}

// nToI maps the integer to its index
func (eratos1) nToI(n uint) uint {
	return n - 2
}

func (e eratos1) GeneratePrimes(n int) []uint {
	// Highest number in the sieve
	max := e.max

	// Its square root
	sq := uint(math.Sqrt(float64(max)))

	// Slice of booleans used to mark composed numbers.
	isComposed := make([]bool, e.nToI(max)+1)

	// Iterate over prime numbers until we reach the
	// square root of the highest number
	for i, c := range isComposed {
		m := e.iToN(uint(i))
		if m > sq {
			break
		}
		if !c {
			for k := 2 * m; k <= max; k = k + m {
				isComposed[e.nToI(k)] = true
			}
		}
	}

	// Generate the list of prime numbers (= unmarked numbers)
	primes := make([]uint, n)
	i := 0
	for j, c := range isComposed {
		if i >= n {
			break
		}
		if !c {
			primes[i] = e.iToN(uint(j))
			i++
		}
	}
	return primes[:i]
}
