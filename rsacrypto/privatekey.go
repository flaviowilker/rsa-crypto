package rsacrypto

import (
	"errors"
	"math"
)

func newPrivateKey(primes *Primes, publicKey *PublicKey) (*PrivateKey, error) {

	eFloat := float64(publicKey.E.Int64())
	zFloat := float64(primes.getNumberZ().Int64())
	d := (zFloat + 1) / eFloat

	mod := math.Mod((eFloat * d), zFloat)

	if mod == 1.0 {
		privateKey := &PrivateKey{
			PublicKey: publicKey,
			D:         d,
		}

		return privateKey, nil
	}

	return nil, errors.New("PrivateKey: create number 'd' not work")
}

// PrivateKey ...
type PrivateKey struct {
	PublicKey *PublicKey
	D         float64
}
