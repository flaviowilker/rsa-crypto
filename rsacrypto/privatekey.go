package rsacrypto

import (
	"math/big"

	"github.com/flaviowilker/rsa-crypto/util"
)

func newPrivateKey(primes *Primes, publicKey *PublicKey) *PrivateKey {

	e := publicKey.E
	z := primes.getNumberZ()

	d := new(big.Int)
	mod := new(big.Int)

	for mod.Cmp(big.NewInt(1)) != 0 {
		d = util.RandomNumber(10, 10000)

		mul := new(big.Int).Mul(e, d)
		mod = mod.Mod(mul, z)
	}

	privateKey := &PrivateKey{
		PublicKey: publicKey,
		D:         d,
	}

	return privateKey
}

// PrivateKey ...
type PrivateKey struct {
	PublicKey *PublicKey
	D         *big.Int
}
