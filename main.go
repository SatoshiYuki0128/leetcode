package main

import (
	"fmt"
	"leetcode/medium"
)

func main() {
	chars := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c', 'd'}

	fmt.Println(medium.Compress(chars))
}
