package rsacrypto

import (
	"fmt"
	"math/big"

	"github.com/flaviowilker/rsa-crypto/util"
)

func newPrimes() *Primes {
	primesSlice := make([]*big.Int, 2, 2)
	primesSlice[0] = new(big.Int)
	primesSlice[1] = new(big.Int)

	for i := 0; i < 2; i++ {
		for primesSlice[i].Cmp(big.NewInt(0)) == 0 {

			number := util.RandomNumber(2, 100)

			if !util.SliceContains(primesSlice, number) && util.IsPrime(number) {
				primesSlice[i] = number
			}
		}
	}

	fmt.Println("p: ", primesSlice[0])
	fmt.Println("q: ", primesSlice[1])

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

func (p *Primes) getNumberN() *big.Int {
	n := new(big.Int).Mul(p.P, p.Q)

	return n
}

func (p *Primes) getNumberZ() *big.Int {
	numA := new(big.Int).Sub(p.P, big.NewInt(1))

	numB := new(big.Int).Sub(p.Q, big.NewInt(1))

	z := new(big.Int).Mul(numA, numB)

	return z
}
