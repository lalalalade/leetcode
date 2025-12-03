package codetop4

func maxAreaOfIsland(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res := 0
	var dfs func(int, int) int
	dfs = func(i, j int) int {
		area := 1
		grid[i][j] = 0
		for _, d := range dirs {
			x, y := i+d.x, j+d.y
			if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == 0 {
				continue
			}
			area += dfs(x, y)
		}
		return area
	}

	for i, row := range grid {
		for j, x := range row {
			if x > 0 {
				res = max(res, dfs(i, j))
			}
		}
	}

	return res
}

var dirs = []struct{ x, y int }{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
