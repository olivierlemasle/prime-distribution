package generation

// NewConcurrent1 creates a prime generator with a concurrent implementation,
// inspired by Golang example (https://play.golang.org/p/iN6HCp_e91p).
// This is elegant but highly inefficient.
func NewConcurrent1() PrimeGenerator {
	return conc1{}
}

type conc1 struct{}

func (conc1) GeneratePrimes(n int) []uint {
	ch := make(chan uint)
	res := make([]uint, n)

	go generate(ch)

	for i := 0; i < n; i++ {
		prime := <-ch
		res[i] = prime
		ch1 := make(chan uint)
		go filter(ch, ch1, prime)
		ch = ch1
	}

	return res
}

// generate sends all numbers >= 2 to a channel ch: 2, 3, 4, 5, etc.
func generate(ch chan<- uint) {
	for i := uint(2); ; i++ {
		ch <- i
	}
}

// filter copies the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func filter(in <-chan uint, out chan<- uint, prime uint) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}
