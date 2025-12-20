package week2

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func substituteCipherText(text string, key string) (string, error) {
	if len(key) != 26 {
		return "", errors.New("key must be 26 chars")
	}

	alphabet, err := createSubstitutionMapping(key)
	if err != nil {
		return "", err
	}

	var cipherText string
	for _, char := range text {
		if unicode.IsLetter(char) {
			cipherText += alphabet[string(char)]
		} else {
			cipherText += string(char)
		}
	}

	return cipherText, nil
}

func createSubstitutionMapping(key string) (map[string]string, error) {
	m := make(map[int32]int)
	alphabet := make(map[string]string)
	for i, char := range key {
		if !unicode.IsLetter(char) {
			return nil, fmt.Errorf("key '%s' must be a letter", string(char))
		}

		c := unicode.ToLower(char)
		if _, ok := m[c]; ok {
			return nil, fmt.Errorf("key contains duplicate '%s'", string(char))
		}
		m[char] = i

		// add lowercase mapping
		alphabet[string('a'+rune(i))] = string(c)
		// add uppercase mapping
		alphabet[string('A'+rune(i))] = strings.ToUpper(string(c))
	}

	return alphabet, nil
}
