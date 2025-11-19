package lib

import (
	"math/rand"
	"time"
)

func GenerateReferenceNumber() string {
	referenceNumber := make([]rune, 9)
	characters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	charactersLength := len(characters)

	numbers := []rune("1234567890")
	numbersLength := len(numbers)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 3; i++ {
		referenceNumber[i] = characters[rand.Intn(charactersLength-1)]
	}

	for i := 3; i < 9; i++ {
		referenceNumber[i] = numbers[rand.Intn(numbersLength-1)]
	}
	return string(referenceNumber)
}

func StrPadLeft(input string, padLength int, padString string) string {
	output := padString

	for padLength > len(output) {
		output += output
	}

	if len(input) >= padLength {
		return input
	}

	return output[:padLength-len(input)] + input
}
