package main

import (
	"fmt"
	"time"

	"github.com/olivierlemasle/prime-distribution/generation"
)

func main() {
	// number of prime numbers
	n := 100_000_004

	// size of the sieve
	max := uint(2_050_000_000)

	t1 := time.Now()
	p := generation.NewErathos2(max).GeneratePrimes(n)
	t2 := time.Now()
	fmt.Printf("Primes generation took %v\n", t2.Sub(t1))

	fmt.Printf("%v primes found\n", len(p))
	if len(p) != n {
		fmt.Printf("The sieve is not big enough to generate %v prime numbers.\n", n)
		return
	}

	t3 := time.Now()
	result := defaultDistribution(p)
	t4 := time.Now()
	fmt.Printf("Distribution took %v\n", t4.Sub(t3))
	fmt.Printf("Total: %v\n", t4.Sub(t1))

	defaultPrint(result)
}
