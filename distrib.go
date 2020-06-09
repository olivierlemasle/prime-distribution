package main

import "fmt"

type distributionParams struct {
	modulo  uint
	classes []int
}

var defaultParams = distributionParams{
	modulo:  uint(10),
	classes: []int{1, 3, 7, 9},
}

func defaultDistribution(primes []uint) map[int]int {
	return distribution(primes, defaultParams)
}

func distribution(primes []uint, p distributionParams) map[int]int {
	result := make(map[int]int)
	l := len(primes)
	for i := 0; i < l-1; i++ {
		prime1 := primes[i]
		prime2 := primes[i+1]
		k := int(p.modulo*(prime1%p.modulo) + prime2%p.modulo)
		result[k]++
	}
	return result
}

func defaultPrint(distrib map[int]int) {
	print(distrib, defaultParams)
}

func print(distrib map[int]int, p distributionParams) {
	fmt.Println()
	for _, k1 := range p.classes {
		for _, k2 := range p.classes {
			k := int(p.modulo)*k1 + k2
			v := distrib[k]
			fmt.Printf("%v: %v\n", []int{k1, k2}, v)
		}
		fmt.Println()
	}
}
