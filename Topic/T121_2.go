package main

import "fmt"

func maxProfit2(prices []int) int {
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
	i := []int{1, 2}
	num := maxProfit2(i)
	fmt.Println(num)
}
