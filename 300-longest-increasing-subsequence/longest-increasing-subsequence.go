func lengthOfLIS(nums []int) int {
	lis := make([]int, 0)

	for _, num := range nums {
		if indexOfNum, _ := slices.BinarySearch(lis, num); indexOfNum == len(lis) {
			lis = append(lis, num)
		} else {
			lis[indexOfNum] = num
		}
	}

	return len(lis)
}