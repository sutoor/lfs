package stock

import (
	"testing"
)

func TestNoProfit(t *testing.T) {
	prices := []int{13, 8, 4, 3, 1}
	res, err := YesterdayMaxProfit(prices)
	if res != 0 {
		t.Fail()
		t.Log("got:", res, ", want: 0 as there is no profit to be made")
	}
	if err != nil {
		t.Fail()
		t.Log("got err:", err, ", want: no error")
	}

}

func TestEmptyPrices(t *testing.T) {
	prices := []int{}
	_, err := YesterdayMaxProfit(prices)
	if err == nil {
		t.Fail()
		t.Log("got: no error, want: error")
	}
}

func TestInvalidPrice(t *testing.T) {
	prices := []int{9, 2, -3, 4, 56}
	_, err := YesterdayMaxProfit(prices)
	if err == nil {
		t.Fail()
		t.Log("got: no error, want: error")
	}
}

func TestMaxProfit(t *testing.T) {
	prices := []int{90, 100, 2, 25, 3, 4, 27, 26}
	res, err := YesterdayMaxProfit(prices)
	if err != nil {
		t.Fail()
		t.Log("got err:", err, ", want: no error")
	}
	//want 25 (2->27) and NOT:
	//98 (100->2): bigger diff but doesnt satisfy buy before sell
	//10 (90->100): only satisfies maxPrice-minPrice == maxDiff (not necessarily true) because
	//				although max - min will give max diff, but given the buy before sell condition,
	//				at a later time, smaller stock prices could yield bigger diff (i.e. higher profit)
	if res != 25 {
		t.Fail()
		t.Log("got:", res, ", want: 25")
	}
}
