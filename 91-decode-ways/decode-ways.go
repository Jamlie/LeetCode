func isZero(b byte) bool {
	return b == '0'
}

func isByteInRange(startRune, nextRune byte) bool {
	return startRune == '1' || startRune == '2' && nextRune >= '0' && nextRune <= '6'
}

func numDecodings(s string) int {
	n := len(s)
	dp := make([]int, n+1)
	dp[n] = 1

	for i := n - 1; i >= 0; i-- {
		if isZero(s[i]) {
			dp[i] = 0
		} else {
			dp[i] = dp[i+1]
		}

		if i+1 < n && isByteInRange(s[i], s[i+1]) {
			dp[i] += dp[i+2]
		}
	}

	return dp[0]
}