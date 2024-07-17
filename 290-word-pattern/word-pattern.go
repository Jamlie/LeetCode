func wordPattern(pattern string, s string) bool {
	words := strings.Fields(s)
	if len(pattern) == 1 && len(words) == 1 {
		return true
	}

	patterns := strings.Split(pattern, "")
	patternsSet := map[string]bool{}

	for _, ptrn := range patterns {
		patternsSet[ptrn] = true
	}

	wordsSet := map[string]bool{}
	for _, word := range words {
		wordsSet[word] = true
	}

	if len(words) != len(patterns) {
		return false
	}

	if len(patternsSet) == 1 && len(wordsSet) == 1 {
		return true
	}

	tempMatches := make(map[string]string)

	for i, word := range words {
		tempMatches[word] = patterns[i]
	}

	if len(tempMatches) != len(patternsSet) {
		return false
	}

	matches := make(map[string]string)

	for i, ptrn := range patterns {
		matches[ptrn] = words[i]
	}

	wordsBuilder := strings.Builder{}

	for i, ptrn := range patterns {
		_, err := wordsBuilder.WriteString(matches[ptrn])
		if err != nil {
			return false
		}

		if i != len(patterns)-1 {
			_, err = wordsBuilder.WriteRune(' ')
			if err != nil {
				return false
			}
		}
	}

	if s == wordsBuilder.String() {
		return true
	}

	return false
}