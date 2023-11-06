package main

import (
	"fmt"
	"leetcode/medium"
)

func main() {
	//nums := []int{1, 2, 3, 4}
	//k := 5

	nums := []int{3, 1, 3, 4, 3}
	k := 6

	fmt.Println(medium.MaxOperations(nums, k))
}
