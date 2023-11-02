package easy

func KidsWithCandies(candies []int, extraCandies int) []bool {
	result := make([]bool, len(candies))
	max := 0

	for i := range candies {
		max = bigger(max, candies[i])
	}

	for i := range candies {
		if candies[i]+extraCandies >= max {
			result[i] = true
		}
	}

	return result
}

func bigger(a, b int) int {
	if a > b {
		return a
	}
	return b
}
