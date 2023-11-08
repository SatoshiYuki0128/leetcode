package main

import (
	"fmt"
	"leetcode/easy"
)

func main() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{2, 4, 6}

	//nums1 := []int{1, 2, 3, 3}
	//nums2 := []int{1, 1, 2, 2}

	fmt.Println(easy.FindDifference(nums1, nums2))
}
