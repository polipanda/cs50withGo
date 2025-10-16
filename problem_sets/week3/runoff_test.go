package week3

import (
	"slices"
	"testing"
)

func TestHandleRunoffElected(t *testing.T) {
	scenarios := []struct {
		candidates  []string
		votes       [][]string
		winners     []string
		expectError bool
	}{
		{[]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}, nil, nil, true},
		{[]string{"1", "1", "2"}, nil, nil, true},
		{[]string{"1", "2", "3"}, [][]string{{"1", "2", "4"}}, nil, true},
		// scenario that ends at first level
		{
			[]string{"a", "b"},
			[][]string{
				{"a", "b"},
				{"a", "b"},
				{"b", "a"},
			},
			[]string{"a"},
			false,
		},
		// scenario that ends at second level
		{
			[]string{"a", "b", "c"},
			[][]string{
				{"a", "b", "c"},
				{"a", "b", "c"},
				{"b", "a", "c"},
				{"b", "a", "c"},
				{"c", "a", "b"},
			},
			[]string{"a"},
			false,
		},
		// scenario that ends at third level
		{
			[]string{"a", "b", "c", "d"},
			[][]string{
				{"a", "b", "c", "d"},
				{"a", "b", "c", "d"},
				{"a", "b", "c", "d"},
				{"b", "c", "a", "d"},
				{"b", "c", "a", "d"},
				{"b", "a", "c", "d"},
				{"c", "a", "b", "d"},
				{"d", "b", "a", "c"},
				{"d", "b", "a", "c"},
			},
			[]string{"b"},
			false,
		},
		// scenario that ends at fourth level
		{
			[]string{"a", "b", "c", "d", "e"},
			[][]string{
				{"a", "b", "c", "d", "e"},
				{"a", "b", "c", "d", "e"},
				{"a", "b", "c", "d", "e"},
				{"b", "c", "a", "d", "e"},
				{"b", "c", "a", "d", "e"},
				{"b", "a", "c", "d", "e"},
				{"c", "a", "b", "d", "e"},
				{"d", "b", "a", "c", "e"},
				{"d", "b", "a", "c", "e"},
			},
			[]string{"b"},
			false,
		},
	}

	for _, scenario := range scenarios {
		result, err := handleRunoffElection(scenario.candidates, scenario.votes)
		if scenario.expectError && err == nil {
			t.Errorf("Scenario %v should have thrown an error", scenario.candidates)
		}

		if !slices.Equal(scenario.winners, result) {
			t.Errorf("winners: %v != %v", result, scenario.winners)
		}
	}
}

func TestInitializeRunoffBallet(t *testing.T) {
	t.Run("error on too many candidates", func(t *testing.T) {
		candidates := make([]string, 10)
		_, err := initializeCandidatesResults(candidates, len(candidates)-1)
		if err == nil {
			t.Errorf("should not have gotten err: %v", err)
		}
	})

	t.Run("error with duplicate candidates", func(t *testing.T) {
		candidates := []string{"a", "b", "b"}
		_, err := initializeCandidatesResults(candidates, len(candidates))
		if err == nil {
			t.Errorf("should not have gotten err: %v", err)
		}
	})

	t.Run("ballot is made", func(t *testing.T) {
		candidates := []string{"a", "b", "c"}
		ballot, err := initializeCandidatesResults(candidates, len(candidates))
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

func TestValidateBallots(t *testing.T) {
	t.Run("ballot votes must equal candidates length", func(t *testing.T) {
		ballot := [][]string{
			{"a", "b"},
		}
		candidates := []string{"a", "b", "c"}
		if err := validateBallots(ballot, candidates); err == nil {
			t.Error("should have received an error")
		}
	})

	t.Run("ballot votes cannot be for an unnamed candidate", func(t *testing.T) {
		ballot := [][]string{
			{"a", "b", "d"},
		}
		candidates := []string{"a", "b", "c"}
		if err := validateBallots(ballot, candidates); err == nil {
			t.Error("should have received an error")
		}
	})

	t.Run("ballot votes cannot name the same person twice", func(t *testing.T) {
		ballot := [][]string{
			{"a", "b", "b"},
		}
		candidates := []string{"a", "b", "c"}
		if err := validateBallots(ballot, candidates); err == nil {
			t.Error("should have received an error")
		}
	})
}

func TestAllocateFirstVotes(t *testing.T) {
	results := map[string][]string{
		"a": {},
		"b": {},
		"c": {},
	}
	ballots := [][]string{
		{"a", "b", "c"},
		{"a", "c", "b"},
		{"b", "c", "a"},
	}

	results = allocateFirstVotes(results, ballots)

	expectedResults := map[string]int{
		"a": 2,
		"b": 1,
		"c": 0,
	}

	for r, expected := range expectedResults {
		if len(results[r]) != expected {
			t.Errorf("expected %s to have %d votes, received %d", r, expected, len(results[r]))
		}
	}
}

func TestMakeWinnersList(t *testing.T) {
	list := map[string][]string{
		"a": {},
		"b": {},
	}

	winners := makeWinnersList(list)
	if !slices.Equal(winners, []string{"a", "b"}) {
		t.Errorf("winners: %v != %v", winners, []string{"a", "b"})
	}
}
