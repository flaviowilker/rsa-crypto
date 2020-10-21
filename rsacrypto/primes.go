package rsacrypto

import (
	"github.com/flaviowilker/rsa-crypto/util"
)

// NewPrimes ...
func NewPrimes() *Primes {
	primesSlice := make([]int, 2)

	for i := 0; i < 2; i++ {
		for primesSlice[i] == 0 {
			number := util.RandomNumber(100, 200)

			if !util.SliceContains(primesSlice, number) && util.IsPrime(number) {
				primesSlice[i] = number
			}
		}
	}

	return &Primes{primesSlice[0], primesSlice[1]}
}

// Primes ...
type Primes struct {
	P int
	Q int
}

// GetNumberN ...
func (p *Primes) GetNumberN() int {
	n := p.P * p.Q
	return n
}

// GetNumberZ ...
func (p *Primes) GetNumberZ() int {
	z := (p.P - 1) * (p.Q - 1)
	return z
}
