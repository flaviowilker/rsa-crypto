package util

import (
	"bufio"
	"math/big"
	"math/rand"
	"os"
	"strings"
	"time"
)

// IsPrime ...
func IsPrime(number *big.Int) bool {
	if number.Cmp(big.NewInt(2)) == -1 {
		return false
	}

	numberInt := number.Int64()
	for i := int64(2); i < numberInt; i++ {
		mod := new(big.Int).Mod(number, big.NewInt(int64(i)))

		if mod.Cmp(big.NewInt(0)) == 0 {
			return false
		}
	}

	mod := new(big.Int).Mod(number, number)

	return mod.Cmp(big.NewInt(0)) == 0
}

// IsMdcPrimes ...
func IsMdcPrimes(num1, num2 *big.Int) bool {

	var mdcPrimes func(num1, num2 *big.Int) *big.Int
	mdcPrimes = func(num1, num2 *big.Int) *big.Int {
		if num2.Cmp(big.NewInt(0)) == 0 {
			return num1
		}

		mod := new(big.Int).Mod(num1, num2)

		return mdcPrimes(num2, mod)
	}

	return mdcPrimes(num1, num2).Cmp(big.NewInt(1)) == 0
}

// RandomNumber ...
func RandomNumber(minNumber int, maxNumber int) *big.Int {
	rand.Seed(time.Now().UnixNano())
	randomInt := rand.Int63n(int64(maxNumber-minNumber+1)) + int64(minNumber)

	return big.NewInt(randomInt)
}

// SliceContains ...
func SliceContains(slice []*big.Int, value *big.Int) bool {
	for _, v := range slice {
		if v.Cmp(value) == 0 {
			return true
		}
	}

	return false
}

// TextReader ...
func TextReader() string {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	text = strings.TrimRight(text, "\r\n")

	return text
}
