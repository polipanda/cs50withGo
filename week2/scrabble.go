package week2

import "strings"

const PLAYER1WINS = "PLAYER 1 WINS!"
const PLAYER2WINS = "PLAYER 2 WINS!"
const TIE = "TIE!"

var pointsMap = map[rune]int{
	'a': 1,
	'b': 3,
	'c': 3,
	'd': 2,
	'e': 1,
	'f': 4,
	'g': 2,
	'h': 4,
	'i': 1,
	'j': 8,
	'k': 5,
	'l': 1,
	'm': 3,
	'n': 1,
	'o': 1,
	'p': 3,
	'q': 10,
	'r': 1,
	's': 1,
	't': 2,
	'u': 1,
	'v': 4,
	'w': 4,
	'x': 8,
	'y': 4,
	'z': 10,
}

func determineScrabbleWinner(player1word, player2word string) string {
	// remove any non letters
	player1word = strings.ToLower(stripNonAlphabetical(player1word))
	player2word = strings.ToLower(stripNonAlphabetical(player2word))

	player1score := countScore(player1word)
	player2score := countScore(player2word)

	if player1score > player2score {
		return PLAYER1WINS
	}

	if player1score < player2score {
		return PLAYER2WINS
	}

	return TIE
}

func stripNonAlphabetical(word string) string {
	var result strings.Builder

	for i := 0; i < len(word); i++ {
		b := word[i]
		if ('a' <= b && b <= 'z') || ('A' <= b && b <= 'Z') || b == ' ' {
			result.WriteByte(b)
		}
	}

	return result.String()
}

func countScore(word string) int {
	var score int
	for _, b := range word {
		if val, ok := pointsMap[b]; ok {
			score += val
		}
	}

	return score
}
