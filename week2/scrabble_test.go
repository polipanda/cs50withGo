package week2

import "testing"

func TestDetermineScrabbleWinner(t *testing.T) {
	t.Run("Player 1 wins", func(t *testing.T) {
		player1word := "exquisite"
		player2word := "ELEPHANT"
		result := determineScrabbleWinner(player1word, player2word)
		if result != PLAYER1WINS {
			t.Errorf("result should be %s, got %s", PLAYER1WINS, result)
		}
	})

	t.Run("Player 2 wins", func(t *testing.T) {
		player1word := "potato"
		player2word := "satellite"
		result := determineScrabbleWinner(player1word, player2word)
		if result != PLAYER2WINS {
			t.Errorf("result should be %s, got %s", PLAYER2WINS, result)
		}
	})

	t.Run("Tie", func(t *testing.T) {
		player1word := "function"
		player2word := "function"
		result := determineScrabbleWinner(player1word, player2word)
		if result != TIE {
			t.Errorf("result should be %s, got %s", TIE, result)
		}
	})
}

func TestStripNonAlphabetical(t *testing.T) {
	input := "12abc12"
	result := stripNonAlphabetical(input)
	if result != "abc" {
		t.Errorf("result should be %s, got %s", "abc", result)
	}
}

func TestCountScore(t *testing.T) {
	result := countScore("exquisite")
	if result != 26 {
		t.Errorf("result should be %d, got %d", 0, result)
	}
}
