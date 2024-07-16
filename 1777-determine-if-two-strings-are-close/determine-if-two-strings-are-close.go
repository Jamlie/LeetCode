import (
   	"slices"
	"sort"
)

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	charsOfFirstWord := make(map[rune]int)
	for _, c := range word1 {
		charsOfFirstWord[c]++
	}

	charsOfSecondWord := make(map[rune]int)
	for _, c := range word2 {
		charsOfSecondWord[c]++
	}

	keysOfWord1 := getKeys(charsOfFirstWord)
	keysOfWord2 := getKeys(charsOfSecondWord)
	if !slices.Equal(keysOfWord1, keysOfWord2) {
		return false
	}

	values1 := getValuesSorted(charsOfFirstWord)
	values2 := getValuesSorted(charsOfSecondWord)
	if !slices.Equal(values1, values2) {
		return false
	}

	return true
}

func getKeys(m map[rune]int) []rune {
	keys := make([]rune, 0, len(m))

	for key := range m {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	return keys
}

func getValuesSorted(m map[rune]int) []int {
	values := make([]int, 0, len(m))

	for _, v := range m {
		values = append(values, v)
	}

	sort.Ints(values)
	return values
}
