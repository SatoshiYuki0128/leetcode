package easy

// 1, 7, 3, 6, 5, 6

func PivotIndex(nums []int) int {
	sum := 0

	for _, v := range nums {
		sum += v
	}

	//fmt.Println(sum)

	count := 0

	for i, v := range nums {
		if count == sum-v {
			return i
		}

		count += v
		sum -= v
	}

	return -1
}
