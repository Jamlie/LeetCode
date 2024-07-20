func longestConsecutive(nums []int) int {
    if len(nums) == 0 {
		return 0
	}

	numsSet := make(map[int]bool)

	for _, num := range nums {
		numsSet[num] = true
	}

	longestSequence := 0

	for num := range numsSet {
		if !numsSet[num-1] {
			currentNum := num
			currentSquenceLength := 1

			for numsSet[currentNum+1] {
				currentNum++
				currentSquenceLength++
			}

			longestSequence = max(currentSquenceLength, longestSequence)
		}
	}

	return longestSequence
}