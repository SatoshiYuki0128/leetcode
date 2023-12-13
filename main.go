package main

import (
	"fmt"
	"leetcode/easy"
)

func main() {
	//num := [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}}
	//num := [][]int{{1, 0, 0}, {0, 0, 1}, {1, 0, 0}}
	num := [][]int{
		{0, 0, 0, 1},
		{1, 0, 0, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 0}}

	fmt.Println(easy.NumSpecial(num))
}
