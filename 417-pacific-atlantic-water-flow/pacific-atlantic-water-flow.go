type PacificAtlantic struct {
	heights     [][]int
	rows        int
	cols        int
	pacificSet  map[[2]int]bool
	atlanticSet map[[2]int]bool
}

func (pa *PacificAtlantic) DFS(r, c int, visit map[[2]int]bool, prevHeight int) {
	if r < 0 || c < 0 || r >= pa.rows || c >= pa.cols || visit[[2]int{r, c}] ||
		pa.heights[r][c] < prevHeight {
		return
	}

	visit[[2]int{r, c}] = true

	pa.DFS(r+1, c, visit, pa.heights[r][c])
	pa.DFS(r-1, c, visit, pa.heights[r][c])
	pa.DFS(r, c+1, visit, pa.heights[r][c])
	pa.DFS(r, c-1, visit, pa.heights[r][c])
}

func pacificAtlantic(heights [][]int) [][]int {
	if len(heights) == 0 {
		return nil
	}

	var (
		rows        = len(heights)
		cols        = len(heights[0])
		pacificSet  = make(map[[2]int]bool)
		atlanticSet = make(map[[2]int]bool)
		pa          = PacificAtlantic{
			heights:     heights,
			rows:        rows,
			cols:        cols,
			pacificSet:  pacificSet,
			atlanticSet: atlanticSet,
		}
	)

	for col := range cols {
		pa.DFS(0, col, pacificSet, heights[0][col])
		pa.DFS(rows-1, col, atlanticSet, heights[rows-1][col])
	}

	for row := range rows {
		pa.DFS(row, 0, pacificSet, heights[row][0])
		pa.DFS(row, cols-1, atlanticSet, heights[row][cols-1])
	}

	res := [][]int{}

	for row := range rows {
		for col := range cols {
			if pacificSet[[2]int{row, col}] && atlanticSet[[2]int{row, col}] {
				res = append(res, []int{row, col})
			}
		}
	}

	return res
}