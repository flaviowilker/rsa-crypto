package run

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/flaviowilker/rsa-crypto/rsacrypto"
	"github.com/flaviowilker/rsa-crypto/util"
)

func readPublicKey(verbose bool) *rsacrypto.PublicKey {
	fmt.Println("Public Key")

	fmt.Println("Write the N number")
	nString := util.TextReader()

	fmt.Println("Write the E number")
	eString := util.TextReader()

	var n, e int
	var err error

	n, err = strconv.Atoi(nString)
	if err != nil {
		fmt.Println("error to read a integer")
		panic(err)
	}

	e, err = strconv.Atoi(eString)
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
