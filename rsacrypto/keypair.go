package rsacrypto

import (
	"fmt"
	"math/big"
)

// KeyPair ...
type KeyPair struct {
	Primes     *Primes
	PrivateKey *PrivateKey
}

// PublicKey ...
type PublicKey struct {
	N *big.Int
	E *big.Int
}

// PrivateKey ...
type PrivateKey struct {
	PublicKey *PublicKey
	D         *big.Int
}

// CreateKeyPair ...
func CreateKeyPair() (*KeyPair, error) {

	primes := ToChosePrimes()

	fmt.Println(primes)

	return &KeyPair{}, nil
}
