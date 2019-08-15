package main

// This is an optimization of "erathos1", where only the integers which
// are not multiple of 2 and 3 are considered.

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
	if n%6 == 5 {
		return (n / 6) * 2
	} else if n%6 == 1 {
		return (n/6)*2 - 1
	}
	panic("Should not call nToI with such a number!")
}

func (e eratos2) GeneratePrimes(n int) []uint {
	// Highest number in the sieve
	max := (e.max/6)*6 + 5

	// Slice of booleans used to mark composed numbers.
	isComposed := make([]bool, e.nToI(max)+1)

	// Iterate over prime numbers until we reach the
	// square root of the highest number
	for i, c := range isComposed {
		m := e.iToN(uint(i))
		if !c {
			for k := m * m; k <= max; k = k + m {
				if k%2 != 0 && k%3 != 0 {
					isComposed[e.nToI(k)] = true
				}
			}
		}
	}

	// Generate the list of prime numbers (= unmarked numbers)
	primes := make([]uint, n)
	primes[0] = 2
	primes[1] = 3
	i := 2
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
