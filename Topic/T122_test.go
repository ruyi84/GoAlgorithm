package main

import (
	"fmt"
	"testing"
)

/*
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。

注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
/*
两个指针，第一个指针固定后，另一个指针向后移动，如果比第一个大就做差，然后计算后面的数组
时间太长，不能成功计算出结果
*/
func maxProfit_122(prices []int) int {
	num, max := 0, 0
	for k := range prices {
		for l := k + 1; l <= len(prices)-1; l++ {
			if prices[l] <= prices[k] {
				continue
			}
			num += prices[l] - prices[k] + maxProfit_122(prices[l+1:])
			if num > max {
				max = num
			}
			num = 0
		}

	}
	return max
}

/*
实际的计算就是计算所有的差值，找出所有的低到高的和
*/
func maxProfit_112_2(prices []int) int {
	len := len(prices)
	profit := 0
	min := prices[0]
	for i := 1; i < len; i++ {
		if prices[i] > min {
			profit += prices[i] - min
		}
		min = prices[i]
	}
	return profit
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	sum := 0
	price := -1
	for i := range prices {
		if price != -1 && (i == len(prices)-1 || prices[i] > prices[i+1]) {
			sum += prices[i] - price
			price = -1
		} else if price == -1 && (i != len(prices)-1 && prices[i] < prices[i+1]) {
			price = prices[i]
		}
	}
	return sum
}

func Test_T122(t *testing.T) {
	ints := []int{1, 4, 4, 3, 2, 4}
	fmt.Println(maxProfit(ints))
}
