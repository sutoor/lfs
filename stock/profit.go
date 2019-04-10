package stock

import (
	"errors"
)

//YesterdayMaxProfit returns 0 if no profit is to be made & an error if prices is empty
//or if it contains invalid value(s) (i.e. a negative price)
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
		//always update the maxDiff if the current diff is greater
		//because we either saw a new min or current price is farther away from
		//existing min than maxPrice-minPrice
		maxDiff = getMax(maxDiff, price-minPrice)
		//only update maxDiff when we see a bigger max
		//because we are interested in the minimum to the left
		//of this max (i.e. must buy before sell condition)
		if price >= maxPrice {
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
