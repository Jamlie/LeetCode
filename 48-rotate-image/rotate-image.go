func rotate(matrix [][]int) {
	n := len(matrix)

	for i := range n {
		for j := i + 1; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}

        slices.Reverse(matrix[i])
	}
}