func containsNearbyDuplicate(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}

	numberOfDups := make(map[int][]int)

	for i, num := range nums {
		numberOfDups[num] = append(numberOfDups[num], i)
	}

	for k, v := range numberOfDups {
		if len(v) < 2 {
			delete(numberOfDups, k)
		}
	}

	if len(numberOfDups) == 0 {
		return false
	}

	for _, v := range numberOfDups {
		for i := 0; i < len(v)-1; i++ {
			if v[i+1]-v[i] <= k {
				return true
			}
		}
	}

	return false
}