package run

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/flaviowilker/rsa-crypto/rsacrypto"
)

func readPrivateKey(verbose bool) *rsacrypto.PrivateKey {
	fmt.Println("Private Key")

	var n, d int

	fmt.Println("Write the N number")
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		fmt.Println("error to read a integer")
		panic(err)
	}

	fmt.Println("Write the D number")
	_, err = fmt.Scanf("%d", &d)
	if err != nil {
		fmt.Println("error to read a integer")
		panic(err)
	}
	fmt.Println()

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
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Write separated by space \" \"")
	fmt.Print("Enter the encrypted text: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	text = strings.TrimSuffix(text, "\n")
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
