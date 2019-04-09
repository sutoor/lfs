package stock

import (
	"errors"
)

//YesterdayMaxProfit returns -1 if no profit found & an error if prices is empty
//or if it contains invalid prices (eg a negative price)
func YesterdayMaxProfit(prices []int) (int, error) {
	if len(prices) < 1 {
		return 0, errors.New("no prices in the list")
	}
	if prices[0] < 0 {
		return 0, errors.New("invalid price found. must be >= 0")
	}
	minPrice := prices[0]
	maxPrice := prices[0]
	maxDiff := -1
	for _, price := range prices[1:] {
		if price < 0 {
			return 0, errors.New("invalid price found. must be >= 0")
		}
		if price < minPrice {
			minPrice = price
		}
		//only calc & update maxDiff when we see a bigger max
		//because we are only interested in minimums to the
		//left of max (i.e. must buy before sell condition)
		if price > maxPrice {
			maxPrice = price
			maxDiff = getMax(maxDiff, maxPrice-minPrice)
		}
	}
	return maxDiff, nil
}

func getMax(i1, i2 int) int {
	if i1 > i2 {
		return i1
	}
	return i2
}
