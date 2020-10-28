package run

import (
	"fmt"

	"github.com/flaviowilker/rsa-crypto/rsacrypto"
)

// App ...
func App() {

	text := "Hello World"

	crypto := rsacrypto.NewCrypto(false)

	encrypted := crypto.EncryptText(text)
	fmt.Println(encrypted)

	decrypted := crypto.DecryptText(encrypted)
	fmt.Println(decrypted)
}
