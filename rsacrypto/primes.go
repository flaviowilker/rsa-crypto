package rsacrypto

import (
	"math/big"

	"github.com/flaviowilker/rsa-crypto/util"
)

func newPrimes() *Primes {
	primesSlice := make([]*big.Int, 2, 2)
	primesSlice[0] = new(big.Int)
	primesSlice[1] = new(big.Int)

	for i := 0; i < 2; i++ {
		for primesSlice[i].Cmp(big.NewInt(0)) == 0 {

			number := util.RandomNumber(10, 100)

			if !util.SliceContains(primesSlice, number) && util.IsPrime(number) {
				primesSlice[i] = number
			}
		}
	}

	return &Primes{
		P: primesSlice[0],
		Q: primesSlice[1],
	}
}

// Primes ...
type Primes struct {
	P *big.Int
	Q *big.Int
}

// GetNumberN ...
func (p *Primes) GetNumberN() *big.Int {
	n := new(big.Int).Mul(p.P, p.Q)

	return n
}

// GetNumberZ ...
func (p *Primes) GetNumberZ() *big.Int {
	numA := new(big.Int).Sub(p.P, big.NewInt(1))

	numB := new(big.Int).Sub(p.Q, big.NewInt(1))

	z := new(big.Int).Mul(numA, numB)

	return z
}
