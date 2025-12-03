package codetop1

// numIslands 岛屿数量
func numIslands(grid [][]byte) int {
	res := 0
	m, n := len(grid), len(grid[0])
	var dfs func(int, int)
	dfs = func(x, y int) {
		if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] != '1' {
			return
		}
		grid[x][y] = '2'
		// 左右上下
		dfs(x, y-1)
		dfs(x, y+1)
		dfs(x-1, y)
		dfs(x+1, y)
	}

	for i, row := range grid {
		for j, c := range row {
			if c == '1' {
				dfs(i, j)
				res++
			}
		}
	}
	return res
}
