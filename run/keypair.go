package run

import (
	"fmt"

	"github.com/flaviowilker/rsa-crypto/rsacrypto"
	"github.com/flaviowilker/rsa-crypto/util"
)

func createKeyPair(verbose bool) *rsacrypto.KeyPair {
	keyPair := rsacrypto.NewKeyPair()

	if verbose {
		fmt.Println("Key numbers")
		fmt.Println("n: ", keyPair.PrivateKey.PublicKey.N)
		fmt.Println("e: ", keyPair.PrivateKey.PublicKey.E)
		fmt.Println("d: ", keyPair.PrivateKey.D)
		fmt.Println("Prime numbers")
		fmt.Println("p: ", keyPair.Primes.P)
		fmt.Println("q: ", keyPair.Primes.Q)
		fmt.Println("Other numbers")
		fmt.Println("z: ", keyPair.Primes.GetNumberZ())
		fmt.Println()
	}

	fmt.Println("Key pair")
	fmt.Printf("Public Key (n, e): (%v, %v)\n", keyPair.PrivateKey.PublicKey.N, keyPair.PrivateKey.PublicKey.E)
	fmt.Printf("Private Key (n, d): (%v, %v)\n", keyPair.PrivateKey.PublicKey.N, keyPair.PrivateKey.D)
	fmt.Println()

	return keyPair
}

func readText() string {
	fmt.Print("Enter text: ")
	text := util.TextReader()
	fmt.Println()

	return text
}

func encryptDecryptText(text string, verbose bool) {
	keyPair := createKeyPair(verbose)

	encrypted := encryptText(text, keyPair.PrivateKey.PublicKey, verbose)

	decryptText(encrypted, keyPair.PrivateKey, verbose)
}
