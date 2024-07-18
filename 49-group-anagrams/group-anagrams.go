func groupAnagrams(strs []string) [][]string {
	res := make([][]string, 0)
	anagrams := make(map[string][]string)

	for _, str := range strs {
		strsBytes := []byte(str)
		slices.Sort(strsBytes)
		tmpStr := string(strsBytes)

		if _, ok := anagrams[tmpStr]; !ok {
			anagrams[tmpStr] = make([]string, 0)
		}

		anagrams[tmpStr] = append(anagrams[tmpStr], str)
	}

	for _, v := range anagrams {
		res = append(res, v)
	}

	return res

}