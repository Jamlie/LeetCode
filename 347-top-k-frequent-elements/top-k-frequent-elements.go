func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)

	for _, num := range nums {
		freqMap[num]++
	}

	uniqueNums := make([]int, 0, len(freqMap))

	for num := range freqMap {
		uniqueNums = append(uniqueNums, num)
	}

	slices.SortFunc(uniqueNums, func(a, b int) int {
		return cmp.Compare(freqMap[b], freqMap[a])
	})

	return uniqueNums[:k]
}