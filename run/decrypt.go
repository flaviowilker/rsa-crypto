package run

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"

	"github.com/flaviowilker/rsa-crypto/rsacrypto"
	"github.com/flaviowilker/rsa-crypto/util"
)

func readPrivateKey(verbose bool) *rsacrypto.PrivateKey {
	fmt.Println("Private Key")

	fmt.Println("Write the N number")
	nString := util.TextReader()

	fmt.Println("Write the D number")
	dString := util.TextReader()
	fmt.Println()

	var n, d int
	var err error

	n, err = strconv.Atoi(nString)
	if err != nil {
		fmt.Println("error to read a integer")
		panic(err)
	}

	d, err = strconv.Atoi(dString)
	if err != nil {
		fmt.Println("error to read a integer")
		panic(err)
	}

	publicKey := &rsacrypto.PublicKey{
		N: big.NewInt(int64(n)),
	}

	privateKey := &rsacrypto.PrivateKey{
		PublicKey: publicKey,
		D:         big.NewInt(int64(d)),
	}

	return privateKey
}

func readEncryptedText() []int {
	fmt.Println("Write separated by space \" \"")
	fmt.Print("Enter the encrypted text: ")
	text := util.TextReader()
	fmt.Println()

	textSlice := strings.Split(text, " ")
	encryptedText := make([]int, len(textSlice), len(textSlice))

	for k, v := range textSlice {
		i, err := strconv.Atoi(v)

		if err != nil {
			fmt.Println("error to read a integer")
			panic(err)
		}

		encryptedText[k] = i
	}

	return encryptedText
}

func decryptText(encrypted []int, privateKey *rsacrypto.PrivateKey, verbose bool) string {

	decrypted := rsacrypto.DecryptText(encrypted, privateKey, verbose)
	fmt.Printf("Decrypted: %v\n", decrypted)
	fmt.Println()

	return decrypted
}
