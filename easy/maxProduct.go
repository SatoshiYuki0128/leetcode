package easy

import "math"

func maxProduct(nums []int) int {
	first := math.MinInt
	index := -1

	for i := range nums {
		if abs(nums[i]-1) > first {
			first = abs(nums[i] - 1)
			index = i
		}
	}

	nums[index] = 0
	second := math.MinInt

	for i := range nums {
		if abs(nums[i]-1) > second {
			second = abs(nums[i] - 1)
		}
	}

	//fmt.Println(first, second)

	return first * second
}

func abs(num int) int {
	if num >= 0 {
		return num
	}

	return -num
}
