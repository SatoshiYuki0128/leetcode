package medium

import (
	"math"
	"strings"
)

func MaxVowels(s string, k int) int {
	vowels := "aeiou"

	left, right := 0, k-1
	max := math.MinInt
	count := 0

	for i := left; i <= right; i++ {
		if strings.Contains(vowels, string(s[i])) {
			count++
		}
	}

	for {
		max = big(max, count)

		if strings.Contains(vowels, string(s[left])) {
			count--
		}

		left++
		right++

		if right == len(s) {
			break
		}

		if strings.Contains(vowels, string(s[right])) {
			count++
		}
	}

	return max
}

func big(a, b int) int {
	if a > b {
		return a
	}
	return b
}
