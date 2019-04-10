package main

import (
	"flag"
	"fmt"
	"lfs/stock"
	"os"
	"strconv"
	"strings"
)

var maxprofitCln = "maxprofit"

func main() {
	flag.Usage = func() {
		fmt.Println("Usage:\n  lfs <subcommand> [options]")
		fmt.Printf("\nAvailable subcommands: \n  o %s", maxprofitCln)
		fmt.Println("\n\nSee 'lfs <subcommand> -h' for usage")
	}

	maxprofitCmd := flag.NewFlagSet(maxprofitCln, flag.ExitOnError)
	stockPrices := maxprofitCmd.String("prices", "100,7,5,8,11,9", "comma separated list (without white space) of stock prices")
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("subcommand required. See 'lfs -h' for help and usage")
		os.Exit(1)
	}
	switch os.Args[1] {
	case maxprofitCln:
		maxprofitCmd.Parse(os.Args[2:])
	default:
		fmt.Printf("%q is not a valid command\n", os.Args[1])
		os.Exit(2)
	}

	if maxprofitCmd.Parsed() {
		priceList := strings.Split(*stockPrices, ",")
		if res, err := getInts(priceList); err != nil {
			fmt.Println("unable to parse prices from command line")
		} else {
			if r, err := stock.YesterdayMaxProfit(res); err != nil {
				fmt.Println("err calculating max profit:", err)
			} else {
				fmt.Println("max profit:", r)
			}
		}
	}
}

func getInts(s []string) ([]int, error) {
	res := make([]int, 0)
	for _, v := range s {
		if r, err := strconv.Atoi(v); err != nil {
			return nil, err
		} else {
			res = append(res, r)
		}
	}
	return res, nil
}
