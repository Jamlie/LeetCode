func combinationSum4(nums []int, target int) int {
	return findCombinations(nums, target, make(map[int]int))
}

func findCombinations(nums []int, target int, memo map[int]int) int {
	if target == 0 {
		return 1
	}

	if target < 0 {
		return 0
	}

	if v, ok := memo[target]; ok {
		return v
	}

	result := 0
	for _, num := range nums {
		result += findCombinations(nums, target-num, memo)
	}

    memo[target] = result

	return memo[target]
}