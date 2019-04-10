package stock

import (
	"testing"
)

func TestNoProfit(t *testing.T) {
	prices := []int{13, 8, 4, 3, 1}
	res, err := YesterdayMaxProfit(prices)
	if res != -1 {
		t.Fail()
		t.Log("got:", res, ", want: -1 as there is no profit to be made")
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
	prices := []int{10, 5, 7, 80, 1, 11, 2, 9}
	res, err := YesterdayMaxProfit(prices)
	if err != nil {
		t.Fail()
		t.Log("got err:", err, ", want: no error")
	}
	//want 75 (5 -> 80), not 79 (80 -> 1) or any other number
	if res != 75 {
		t.Fail()
		t.Log("got:", res, ", want: 75")
	}
}
