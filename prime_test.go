package main

import (
	"math/big"
	"sort"
	"testing"
)

// The most naive prime generator

type goProbablyPrime bool

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

func NewGoProbablyPrime() PrimeGenerator {
	return goProbablyPrime(true)
}

// List of generators that should be compared
var generators = []struct {
	name      string
	generator PrimeGenerator
}{
	{"go-probably-prime", NewGoProbablyPrime()},
	{"eratos1", eratos1{2000000}},
	{"eratos2", eratos2{2000000}},
}

// uintSlice attaches the methods of sort.Interface to []uint, sorting in increasing order.
type uintSlice []uint

func (p uintSlice) Len() int           { return len(p) }
func (p uintSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p uintSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func TestPrimes(t *testing.T) {
	n := 100000

	expectedPrimes := NewGoProbablyPrime().GeneratePrimes(n)

	for _, gen := range generators[1:] { //  first element removed
		t.Run(gen.name, func(t *testing.T) {
			primes := gen.generator.GeneratePrimes(n)

			if len(primes) != n {
				t.Errorf("Generating primes with %v did not provide %v prime numbers, but %v.", gen.name, n, len(primes))
			}

			if !sort.IsSorted(uintSlice(primes)) {
				t.Errorf("List of primes generated with %v is not ordered ascending.", gen.name)
			}

			//Compare "primes" and "expectedPrimes", which should both be ordered ascending
			for i, j := 0, 0; i < len(primes) && j < len(expectedPrimes); {
				if primes[i] == expectedPrimes[j] {
					i++
					j++
				} else if primes[i] > expectedPrimes[j] {
					t.Errorf("%v is missing in prime numbers list.", expectedPrimes[j])
					j++
				} else {
					t.Errorf("%v should not be in prime numbers list (index %v).", primes[i], i)
					i++
				}
			}
		})
	}
}

var result []uint

func BenchmarkPrimes(b *testing.B) {
	n := 100000

	for _, gen := range generators {
		b.Run(gen.name, func(b *testing.B) {
			var r []uint
			for i := 0; i < b.N; i++ {
				// The result is recorded to prevent the compilator to eliminate the function call.
				r = gen.generator.GeneratePrimes(n)
			}

			// This assignation is just to prevent the compilator to eliminate the benchmark itself.
			result = r
		})
	}
}
