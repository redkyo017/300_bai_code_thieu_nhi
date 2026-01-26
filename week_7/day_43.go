package week7

// 200. Number of Islands https://leetcode.com/problems/number-of-islands/description/

func NumIslands(grid [][]byte) int {
	ROWS := len(grid)
	COLS := len(grid[0])
	visited := make(map[[2]int]bool, 0)
	res := 0
	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r == ROWS || r < 0 || c == COLS || c < 0 || grid[r][c] == '0' || visited[[2]int{r, c}] {
			return
		}
		visited[[2]int{r, c}] = true
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r+1, c)
		dfs(r, c-1)
	}
	for i := 0; i < ROWS; i++ {
		for j := 0; j < COLS; j++ {
			if !visited[[2]int{i, j}] && grid[i][j] == '1' {
				dfs(i, j)
				res++
			}
		}
	}
	return res
}
