package feed

func isBlacklisted(blacklist []int, ingredients []int) bool {
	i, j := 0, 0
	for i < len(blacklist) && j < len(ingredients) {
		if blacklist[i] == ingredients[j] {
			return true
		}
		if blacklist[i] < ingredients[j] {
			i++
			continue
		}
		j++

	}
	return false
}

func calculateScore(allowList []int, ingredients []int) int {
	score := 0
	i, j := 0, 0
	for i < len(allowList) && j < len(ingredients) {
		if allowList[i] == ingredients[j] {
			i++
			j++
			score++
			continue
		}
		if allowList[i] < ingredients[j] {
			i++
			continue
		}
		j++
	}
	return score
}
