package week3

import (
	"slices"
	"testing"
)

func TestHandleElection(t *testing.T) {
	scenarios := []struct {
		candidates  []string
		votes       []string
		winners     []string
		expectError bool
	}{
		{[]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}, nil, nil, true},
		{[]string{"1", "1", "2"}, nil, nil, true},
		{[]string{"a", "b", "c"}, []string{"a", "a", "b"}, []string{"a"}, false},
		{[]string{"a", "b", "c"}, []string{"a", "a", "e"}, []string{"a"}, false},
		{[]string{"a", "b", "c"}, []string{"a", "b", "c"}, []string{"a", "b", "c"}, false},
	}

	for _, scenario := range scenarios {
		result, err := handlePluralityElection(scenario.candidates, scenario.votes)
		if err != nil && !scenario.expectError {
			t.Errorf("should not have gotten err: %v", err)
		}

		slices.Sort(result)
		if !slices.Equal(result, scenario.winners) {
			t.Errorf("winners: %v != %v", result, scenario.winners)
		}
	}
}

func TestInitializePluralityBallots(t *testing.T) {
	t.Run("error on too many candidates", func(t *testing.T) {
		candidates := make([]string, 10)
		_, err := initializePluralityBallot(candidates, len(candidates)-1)
		if err == nil {
			t.Errorf("should not have gotten err: %v", err)
		}
	})

	t.Run("error with duplicate candidates", func(t *testing.T) {
		candidates := []string{"a", "b", "b"}
		_, err := initializePluralityBallot(candidates, len(candidates))
		if err == nil {
			t.Errorf("should not have gotten err: %v", err)
		}
	})

	t.Run("ballot is made", func(t *testing.T) {
		candidates := []string{"a", "b", "c"}
		ballot, err := initializePluralityBallot(candidates, len(candidates))
		if err != nil {
			t.Errorf("should not have gotten err: %v", err)
		}

		for _, candidate := range candidates {
			_, ok := ballot[candidate]
			if !ok {
				t.Errorf("ballot should contain candidate: %v", candidate)
			}
		}
	})
}

func TestCountVotes(t *testing.T) {
	ballot := map[string]int{
		"a": 0,
		"b": 0,
	}

	votes := []string{"a", "a", "b"}

	result := countVotes(ballot, votes)

	if result["a"] != 2 || result["b"] != 1 {
		t.Errorf("ballot should contain all votes: %v", votes)
	}
}

func TestSelectWinners(t *testing.T) {
	t.Run("single winner", func(t *testing.T) {
		results := map[string]int{
			"a": 2,
			"b": 1,
			"c": 1,
		}

		winners := selectWinners(results)

		if len(winners) > 1 || winners[0] != "a" {
			t.Errorf("winners should only contain a, not: %v", winners)
		}
	})

	t.Run("multiple winners", func(t *testing.T) {
		results := map[string]int{
			"a": 2,
			"b": 2,
			"c": 1,
		}

		winners := selectWinners(results)
		slices.Sort(winners)
		expectedList := []string{"a", "b"}

		if len(winners) != 2 || !slices.Equal(winners, expectedList) {
			t.Errorf("winners should match expectedList: %v, got %v", expectedList, winners)
		}
	})
}
