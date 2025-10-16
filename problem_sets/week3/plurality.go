package week3

import (
	"fmt"
)

func handlePluralityElection(candidates []string, votes []string) ([]string, error) {
	ballot, err := initializePluralityBallot(candidates, 9)
	if err != nil {
		return nil, err
	}

	results := countVotes(ballot, votes)
	winners := selectWinners(results)

	return winners, nil
}

func initializePluralityBallot(candidates []string, maxCandidates int) (map[string]int, error) {
	if len(candidates) > maxCandidates {
		return nil, fmt.Errorf("the maximum number of candidates is %d", maxCandidates)
	}

	ballot := make(map[string]int)
	for _, candidate := range candidates {
		if _, ok := ballot[candidate]; ok {
			return nil, fmt.Errorf("candidate %s is duplicated", candidate)
		}
		ballot[candidate] = 0
	}

	return ballot, nil
}

func countVotes(ballot map[string]int, votes []string) map[string]int {
	for _, vote := range votes {
		if _, ok := ballot[vote]; ok {
			ballot[vote]++
		}
	}

	return ballot
}

func selectWinners(results map[string]int) []string {
	var winners []string
	var leadingTotal int
	for candidate, result := range results {
		if result == leadingTotal {
			winners = append(winners, candidate)
		}

		if result > leadingTotal {
			leadingTotal = result
			winners = nil
			winners = append(winners, candidate)
		}
	}

	return winners
}
