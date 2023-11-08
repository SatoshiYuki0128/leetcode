package easy

//nums1 := []int{1, 2, 3}
//nums2 := []int{2, 4, 6}

func FindDifference(nums1 []int, nums2 []int) [][]int {
	record1 := map[int]int{}
	record2 := map[int]int{}

	for _, v := range nums1 {
		record1[v]++
	}

	for _, v := range nums2 {
		record2[v]++
	}

	resultMap1 := map[int]int{}
	resultMap2 := map[int]int{}

	for _, v := range nums1 {
		_, ok := record2[v]
		if !ok {
			resultMap1[v]++
		}
	}

	for _, v := range nums2 {
		_, ok := record1[v]
		if !ok {
			resultMap2[v]++
		}
	}

	result := make([][]int, 2)

	for k, _ := range resultMap1 {
		result[0] = append(result[0], k)
	}

	for k, _ := range resultMap2 {
		result[1] = append(result[1], k)
	}

	return result
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	return [][]int{setDifference(nums2, nums1), setDifference(nums1, nums2)}
}
func setDifference(nums1 []int, nums2 []int) []int {
	temp := [2001]bool{}
	var res []int
	for _, val1 := range nums1 {
		temp[val1+1000] = true
	}
	for _, val2 := range nums2 {
		if !temp[val2+1000] {
			res = append(res, val2)
			temp[val2+1000] = true
		}
	}
	return res
}
