package generation

import (
	"fmt"
	"sort"
	"testing"
)

// List of generators that should be compared
var generators = []struct {
	name      string
	max       int // maximum number of primes to generate (in tests)
	generator PrimeGenerator
}{
	{"go-probably-prime", 100_000, NewProbablyPrime()},
	{"eratos1", 100_000, NewErathos1(2_000_000)},
	{"eratos2", 100_000, NewErathos2(2_000_000)},
	{"concurent1", 10_000, NewConcurrent1()},
}

// List of numbers of primes that should be generated to test the generators
var lengths = []int{100, 1000, 10_000, 100_000}

// uintSlice attaches the methods of sort.Interface to []uint, sorting in increasing order.
type uintSlice []uint

func (p uintSlice) Len() int           { return len(p) }
func (p uintSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p uintSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func TestPrimes(t *testing.T) {
	// Generate all primes that will be tested, using "go-probably-prime"
	nmax := lengths[len(lengths)-1]
	expectedPrimes := NewProbablyPrime().GeneratePrimes(nmax)

	for _, n := range lengths {

		// we iterate over generators except the first one, because we will compare
		//each generator to the first one
		for _, gen := range generators[1:] {
			if n <= gen.max {
				t.Run(fmt.Sprintf("%v/%v", gen.name, n), func(t *testing.T) {
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
	}
}

var result []uint

func BenchmarkPrimes(b *testing.B) {
	for _, n := range lengths {
		for _, gen := range generators {
			if n <= gen.max {
				b.Run(fmt.Sprintf("%v/%v", gen.name, n), func(b *testing.B) {
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
	}
}
