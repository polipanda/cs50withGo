package week1

import "testing"

func TestValidateCreditCardNumbers(t *testing.T) {
	scenarios := []struct {
		number int
		card   string
	}{
		{
			number: 1,
			card:   INVALID,
		},
		{
			number: 3759678978,
			card:   INVALID,
		},
		{
			number: 378282246310005,
			card:   AMEX,
		},
		{
			number: 371449635398431,
			card:   AMEX,
		},
		{
			number: 2221000000000009,
			card:   MASTERCARD,
		},
		{
			number: 2223000048400011,
			card:   MASTERCARD,
		},
		{
			number: 2223016768739313,
			card:   MASTERCARD,
		},
		{
			number: 4111111111111111,
			card:   VISA,
		},
		{
			number: 4012888888881881,
			card:   VISA,
		},
	}

	for _, scenario := range scenarios {
		result := ValidateCreditCardNumbers(scenario.number)
		if result != scenario.card {
			t.Errorf("expected %s with number %d, got %s", scenario.card, scenario.number, result)
		}
	}
}

func TestFirstNDigits(t *testing.T) {
	scenarios := []struct {
		number   int
		digits   int
		expected int
	}{
		{
			number:   1,
			digits:   1,
			expected: 1,
		},
		{
			number:   3759678978,
			digits:   2,
			expected: 37,
		},
		{
			number:   -1,
			digits:   1,
			expected: -1,
		},
	}

	for _, scenario := range scenarios {
		result := firstNDigits(scenario.number, scenario.digits)
		if result != scenario.expected {
			t.Errorf("expected %d, got %d", scenario.expected, result)
		}
	}
}

func TestIntLength(t *testing.T) {
	scenarios := []struct {
		number   int
		expected int
	}{
		{
			number:   0,
			expected: 1,
		},
		{
			number:   1,
			expected: 1,
		},
		{
			number:   375967,
			expected: 6,
		},
		{
			number:   3759678978,
			expected: 10,
		},
	}

	for _, scenario := range scenarios {
		result := intLength(scenario.number)
		if result != scenario.expected {
			t.Errorf("expected %d, got %d", scenario.expected, result)
		}
	}
}
