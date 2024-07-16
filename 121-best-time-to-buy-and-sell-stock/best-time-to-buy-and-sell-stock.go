func maxProfit(prices []int) int {
    if len(prices) == 0 {
		return 0
	}

	maxProfit := 0
	maxAmount := prices[0]

	for _, price := range prices {
		if price < maxAmount {
			maxAmount = price
			continue
		}

		currentProfit := price - maxAmount
		maxProfit = max(maxProfit, currentProfit)
	}

	return maxProfit

}