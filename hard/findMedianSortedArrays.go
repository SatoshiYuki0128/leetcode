package hard

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	size1, size2 := len(nums1), len(nums2)

	if size1 == 0 {
		if size2%2 == 0 {
			return float64(nums2[size2/2-1]+nums2[size2/2]) / 2
		} else {
			return float64(nums2[size2/2])
		}
	}

	if size2 == 0 {
		if size1%2 == 0 {
			return float64(nums1[size1/2-1]+nums1[size1/2]) / 2
		} else {
			return float64(nums1[size1/2])
		}
	}

	indexM, indexN := 0, 0
	now := float64(0)

	if (size1+size2)%2 == 1 {
		for indexM < size1 && indexN < size2 {
			if nums1[indexM] <= nums2[indexN] {
				now = float64(nums1[indexM])
				indexM++
			} else {
				now = float64(nums2[indexN])
				indexN++
			}

			if indexM+indexN > (size1+size2)/2 {
				return now
			}
		}
	} else {
		medium := float64(0)

		for indexM < size1 && indexN < size2 {
			if nums1[indexM] <= nums2[indexN] {
				now = float64(nums1[indexM])
				indexM++
			} else {
				now = float64(nums2[indexN])
				indexN++
			}

			if indexM+indexN == (size1+size2)/2 {
				medium = now
			}

			if indexM+indexN == (size1+size2)/2+1 {
				return (medium + now) / 2
			}
		}
	}

	return 0
}
