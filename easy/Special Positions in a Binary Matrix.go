package easy

import "fmt"

func NumSpecial(mat [][]int) int {
	if len(mat) == 0 {
		return 0
	}

	rowSize := len(mat)
	columnSize := len(mat[0])

	row := make([]int, rowSize)
	column := make([]int, columnSize)

	for i := 0; i < rowSize; i++ {
		for j := 0; j < columnSize; j++ {
			if mat[i][j] == 1 {
				row[i]++
			}
		}
	}

	for i := 0; i < columnSize; i++ {
		for j := 0; j < rowSize; j++ {
			if mat[j][i] == 1 {
				column[i]++
			}
		}
	}

	fmt.Println(row)
	fmt.Println(column)

	count := 0

	for i := 0; i < rowSize; i++ {
		for j := 0; j < columnSize; j++ {
			if mat[i][j] == 1 && row[i] == 1 && column[j] == 1 {
				count++
			}
		}
	}

	return count
}
