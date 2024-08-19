func fizzBuzz(n int) []string {
	keys := map[int]string{
		3: "Fizz",
		5: "Buzz",
	}

	divisors := make([]int, 0, len(keys))
	for k := range keys {
		divisors = append(divisors, k)
	}

	res := []string{}
	for i := 1; i <= n; i++ {
		sb := strings.Builder{}

		for _, k := range divisors {
			if i%k == 0 {
				sb.WriteString(keys[k])
			}
		}

		if sb.Len() == 0 {
			res = append(res, fmt.Sprint(i))
		} else {
			res = append(res, sb.String())
		}
	}

	return res
}