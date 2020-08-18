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
*/
func maxProfit_122(prices []int) int {

	num, max := 0, 0
	for k := range prices {
		for l := k + 1; k < len(prices)-1; k++ {
			if prices[l] > prices[k] {
				num += prices[l] - prices[k] + maxProfit_122(prices[l+1:])
				fmt.Println(prices[l]-prices[k], maxProfit_122(prices[l+1:]), num)
			}
			if num > max {
				//fmt.Println(prices[k],prices[l])
				max = num

			}
			num = 0
		}

	}
	return max
}

func Test_T122(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5}
	fmt.Println(maxProfit_122(ints))
}
