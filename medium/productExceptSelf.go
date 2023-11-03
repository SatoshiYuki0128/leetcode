package medium

func ProductExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	total := 1
	var zero []int

	for i, v := range nums {
		if v == 0 {
			zero = append(zero, i)
		} else {
			total *= v
		}

		if len(zero) > 1 {
			return result
		}
	}

	for i := range nums {
		if len(zero) > 0 {
			if i == zero[0] {
				result[i] = total
			} else {
				result[i] = 0
			}
		} else {
			result[i] = total / nums[i]
		}
	}

	return result
}

func ProductExceptSelf2(nums []int) []int {
	prod := make([]int, len(nums))
	revprod := make([]int, len(nums))

	cmult1 := 1
	cmult2 := 1

	for i := 0; i < len(nums); i++ {
		cmult1 = cmult1 * nums[i]
		prod[i] = cmult1

		revindex := len(nums) - 1 - i
		cmult2 = cmult2 * nums[revindex]
		revprod[revindex] = cmult2
	}

	ans := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		left := 1
		right := 1

		if i > 0 {
			left = prod[i-1]
		}

		if i < len(nums)-1 {
			right = revprod[i+1]
		}
		ans[i] = left * right
	}

	return ans
}
