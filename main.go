package main

import (
	"fmt"
	"leetcode/medium"
)

func main() {
	//nums := []int{1, 0, 1, 0}
	nums := []int{0, 1, 1, 1, 0, 1, 1, 0, 1}

	fmt.Println(medium.LongestSubarray(nums))
}
