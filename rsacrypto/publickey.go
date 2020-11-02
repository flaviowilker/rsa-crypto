package rsacrypto

import (
	"math/big"

	"github.com/flaviowilker/rsa-crypto/util"
)

func newPublicKey(primes *Primes) *PublicKey {
	n := primes.getNumberN()
	z := primes.getNumberZ()
	e := new(big.Int)

	for e.Cmp(big.NewInt(0)) == 0 {
		randomE := util.RandomNumber(2, 200)

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
