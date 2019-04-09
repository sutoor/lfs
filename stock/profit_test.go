package stock

import (
	"fmt"
	"testing"
)

func TestYesterdayMaxProfit(t *testing.T) {
	s := []int{9, 9, 20}
	res, err := YesterdayMaxProfit(s)
	if err != nil {
		fmt.Println("err:", err)
	} else {
		fmt.Println("res=", res)
	}
}
