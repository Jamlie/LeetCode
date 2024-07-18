func canJump(nums []int) bool {
	if len(nums) == 0 {
		return false
	}

	maxReach := 0
	for i, num := range nums {
		if i > maxReach {
			return false
		}

		maxReach = max(maxReach, i+num)
		if maxReach >= len(nums)-1 {
			return true
		}
	}

	return false
}