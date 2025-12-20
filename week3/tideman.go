package week3

import (
	"slices"
	"sort"
)

type TidemanElection struct {
	Pairs []Pair
}

type Pair struct {
	First  string
	Second string
	Margin int
}

func RunTidemanRankedPairsElection(candidates []string, ballots [][]string) ([]string, error) {
	pairs := calculatePairMargins(candidates, ballots)
	pairs = sortPairMargins(pairs)
	lockedPairs := lockPairs(pairs)
	winners := findTidemanWinner(lockedPairs, candidates)

	return winners, nil
}

func calculatePairMargins(candidates []string, ballots [][]string) []Pair {
	var pairs []Pair
	for i := 0; i < len(candidates); i++ {
		if i == len(candidates)-1 {
			break
		}

		for j := i + 1; j < len(candidates); j++ {
			// candidates to compare
			c1, c2 := candidates[i], candidates[j]
			var count1, count2 int
			for _, ballot := range ballots {
				for _, candidate := range ballot {
					if candidate == c1 {
						count1++
						break
					}
					if candidate == c2 {
						count2++
						break
					}
				}
			}

			margin := count1 - count2
			if margin > 0 {
				pairs = append(pairs, Pair{First: c1, Second: c2, Margin: margin})
			}

			if margin < 0 {
				pairs = append(pairs, Pair{First: c2, Second: c1, Margin: -margin})
			}
		}
	}

	return pairs
}

func sortPairMargins(pairs []Pair) []Pair {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Margin > pairs[j].Margin
	})

	return pairs
}

func lockPairs(pairs []Pair) []Pair {
	var locked []Pair
	for _, pair := range pairs {
		// if there's a cycle, skip, else add to locked
		if !createsCycle(locked, pair.Second, []string{pair.First}) {
			locked = append(locked, pair)
		}
	}

	return locked
}

func createsCycle(lockedPairs []Pair, second string, cycle []string) bool {
	// cycle contains A
	for _, pair := range lockedPairs {
		// looking for a pair where B > C
		if pair.First == second {
			// if C is already in the cycle, return true
			if slices.Contains(cycle, pair.Second) {
				return true
			}

			// add B to the cycle
			cycle = append(cycle, pair.First)
			newCycle := make([]string, len(cycle))
			copy(newCycle, cycle)

			return createsCycle(lockedPairs, pair.Second, newCycle)
		}
	}

	return false
}

func findTidemanWinner(lockedPairs []Pair, candidates []string) []string {
	var winners []string
	for _, candidate := range candidates {
		hasLost := false
		for _, pair := range lockedPairs {
			if pair.Second == candidate {
				hasLost = true
				break
			}
		}

		if !hasLost {
			winners = append(winners, candidate)
		}
	}

	return winners
}
