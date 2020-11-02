package rsacrypto

import (
	"fmt"
	"math/big"
)

// EncryptText ...
func EncryptText(text string, publicKey *PublicKey, verbose bool) []int {

	if verbose {
		fmt.Println("Encrypt text")
	}

	runes := []rune(text)
	encrypted := make([]int, len(runes), len(runes))

	for k, v := range runes {
		c := new(big.Int).Exp(big.NewInt(int64(v)), publicKey.E, publicKey.N)

		encrypted[k] = int(c.Int64())

		if verbose {
			fmt.Printf("char %d: %s\n", k+1, string(rune(v)))
			fmt.Println("char ascii: ", big.NewInt(int64(v)))
			fmt.Println("char encrypted: ", c)
		}
	}

	if verbose {
		fmt.Println()
	}

	return encrypted
}

// DecryptText ...
func DecryptText(encrypted []int, privateKey *PrivateKey, verbose bool) string {

	if verbose {
		fmt.Println("Decrypt text")
	}

	decrypted := make([]int, len(encrypted), len(encrypted))

	for k, v := range encrypted {
		m := new(big.Int).Exp(big.NewInt(int64(v)), privateKey.D, privateKey.PublicKey.N)

		decrypted[k] = int(m.Int64())

		if verbose {
			fmt.Printf("char %d: %s\n", k+1, string(rune(m.Int64())))
			fmt.Println("char encrypted: ", big.NewInt(int64(v)))
			fmt.Println("char ascii: ", m)
		}
	}

	if verbose {
		fmt.Println()
	}

	var decryptedText string

	for _, v := range decrypted {
		decryptedText += string(rune(v))
	}

	return decryptedText
}
