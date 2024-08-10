type Matrix struct {
	maze    [][]int
	rows    int
	columns int
}

func NewMatrix(rows, columns int) *Matrix {
	maze := make([][]int, rows)

	for i := range maze {
		maze[i] = make([]int, columns)
	}

	return &Matrix{
		maze:    maze,
		rows:    rows,
		columns: columns,
	}
}

func (m Matrix) RowsAsIndex() int {
	return m.rows - 1
}

func (m Matrix) ColumnAsIndex() int {
	return m.columns - 1
}

func (m *Matrix) findUniquePaths(i, j int) int {
	if i >= m.rows || j >= m.columns {
		return 0
	}

	if i == m.RowsAsIndex() || j == m.ColumnAsIndex() {
		return 1
	}

	if m.maze[i][j] > 0 {
		return m.maze[i][j]
	}

	m.maze[i][j] = m.findUniquePaths(i+1, j) + m.findUniquePaths(i, j+1)
	return m.maze[i][j]
}

func uniquePaths(m, n int) int {
	maze := NewMatrix(m, n)

	return maze.findUniquePaths(0, 0)
}