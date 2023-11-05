package easy

func MoveZeroes(nums []int) {
	size := len(nums)

	if size < 2 {
		return
	}

	left, right := 0, 1

	for right < size {
		if nums[left] == 0 {
			if nums[right] != 0 {
				nums[left], nums[right] = nums[right], nums[left]
				left++
				right = left + 1
			} else {
				right++
			}
		} else {
			left++
			right++
		}

		if left >= size {
			break
		}
	}
}

func moveZeroes(nums []int) {
	l := 0

	for r := range nums {
		if nums[r] != 0 {
			nums[l] = nums[r]
			l++
		}
	}

	for ; l < len(nums); l++ {
		nums[l] = 0
	}

	return

}
