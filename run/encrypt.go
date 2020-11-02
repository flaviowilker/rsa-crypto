package run

import (
	"fmt"
	"math/big"

	"github.com/flaviowilker/rsa-crypto/rsacrypto"
)

func readPublicKey(verbose bool) *rsacrypto.PublicKey {
	fmt.Println("Public Key")

	var n, e int

	fmt.Println("Write the N number")
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println("error to read a integer")
		panic(err)
	}

	fmt.Println("Write the E number")
	_, err = fmt.Scanf("%d", &e)
	if err != nil {
		fmt.Println("error to read a integer")
		panic(err)
	}

	publicKey := &rsacrypto.PublicKey{
		N: big.NewInt(int64(n)),
		E: big.NewInt(int64(e)),
	}

	return publicKey
}

func encryptText(text string, publicKey *rsacrypto.PublicKey, verbose bool) []int {
	encrypted := rsacrypto.EncryptText(text, publicKey, verbose)
	fmt.Printf("Encrypted: %v\n", encrypted)
	fmt.Println()

	return encrypted
}
