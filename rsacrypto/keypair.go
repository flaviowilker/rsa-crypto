package rsacrypto

import (
	"errors"
	"fmt"

	"github.com/flaviowilker/rsa-crypto/util"
)

// CreateKeyPair ...
func CreateKeyPair() (*KeyPair, error) {

	primes := NewPrimes()

	fmt.Println(primes)

	return &KeyPair{}, nil
}

// KeyPair ...
type KeyPair struct {
	Primes     *Primes
	PrivateKey *PrivateKey
}

// CreateNumberN ...
func (k *KeyPair) CreateNumberN() error {
	if k.Primes.P != 0 && k.Primes.Q != 0 {
		k.PrivateKey.PublicKey.N = k.Primes.GetNumberN()
	} else {
		return errors.New("field Primes is null")
	}
	return nil
}

// CreateNumberE ...
func (k *KeyPair) CreateNumberE() error {
	if k.Primes.P != 0 && k.Primes.Q != 0 {
		var numberECreated bool
		var e int
		z := k.Primes.GetNumberZ()

		for numberECreated == false {
			e = util.RandomNumber(2, 1000)
			numberECreated = util.IsMdcPrimes(z, e)
		}
		k.PrivateKey.PublicKey.E = e
	} else {
		return errors.New("field Primes is null")
	}
	return nil
}

// CreateNumberD ...
func (k *KeyPair) CreateNumberD() error {
	if k.PrivateKey.PublicKey.E != 0 && k.Primes.P != 0 && k.Primes.Q != 0 {
		var d int
		e := k.PrivateKey.PublicKey.E
		z := k.Primes.GetNumberZ()

		for d == 0 {
			randomD := util.RandomNumber(2, 1000)
			if (e * randomD % z) == 1 {
				d = randomD
			}
		}
	} else {
		return errors.New("field PrivateKey.PublicKey.E and Primes is null")
	}
	return nil
}

// PublicKey ...
type PublicKey struct {
	N int
	E int
}

// PrivateKey ...
type PrivateKey struct {
	PublicKey *PublicKey
	D         int
}
