package main

import (
	"fmt"
	"unicode"
)

func main() {
}

func fmtin() {
	fmt.Println()
}

func maxProfit(prices []int) int {
	buy_idx := 0
	sell_idx := 1
	max_profit := 0

	for sell_idx < len(prices) {
		buy := prices[buy_idx]
		sell := prices[sell_idx]

		if sell < buy {
			buy_idx = sell_idx
			sell_idx++
			continue
		}

		profit := sell - buy
		if max_profit < profit {
			max_profit = profit
		}
		sell_idx++
	}

	return max_profit
}
