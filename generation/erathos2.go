package generation

import (
	"math"
)

// NewErathos2 is an optimization of "erathos1", where only the integers which
// are not multiple of 2 and 3 are considered.
func NewErathos2(max uint) PrimeGenerator {
	return eratos2{max}
}

type eratos2 struct {
	max uint
}

// iToN maps an index to an integer (which is not multiple of 2 and 3)
func (eratos2) iToN(i uint) uint {
	if i%2 == 0 {
		return (i/2+1)*6 - 1
	}
	return (i/2+1)*6 + 1
}

// nToI maps the integer to its index
func (eratos2) nToI(n uint) uint {
	switch n % 6 {
	case 1:
		return (n/6)*2 - 1
	case 5:
		return (n / 6) * 2
	default:
		panic("Should not call nToI with such a number!")
	}
}

func (e eratos2) GeneratePrimes(n int) []uint {
	// Highest number in the sieve
	max := (e.max/6)*6 + 5

	// Its square root
	sq := uint(math.Sqrt(float64(max)))

	// Slice of booleans used to mark composed numbers.
	isComposed := make([]bool, e.nToI(max)+1)

	// Create the list of prime numbers
	primes := make([]uint, n)
	primes[0] = 2
	primes[1] = 3

	j := 2

	// Iterate over prime numbers until we reach the
	// square root of the highest number
	for i, c := range isComposed {
		m := e.iToN(uint(i))
		if j >= n {
			break
		}
		if !c {
			primes[j] = m
			j++
			if m <= sq {
				for k := m * m; k <= max; k = k + m {
					if k%2 != 0 && k%3 != 0 {
						isComposed[e.nToI(k)] = true
					}
				}
			}
		}
	}

	return primes[:j]
}
