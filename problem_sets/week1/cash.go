package week1

func coinsOwed(centsOwed int) int {
	if centsOwed <= 0 {
		return 0
	}

	coinTypes := []int{25, 10, 5, 1}
	var total int

	for _, coin := range coinTypes {
		coinsToAdd := centsOwed / coin
		total += coinsToAdd

		centsOwed -= coinsToAdd * coin

		if centsOwed <= 0 {
			break
		}
	}

	return total
}
