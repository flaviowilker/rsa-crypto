package rsacrypto

import (
	"fmt"
	"math/big"
)

// NewCrypto ...
func NewCrypto(verbose bool) *Crypto {

	keyPair := newKeyPair()

	crypto := &Crypto{
		KeyPair: keyPair,
		Verbose: verbose,
	}

	if crypto.Verbose {
		fmt.Println("Key numbers")
		fmt.Println("n: ", crypto.KeyPair.PrivateKey.PublicKey.N)
		fmt.Println("e: ", crypto.KeyPair.PrivateKey.PublicKey.E)
		fmt.Println("d: ", crypto.KeyPair.PrivateKey.D)
		fmt.Println("Prime numbers")
		fmt.Println("p: ", crypto.KeyPair.Primes.P)
		fmt.Println("q: ", crypto.KeyPair.Primes.Q)
		fmt.Println("Other numbers")
		fmt.Println("z: ", crypto.KeyPair.Primes.getNumberZ())
		fmt.Println()
	}

	return crypto
}

// Crypto ...
type Crypto struct {
	KeyPair *KeyPair
	Verbose bool
}

// EncryptText ...
func (cr *Crypto) EncryptText(text string) []int {

	if cr.Verbose {
		fmt.Println()
		fmt.Println("Encrypt text")
	}

	runes := []rune(text)
	encrypted := make([]int, len(runes), len(runes))

	for k, v := range runes {
		c := new(big.Int).Exp(big.NewInt(int64(v)), cr.KeyPair.PrivateKey.PublicKey.E, cr.KeyPair.PrivateKey.PublicKey.N)

		encrypted[k] = int(c.Int64())

		if cr.Verbose {
			fmt.Println("char ", k+1)
			fmt.Println("char number: ", big.NewInt(int64(v)))
			fmt.Println("char encrypted: ", c)
		}
	}

	return encrypted
}

// DecryptText ...
func (cr *Crypto) DecryptText(encrypted []int) string {

	if cr.Verbose {
		fmt.Println()
		fmt.Println("Decrypt text")
	}

	decrypted := make([]int, len(encrypted), len(encrypted))

	for k, v := range encrypted {
		m := new(big.Int).Exp(big.NewInt(int64(v)), cr.KeyPair.PrivateKey.D, cr.KeyPair.PrivateKey.PublicKey.N)

		decrypted[k] = int(m.Int64())

		if cr.Verbose {
			fmt.Println("char ", k+1)
			fmt.Println("char number: ", big.NewInt(int64(v)))
			fmt.Println("char decrypted: ", m)
		}
	}

	var decryptedText string

	for _, v := range decrypted {
		decryptedText += string(rune(v))
	}

	return decryptedText
}
