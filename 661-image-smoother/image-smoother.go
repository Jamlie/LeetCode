func imageSmoother(img [][]int) [][]int {
	if len(img) == 0 {
		return nil
	}

	m, n := len(img), len(img[0])
	result := make([][]int, m)
	for i := range m {
		result[i] = make([]int, n)
	}

	directions := [...][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 0},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}

	for i := range m {
		for j := range n {
			count, total := 0, 0

			for _, v := range directions {
				directionI, directionJ := v[0], v[1]
				ni, nj := i+directionI, j+directionJ
				if 0 <= ni && ni < m && 0 <= nj && nj < n {
					total += img[ni][nj]
					count += 1
				}
                if count != 0 {
				    result[i][j] = total / count
                }
			}
		}
	}

	return result
}