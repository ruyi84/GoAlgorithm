package main

import "fmt"

func maxProfit2(prices []int) int {
	max := 0
	for i := 0; i < len(prices)-1; i++ {
		for j := i + 1; j < len(prices); j++ {
			num := (prices[j] - prices[i])
			if max < num {
				max = num
			}
		}
	}
	return max
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	min, max := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if min > prices[i] {
			min = prices[i]
		}
		if max < (prices[i] - min) {
			max = (prices[i] - min)
		}
	}
	return max
}

func main() {
	i := []int{7, 1, 5, 3, 6, 4}
	num := maxProfit(i)
	fmt.Println(num)
}
