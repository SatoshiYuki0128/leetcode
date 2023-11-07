package medium

// 1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0

func LongestOnes(nums []int, k int) int {
	left, right := 0, 0

	for right < len(nums) {
		if nums[right] == 0 {
			k--
		}

		if k < 0 {
			if nums[left] == 0 {
				k++
			}
			left++
		}

		right++
	}

	return right - left
}
