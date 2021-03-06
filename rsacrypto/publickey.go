package rsacrypto

import (
	"math/big"

	"github.com/flaviowilker/rsa-crypto/util"
)

func newPublicKey(primes *Primes) *PublicKey {
	n := primes.GetNumberN()
	z := primes.GetNumberZ()
	e := new(big.Int)

	for e.Cmp(big.NewInt(0)) == 0 {
		randomE := util.RandomNumber(10, 200)

		if util.IsMdcPrimes(z, randomE) {
			e = randomE
		}
	}

	return &PublicKey{
		N: n,
		E: e,
	}
}

// PublicKey ...
type PublicKey struct {
	N *big.Int
	E *big.Int
}
