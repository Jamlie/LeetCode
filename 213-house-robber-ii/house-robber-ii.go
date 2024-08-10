func rob(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	if n == 1 {
		return nums[0]
	}

	if n == 2 {
		return max(nums[0], nums[1])
	}

	return max(robHelper(nums[:n-1]), robHelper(nums[1:]))
}

func robHelper(nums []int) int {
	first, second := 0, 0
	for _, num := range nums {
		first, second = max(first, second+num), first
	}

	return first
}