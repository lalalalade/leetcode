package codetop4

import "slices"

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := range n {
		for j := range i {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	for _, row := range matrix {
		slices.Reverse(row)
	}
}
