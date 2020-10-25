package rsacrypto

import (
	"errors"
	"math"

	"github.com/flaviowilker/rsa-crypto/util"
)

// NewKeyPair ...
func NewKeyPair() *KeyPair {

	primes := NewPrimes()
	keyPair := &KeyPair{Primes: primes}

	err := keyPair.createNumberN()
	if err != nil {
		panic(err)
	}

	err = keyPair.createNumberE()
	if err != nil {
		panic(err)
	}

	err = keyPair.createNumberD()
	if err != nil {
		panic(err)
	}

	return keyPair
}

// KeyPair ...
type KeyPair struct {
	Primes     *Primes
	PrivateKey *PrivateKey
}

func (k *KeyPair) createNumberN() error {
	if k.Primes.P != 0 && k.Primes.Q != 0 {
		k.PrivateKey = &PrivateKey{}
		k.PrivateKey.PublicKey = &PublicKey{}

		k.PrivateKey.PublicKey.N = k.Primes.getNumberN()
	} else {
		return errors.New("field Primes is null")
	}
	return nil
}

func (k *KeyPair) createNumberE() error {
	if k.Primes.P != 0 && k.Primes.Q != 0 {
		var e int
		z := k.Primes.getNumberZ()

		for k.PrivateKey.PublicKey.E == 0 {
			e = util.RandomNumber(2, 1000)
			if util.IsMdcPrimes(z, e) {
				k.PrivateKey.PublicKey.E = e
			}
		}
	} else {
		return errors.New("field Primes is null")
	}
	return nil
}

func (k *KeyPair) createNumberD() error {
	if k.PrivateKey.PublicKey.E != 0 && k.Primes.P != 0 && k.Primes.Q != 0 {
		e := float64(k.PrivateKey.PublicKey.E)
		z := float64(k.Primes.getNumberZ())
		d := float64((z + 1) / e)

		mod := math.Mod((e * d), z)

		if mod == 1.0 {
			k.PrivateKey.D = d
		} else {
			return errors.New("create number 'd' not work")
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
	D         float64
}
