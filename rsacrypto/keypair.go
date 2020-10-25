package rsacrypto

// NewKeyPair ...
func NewKeyPair() *KeyPair {

	primes := newPrimes()
	publicKey := newPublicKey(primes)
	privateKey, err := newPrivateKey(primes, publicKey)
	if err != nil {
		panic(err)
	}

	keyPair := &KeyPair{Primes: primes, PrivateKey: privateKey}

	return keyPair
}

// KeyPair ...
type KeyPair struct {
	Primes     *Primes
	PrivateKey *PrivateKey
}
