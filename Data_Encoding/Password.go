package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

func main() {
	password := generatePassword(8, 20)
	fmt.Println(password)
}

func generatePassword(minLength, maxLength int) string {
	specialChars := "@%+\\/'!#$^?:,(){}[]~-_."
	lowercaseChars := "abcdefghijklmnopqrstuvwxyz"
	uppercaseChars := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numberChars := "0123456789"

	characterSet := []string{specialChars, lowercaseChars, uppercaseChars, numberChars}

	var password strings.Builder

	// Add at least one character from each character set
	for _, charSet := range characterSet {
		char, _ := getRandomCharacter(charSet)
		password.WriteByte(char)
	}

	// Add remaining characters up to maxLength
	for i := 4; i <= maxLength; i++ {
		charSet := characterSet[i%4]
		char, _ := getRandomCharacter(charSet)
		password.WriteByte(char)
	}

	return password.String()
}

func getRandomCharacter(charSet string) (byte, error) {
	randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(charSet))))
	if err != nil {
		return 0, err
	}
	return charSet[randomIndex.Int64()], nil
}
