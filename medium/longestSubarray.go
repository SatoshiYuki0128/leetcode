package medium

import "fmt"

// 1, 1, 0, 1
// 0,1,1,1,0,1,1,0,1

// 00 01 02 03 04 14 15

func LongestSubarray(nums []int) int {
	left, right := 0, 0
	count := 0
	result := 0

	for right < len(nums) {
		if count < 2 {
			if nums[right] == 0 {
				count++
			}

			if count < 2 && right-left > result {
				result = right - left
				fmt.Println(left, right, result, count)
			}

			right++
		} else {
			if nums[left] == 0 {
				count--
			}

			left++
		}
	}

	return result
}
