package week3

import (
	"fmt"
	"slices"
)

func handleRunoffElection(candidates []string, votes [][]string) ([]string, error) {
	results, err := initializeCandidatesResults(candidates, 9)
	if err != nil {
		return nil, err
	}

	if err = validateBallots(votes, candidates); err != nil {
		return nil, err
	}

	results = allocateFirstVotes(results, votes)

	winningThreshold := len(votes) / 2

	winners := findWinner(results, winningThreshold)

	return winners, nil
}

func initializeCandidatesResults(candidates []string, maxCandidates int) (map[string][]string, error) {
	if len(candidates) > maxCandidates {
		return nil, fmt.Errorf("the maximum number of candidates is %d", maxCandidates)
	}

	results := make(map[string][]string)
	for _, candidate := range candidates {
		if _, ok := results[candidate]; ok {
			return nil, fmt.Errorf("candidate %s is duplicated", candidate)
		}
		results[candidate] = []string{}
	}

	return results, nil
}

func validateBallots(ballots [][]string, candidates []string) error {
	total := len(candidates)
	for _, ballot := range ballots {
		if len(ballot) != total {
			return fmt.Errorf("ballot is missing votes %v", ballot)
		}

		var s []string
		for _, v := range ballot {
			if !slices.Contains(candidates, v) {
				return fmt.Errorf("ballot was for an incorrect candidate: %v", ballot)
			}

			if slices.Contains(s, v) {
				return fmt.Errorf("ballot cannot vote for the same candidate twice: %v", ballot)
			}

			s = append(s, v)
		}
	}

	return nil
}

func allocateFirstVotes(results map[string][]string, ballots [][]string) map[string][]string {
	for _, ballot := range ballots {
		if _, ok := results[ballot[0]]; ok {
			results[ballot[0]] = append(results[ballot[0]], ballot[0])
		}
	}

	return results
}

func findWinner(results map[string][]string, winThreshold int) (winners []string) {
	if len(results) == 1 {
		return makeWinnersList(results)
	}

	var lowest int
	var relegatedCandidate string
	var counts []int
	for candidate, votes := range results {
		// found someone with majority
		if len(votes) >= winThreshold {
			return []string{candidate}
		}

		if len(votes) < lowest || lowest == 0 {
			lowest = len(votes)
			relegatedCandidate = candidate
		}

		counts = append(counts, len(votes))
	}

	// candidates tie for victory
	if counts[0] == len(counts)/len(results) {
		return makeWinnersList(results)
	}

	var lastWasCandidate bool
	for _, vote := range results[relegatedCandidate] {
		if lastWasCandidate {
			// can't assign to a candidate that was already knocked out, check if they're still in the running
			if currentVotes, ok := results[vote]; ok {
				results[vote] = append(currentVotes, results[relegatedCandidate]...)
				lastWasCandidate = false

				break
			}
			continue
		}

		if vote == relegatedCandidate {
			lastWasCandidate = true
		}
	}

	delete(results, relegatedCandidate)

	return findWinner(results, winThreshold)
}

func makeWinnersList(list map[string][]string) (winners []string) {
	for c, _ := range list {
		winners = append(winners, c)
	}
	return winners
}
