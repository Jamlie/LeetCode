func kidsWithCandies(candies []int, extraCandies int) []bool {
	if len(candies) == 1 {
		return []bool{true}
	}

	kidsWithMostCandies := []bool{}

	maxNum := slices.Max(candies)

	for _, candy := range candies {
		newCandyCount := candy + extraCandies

		if newCandyCount < maxNum {
			kidsWithMostCandies = append(kidsWithMostCandies, false)
		} else {
			kidsWithMostCandies = append(kidsWithMostCandies, true)
		}
	}

	return kidsWithMostCandies
}
