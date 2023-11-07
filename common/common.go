package common

import "fmt"

func PrintIntMatrix(matrix [][]int) {
	for i := range matrix {
		for j := range matrix[i] {

			fmt.Printf("%d ", matrix[i][j])

		}
		fmt.Println()
	}
}
