package code_top

import "math"

// dirs 右下左上 四个方向
var dirs = [4][2]int{
	{0, 1}, {1, 0}, {0, -1}, {-1, 0},
}

// spiralOrder 螺旋矩阵
func spiralOrder(matrix [][]int) []int {
	m, n := len(matrix), len(matrix[0])
	res := make([]int, m*n)
	i, j, di := 0, 0, 0
	for k := range res {
		res[k] = matrix[i][j]
		matrix[i][j] = math.MaxInt
		// 下一步的位置
		x, y := i+dirs[di][0], j+dirs[di][1]
		if x < 0 || x >= m || y < 0 || y >= n || matrix[x][y] == math.MaxInt {
			di = (di + 1) % 4
		}
		i += dirs[di][0]
		j += dirs[di][1]
	}
	return res
}
