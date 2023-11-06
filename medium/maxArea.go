package medium

import "math"

func MaxArea(height []int) int {
	size := len(height)

	if size < 2 {
		return 0
	}

	left, right := 0, size-1
	max := math.MinInt

	for left < right {
		max = bigger(max, (right-left)*smaller(height[left], height[right]))

		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}

	return max
}

func bigger(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func smaller(a, b int) int {
	if a < b {
		return a
	}
	return b
}
