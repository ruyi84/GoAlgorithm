package main

import "fmt"

func maxProfit(prices []int) int {
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

func main() {
	i := []int{7, 1, 5, 3, 6, 4}
	num := maxProfit(i)
	fmt.Println(num)
}
