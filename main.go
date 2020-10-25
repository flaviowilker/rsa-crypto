package main

import (
	"fmt"

	"github.com/flaviowilker/rsa-crypto/rsacrypto"
)

func main() {

	text := "A"

	keyPair := rsacrypto.NewKeyPair()

	encrypted := rsacrypto.EncryptText(keyPair.PrivateKey.PublicKey, text)

	fmt.Println(encrypted)

	decrypted := rsacrypto.DecryptText(keyPair.PrivateKey, encrypted)

	fmt.Println(decrypted)
}
