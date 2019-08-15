package main

// PrimeGenerator generates the list of n first prime numbers
type PrimeGenerator interface {
	GeneratePrimes(n int) []uint
}
