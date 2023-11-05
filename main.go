package main

import (
	"fmt"
	"leetcode/easy"
)

func main() {
	//nums := []int{0, 1, 0, 3, 12}
	nums := []int{4, 2, 4, 0, 0, 3, 0, 5, 1, 0}

	easy.MoveZeroes(nums)

	fmt.Println(nums)
}
