package rsacrypto

// NewKeyPair ...
func NewKeyPair() *KeyPair {

	primes := newPrimes()
	publicKey := newPublicKey(primes)
	privateKey := newPrivateKey(primes, publicKey)

	keyPair := &KeyPair{Primes: primes, PrivateKey: privateKey}

	return keyPair
}

// KeyPair ...
type KeyPair struct {
	Primes     *Primes
	PrivateKey *PrivateKey
}
