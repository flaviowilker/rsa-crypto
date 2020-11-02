package run

import (
	"flag"
	"fmt"
)

var verbose = flag.Bool("v", false, "verbose for algorithm")
var keyPair = flag.Bool("keypair", false, "create keypair")
var encrypt = flag.Bool("encrypt", false, "encrypt string")
var decrypt = flag.Bool("decrypt", false, "encrypt string")

// App ...
func App() {
	flag.Parse()

	fmt.Println("RSA Crypto")
	fmt.Println("----------")

	if *keyPair {
		createKeyPair(*verbose)
	}
	if *encrypt {
		publicKey := readPublicKey(*verbose)
		text := readText()
		encryptText(text, publicKey, *verbose)
	}
	if *decrypt {
		privateKey := readPrivateKey(*verbose)
		encryptedText := readEncryptedText()
		decryptText(encryptedText, privateKey, *verbose)
	}

	if !*keyPair && !*encrypt && !*decrypt {
		text := readText()
		encryptDecryptText(text, *verbose)
	}
}
