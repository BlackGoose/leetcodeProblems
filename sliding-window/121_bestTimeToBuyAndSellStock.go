package main

func maxProfit(prices []int) int {
	minPrice := prices[0]
	result := -1

	for _, price := range prices {
		if result < price-minPrice {
			result = price - minPrice
		}
		if price < minPrice {
			minPrice = price
		}
	}

	if result < 0 {
		return 0
	}
	return result
}
