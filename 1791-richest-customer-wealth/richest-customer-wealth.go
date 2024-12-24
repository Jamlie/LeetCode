func maximumWealth(accounts [][]int) int {
	reducedSlice := make([]int, 0, len(accounts))

	sum := 0

	for _, account := range accounts {
		for _, v := range account {
			sum += v
		}
		reducedSlice = append(reducedSlice, sum)
		sum = 0
	}

	return slices.Max(reducedSlice)
}