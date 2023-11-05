package medium

import "math"

func IncreasingTriplet(nums []int) bool {
	size := len(nums)

	if size < 3 {
		return false
	}

	max := -1 << 31
	maxArray := make([]int, len(nums))

	min := (1 << 31) - 1
	minArray := make([]int, len(nums))

	for i := range nums {
		if nums[i] < min {
			min = nums[i]
		}

		if nums[size-1-i] > max {
			max = nums[size-1-i]
		}

		minArray[i] = min
		maxArray[size-1-i] = max
	}

	//fmt.Println(minArray)
	//fmt.Println(maxArray)

	for i := range nums {
		if minArray[i] < nums[i] && maxArray[i] > nums[i] {
			return true
		}
	}

	return false
}

func IncreasingTriplet2(nums []int) bool {
	a, b, n := math.MaxInt, math.MaxInt, len(nums)
	for i := 0; i < n; i++ {
		num := nums[i]
		if a >= num {
			a = num
		} else if b >= num {
			b = num
		} else {
			return true
		}
	}
	return false
}
