func romanToInt(roman string) int {
	romanNumbers := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0

	for i, c := range roman {
		value := romanNumbers[byte(c)]

		if i < len(roman)-1 && romanNumbers[roman[i+1]] > value {
			result -= value
		} else {
			result += value
		}
	}

	return result
}
