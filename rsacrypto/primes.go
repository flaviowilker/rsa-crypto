package rsacrypto

import (
	"math/big"

	"github.com/flaviowilker/rsa-crypto/util"
)

// ToChosePrimes ...
func ToChosePrimes() *Primes {
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
func (p *Primes) GetNumberN() *big.Int {
	n := p.P * p.Q

	return big.NewInt(int64(n))
}

// GetNumberZ ...
func (p *Primes) GetNumberZ() *big.Int {
	z := (p.P - 1) * (p.Q - 1)

	return big.NewInt(int64(z))
}
