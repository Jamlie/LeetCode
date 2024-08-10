func wordBreak(s string, wordDict []string) bool {
	wordsSet := make(map[string]bool)
	for _, word := range wordDict {
		wordsSet[word] = true
	}

	memo := make(map[string]bool)

	return canSegment(s, wordsSet, memo)
}

func canSegment(s string, wordsSet, memo map[string]bool) bool {
	if res, found := memo[s]; found {
		return res
	}

	if len(s) == 0 {
		return true
	}

	for word := range wordsSet {
		if strings.HasPrefix(s, word) {
			remaining := s[len(word):]
			if canSegment(remaining, wordsSet, memo) {
				memo[s] = true
				return true
			}
		}
	}

	memo[s] = false
	return false
}