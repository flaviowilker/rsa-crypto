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

	return crypto
}

// Crypto ...
type Crypto struct {
	KeyPair *KeyPair
	Verbose bool
}

// EncryptText ...
func (cr *Crypto) EncryptText(text string) []int {

	runes := []rune(text)
	encrypted := make([]int, len(runes), len(runes))

	for k, v := range runes {
		c := new(big.Int).Exp(big.NewInt(int64(v)), cr.KeyPair.PrivateKey.PublicKey.E, cr.KeyPair.PrivateKey.PublicKey.N)

		fmt.Println()
		fmt.Println("char: ", big.NewInt(int64(v)))
		fmt.Println("char crypto: ", c)

		encrypted[k] = int(c.Int64())
	}

	fmt.Println("e: ", cr.KeyPair.PrivateKey.PublicKey.E)
	fmt.Println("n: ", cr.KeyPair.PrivateKey.PublicKey.N)

	return encrypted
}

// DecryptText ...
func (cr *Crypto) DecryptText(encrypted []int) string {

	decrypted := make([]int, len(encrypted), len(encrypted))

	for k, v := range encrypted {
		m := new(big.Int).Exp(big.NewInt(int64(v)), cr.KeyPair.PrivateKey.D, cr.KeyPair.PrivateKey.PublicKey.N)

		fmt.Println()
		fmt.Println("char crypto: ", big.NewInt(int64(v)))
		fmt.Println("char: ", m)

		decrypted[k] = int(m.Int64())
	}

	var decryptedText string

	for _, v := range decrypted {
		decryptedText += string(rune(v))
	}

	fmt.Println("d: ", cr.KeyPair.PrivateKey.D)
	fmt.Println("n: ", cr.KeyPair.PrivateKey.PublicKey.N)

	return decryptedText
}
