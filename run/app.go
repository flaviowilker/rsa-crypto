package run

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/flaviowilker/rsa-crypto/rsacrypto"
)

var verbose = flag.Bool("v", false, "verbose for algorithm")

// App ...
func App() {
	flag.Parse()

	fmt.Println("RSA Crypto")
	fmt.Println("---------------------")

	text := textReader()

	encryptDecryptText(text, *verbose)
}

func textReader() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	text = strings.TrimSuffix(text, "\n")
	fmt.Println()

	return text
}

func encryptDecryptText(text string, verbose bool) {
	crypto := rsacrypto.NewCrypto(verbose)

	encrypted := crypto.EncryptText(text)
	fmt.Print("Encrypted: ")
	fmt.Println(encrypted)

	decrypted := crypto.DecryptText(encrypted)
	fmt.Print("Decrypted: ")
	fmt.Println(decrypted)
}
