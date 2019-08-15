package main

import "fmt"

type args struct {
	modulo  uint
	classes []int
}

var defaultArgs = args{
	modulo:  uint(10),
	classes: []int{1, 3, 7, 9},
}

func defaultDistribution(primes []uint) map[int]int {
	return distribution(primes, defaultArgs)
}

func distribution(primes []uint, a args) map[int]int {
	result := make(map[int]int)
	l := len(primes)
	for i := 0; i < l-1; i++ {
		p1 := primes[i]
		p2 := primes[i+1]
		k := int(a.modulo*(p1%a.modulo) + p2%a.modulo)
		result[k]++
	}
	return result
}

func defaultPrint(distrib map[int]int) {
	print(distrib, defaultArgs)
}

func print(distrib map[int]int, a args) {
	fmt.Println()
	for _, k1 := range a.classes {
		for _, k2 := range a.classes {
			k := int(a.modulo)*k1 + k2
			v := distrib[k]
			fmt.Printf("%v: %v\n", []int{k1, k2}, v)
		}
		fmt.Println()
	}
}
