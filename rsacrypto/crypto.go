package rsacrypto

import (
	"fmt"
	"math/big"
)

// EncryptText ...
func EncryptText(publicKey *PublicKey, text string) []int {

	runes := []rune(text)
	encrypted := make([]int, len(runes), len(runes))

	for k, v := range runes {
		c := new(big.Int).Exp(big.NewInt(int64(v)), publicKey.E, publicKey.N)

		fmt.Println()
		fmt.Println("char: ", big.NewInt(int64(v)))
		fmt.Println("e: ", publicKey.E)
		fmt.Println("n: ", publicKey.N)
		fmt.Println("char crypto: ", c)

		encrypted[k] = int(c.Int64())
	}

	return encrypted
}

// DecryptText ...
func DecryptText(privateKey *PrivateKey, encrypted []int) string {

	decrypted := make([]int, len(encrypted), len(encrypted))

	for k, v := range encrypted {
		m := new(big.Int).Exp(big.NewInt(int64(v)), privateKey.D, privateKey.PublicKey.N)

		fmt.Println()
		fmt.Println("char crypto: ", big.NewInt(int64(v)))
		fmt.Println("d: ", privateKey.D)
		fmt.Println("n: ", privateKey.PublicKey.N)
		fmt.Println("char: ", m)

		decrypted[k] = int(m.Int64())
	}

	var decryptedText string

	for _, v := range decrypted {
		decryptedText += string(rune(v))
	}

	return decryptedText
}
