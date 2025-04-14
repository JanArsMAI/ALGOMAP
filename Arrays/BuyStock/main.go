package main

import "fmt"

func maxProfit(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}
	max_profit := 0
	min_price := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < min_price {
			min_price = prices[i]
		}
		profit := prices[i] - min_price
		if profit > max_profit {
			max_profit = profit
		}
	}
	return max_profit
}

func main() {
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))
}
