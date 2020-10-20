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
