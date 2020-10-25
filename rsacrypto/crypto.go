package rsacrypto

import "math"

// EncryptText ...
func EncryptText(publicKey *PublicKey, text string) []int {

	runes := []rune(text)
	encrypted := make([]int, len(runes), len(runes))

	for k, v := range runes {
		power := math.Pow(float64(v), float64(publicKey.E.Int64()))
		mod := int(math.Mod(power, float64(publicKey.N.Int64())))

		encrypted[k] = mod
	}

	return encrypted
}

// DecryptText ...
func DecryptText(privateKey *PrivateKey, encrypted []int) string {

	decrypted := make([]int, len(encrypted), len(encrypted))

	for k, v := range encrypted {
		power := math.Pow(float64(v), privateKey.D)
		mod := math.Mod(power, float64(privateKey.PublicKey.N.Int64()))

		decrypted[k] = int(mod)
	}

	var decryptedText string

	for _, v := range decrypted {
		decryptedText += string(rune(v))
	}

	return decryptedText
}
