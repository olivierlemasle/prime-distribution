package main

import (
	"fmt"
	"time"
)

func main() {
	n := 100000004 // number of prime numbers

	t1 := time.Now()
	e := eratos2{
		max: 2350000000, // size of the sieve
	}
	p := e.GeneratePrimes(n)
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
