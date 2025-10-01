package week1

import "testing"

func TestCoinsOwed(t *testing.T) {
	scenarios := []struct {
		change        int
		expectedCoins int
	}{
		{
			change:        -1,
			expectedCoins: 0,
		},
		{
			change:        0,
			expectedCoins: 0,
		},
		{
			change:        1,
			expectedCoins: 1,
		},
		{
			change:        4,
			expectedCoins: 4,
		},
		{
			change:        5,
			expectedCoins: 1,
		},
		{
			change:        24,
			expectedCoins: 6,
		},
		{
			change:        25,
			expectedCoins: 1,
		},
		{
			change:        26,
			expectedCoins: 2,
		},
		{
			change:        99,
			expectedCoins: 9,
		},
	}

	for _, scenario := range scenarios {
		result := coinsOwed(scenario.change)
		if result != scenario.expectedCoins {
			t.Errorf("coinsOwed(%d): expected %d, got %d", scenario.change, scenario.expectedCoins, result)
		}
	}
}
