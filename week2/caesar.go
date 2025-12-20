package week2

import (
	"fmt"
	"unicode"
)

func cipherText(plainText string, key int) (string, error) {
	if key < 0 {
		return "", fmt.Errorf("key must be >= 0")
	}

	shift := rune(key % 26)
	var cipherText string
	for _, char := range plainText {
		if !unicode.IsLetter(char) {
			cipherText += string(char)
			continue
		}

		if unicode.IsUpper(char) && char+shift > 90 {
			cipherText += string(64 + (char + shift - 90))
		} else if unicode.IsLower(char) && char+shift > 122 {
			cipherText += string(96 + (char + shift - 122))
		} else {
			cipherText += string(char + shift)
		}
	}

	return cipherText, nil
}
