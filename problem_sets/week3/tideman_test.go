package week3

import (
	"slices"
	"testing"
)

func TestRunTidemanRankedPairsElection(t *testing.T) {
	t.Run("3 candidate election", func(t *testing.T) {
		candidates := []string{"Alice", "Carol", "Bob"}

		ballots := [][]string{
			{"Alice", "Bob", "Carol"},
			{"Alice", "Carol", "Bob"},
			{"Bob", "Alice", "Carol"},
			{"Bob", "Carol", "Alice"},
			{"Carol", "Bob", "Alice"},
		}

		winners, err := RunTidemanRankedPairsElection(candidates, ballots)
		if err != nil {
			t.Errorf("Error running tideman ranked pairs election: %v", err)
		}

		if !slices.Contains(winners, "Bob") {
			t.Errorf("winners = %v; want Bob", winners)
		}
	})

	t.Run("4 candidate election", func(t *testing.T) {
		candidates := []string{"Alice", "Carol", "Bob", "Dave"}

		ballots := [][]string{
			{"Alice", "Bob", "Carol", "Dave"},
			{"Alice", "Carol", "Bob", "Dave"},
			{"Bob", "Alice", "Carol", "Dave"},
			{"Bob", "Carol", "Alice", "Dave"},
			{"Carol", "Alice", "Bob", "Dave"},
			{"Carol", "Bob", "Alice", "Dave"},
			{"Dave", "Carol", "Alice", "Bob"},
			{"Dave", "Alice", "Carol", "Bob"},
		}

		winners, err := RunTidemanRankedPairsElection(candidates, ballots)
		if err != nil {
			t.Errorf("Error running tideman ranked pairs election: %v", err)
		}

		if !slices.Equal(winners, []string{"Alice", "Carol"}) {
			t.Errorf("winners = %v; want Alice", winners)
		}
	})

	t.Run("5 candidate election", func(t *testing.T) {
		candidates := []string{"Alice", "Carol", "Bob", "Dave", "Eve"}

		ballots := [][]string{
			{"Alice", "Bob", "Carol", "Dave", "Eve"},
			{"Alice", "Carol", "Bob", "Eve", "Dave"},
			{"Bob", "Alice", "Carol", "Dave", "Eve"},
			{"Bob", "Carol", "Alice", "Eve", "Dave"},
			{"Carol", "Alice", "Bob", "Dave", "Eve"},
			{"Carol", "Bob", "Alice", "Eve", "Dave"},
			{"Dave", "Eve", "Alice", "Bob", "Carol"},
			{"Eve", "Dave", "Alice", "Carol", "Bob"},
			{"Alice", "Bob", "Carol", "Eve", "Dave"},
			{"Bob", "Carol", "Alice", "Dave", "Eve"},
		}

		winners, err := RunTidemanRankedPairsElection(candidates, ballots)
		if err != nil {
			t.Errorf("Error running tideman ranked pairs election: %v", err)
		}

		if !slices.Contains(winners, "Alice") {
			t.Errorf("winners = %v; want Alice", winners)
		}
	})
}
