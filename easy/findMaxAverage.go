package easy

func FindMaxAverage(nums []int, k int) float64 {
	left, right := 0, k-1
	max := -1e9
	sum := 0

	for i := 0; i < k; i++ {
		sum += nums[i]
	}

	for {
		max = big(max, float64(sum)/float64(k))

		sum -= nums[left]

		left++
		right++

		if right >= len(nums) {
			break
		}

		sum += nums[right]
	}

	return max
}

func big(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}
