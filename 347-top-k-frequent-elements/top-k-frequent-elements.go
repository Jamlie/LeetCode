type Item struct {
	key   int
	value int
}

func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)

	for _, num := range nums {
		freqMap[num]++
	}

	freqNums := make([]Item, 0, len(freqMap))
	for k, v := range freqMap {
		freqNums = append(freqNums, Item{
			key:   k,
			value: v,
		})
	}

	uniqueNums := make([]int, 0, len(freqNums))

	for {
		item := slices.MaxFunc(freqNums, func(a, b Item) int {
			return cmp.Compare(a.value, b.value)
		})

		uniqueNums = append(uniqueNums, item.key)

		freqNums = slices.DeleteFunc(freqNums, func(i Item) bool {
			return i == item
		})

		if len(uniqueNums) >= k {
			break
		}
	}

	return uniqueNums
}