package rsacrypto

import (
	"fmt"
	"math/big"

	"github.com/flaviowilker/rsa-crypto/util"
)

func newPrivateKey(primes *Primes, publicKey *PublicKey) *PrivateKey {

	e := publicKey.E
	z := primes.getNumberZ()

	fmt.Println("e: ", e)
	fmt.Println("z: ", z)

	d := new(big.Int)
	mod := new(big.Int)

	for mod.Cmp(big.NewInt(1)) != 0 {
		d = util.RandomNumber(1, 10000)

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
