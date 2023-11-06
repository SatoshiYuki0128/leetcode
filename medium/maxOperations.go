package medium

import "sort"

func MaxOperations(nums []int, k int) int {
	size := len(nums)

	if size < 2 {
		return 0
	}

	sort.Ints(nums)

	left, right := 0, size-1
	count := 0

	for left < right {
		if nums[left]+nums[right] == k {
			count++
			left++
			right--
		} else if nums[left]+nums[right] > k {
			right--
		} else {
			left++
		}
	}

	return count
}
