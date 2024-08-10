func rob(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	if n == 1 {
		return nums[0]
	}

	first, second := nums[0], max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		first, second = second, max(second, first+nums[i])
	}

	return second
}