type CoinGame struct {
	coins  []int
	amount int
	memo   map[int]int
}

func coinChange(coins []int, amount int) int {
	coinGame := CoinGame{
		coins:  coins,
		amount: amount,
		memo:   make(map[int]int),
	}

	return coinChangeRecursive(amount, coinGame)
}

func coinChangeRecursive(n int, coinGame CoinGame) int {
	if n == 0 {
		return 0
	}

	if n < 0 {
		return -1
	}

	if v, ok := coinGame.memo[n]; ok {
		return v
	}

	minCoins := coinGame.amount + 1
	for _, coin := range coinGame.coins {
		res := coinChangeRecursive(n-coin, coinGame)
		if res >= 0 && res < minCoins {
			minCoins = res + 1
		}
	}

	if minCoins == coinGame.amount+1 {
		coinGame.memo[n] = -1
	} else {
		coinGame.memo[n] = minCoins
	}

	return coinGame.memo[n]
}