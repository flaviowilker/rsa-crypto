package util

import (
	"math/rand"
	"time"
)

// IsPrime ...
func IsPrime(number int) bool {
	if number < 2 {
		return false
	}

	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return number%number == 0
}

// IsMdcPrimes ...
func IsMdcPrimes(num1, num2 int) bool {

	var mdcPrimes func(num1, num2 int) int
	mdcPrimes = func(num1, num2 int) int {
		if num2 == 0 {
			return num1
		}
		return mdcPrimes(num2, num1%num2)
	}
	return mdcPrimes(num1, num2) == 1
}

// RandomNumber ...
func RandomNumber(minNumber int, maxNumber int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(maxNumber-minNumber+1) + minNumber
}

// SliceContains ...
func SliceContains(slice []int, value int) bool {
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
