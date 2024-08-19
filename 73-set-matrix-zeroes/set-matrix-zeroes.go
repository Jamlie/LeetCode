type Coordinate struct {
	row int
	col int
}

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}

	zeroesLocations := []Coordinate{}
	n := len(matrix)
	m := len(matrix[0])

	for i := range n {
		for j := range m {
			if matrix[i][j] == 0 {
				zeroesLocations = append(zeroesLocations, Coordinate{
					row: i,
					col: j,
				})
			}
		}
	}

	for _, location := range zeroesLocations {
		for i := range n {
			matrix[i][location.col] = 0
		}

		for j := range m {
			matrix[location.row][j] = 0
		}
	}
}