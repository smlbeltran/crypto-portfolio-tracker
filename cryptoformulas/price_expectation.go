package cryptoformulas

func ProfitPercentageExpectation(currentPrice, targetPrice float64) float64 {
	return ((targetPrice - currentPrice) / currentPrice) * 100
}

func ProfitExpectation(totalCoins, targetPrice float64) float64 {
	return totalCoins * targetPrice
}
