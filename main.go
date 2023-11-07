package main

import (
	"fmt"
	"leetcode/medium"
)

func main() {
	nums := []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	k := 2

	fmt.Println(medium.LongestOnes(nums, k))
}
