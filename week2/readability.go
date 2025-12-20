package week2

import (
	"math"
	"strings"
	"unicode"
)

const BEFOREGRADE1 = "Before Grade 1"
const GRADE2 = "Grade 2"
const GRADE3 = "Grade 3"
const GRADE5 = "Grade 5"
const GRADE7 = "Grade 7"
const GRADE8 = "Grade 8"
const GRADE9 = "Grade 9"
const GRADE10 = "Grade 10"
const GRADE16PLUS = "Grade 16+"

func determineReadingLevel(text string) string {
	letters := countLetters(text)
	sentences := countSentences(text)
	words := countWords(text)
	rate := 100.0
	avgLetters := (float64(letters) / float64(words)) * rate
	avgSentences := (float64(sentences) / float64(words)) * rate

	index := calculateIndex(avgLetters, avgSentences)

	return gradeToString(int(math.Round(index)))
}

func countLetters(text string) int {
	count := 0

	for _, letter := range text {
		if unicode.IsLetter(letter) {
			count++
		}
	}

	return count
}

func countSentences(text string) int {
	count := 0
	for _, r := range text {
		switch r {
		case '.', '!', '?':
			count++
		}
	}

	return count
}

func countWords(text string) int {
	return len(strings.Fields(text))
}

func calculateIndex(letterOccurrance, sentenceOccurrance float64) float64 {

	return 0.0588*letterOccurrance - 0.296*sentenceOccurrance - 15.8
}

func gradeToString(grade int) string {
	if grade < 1 {
		return BEFOREGRADE1
	}

	if grade >= 16 {
		return GRADE16PLUS
	}

	switch grade {
	case 2:
		return GRADE2
	case 3:
		return GRADE3
	case 5:
		return GRADE5
	case 7:
		return GRADE7
	case 8:
		return GRADE8
	case 9:
		return GRADE9
	case 10:
		return GRADE10
	}

	return "Not a valid grade"
}
