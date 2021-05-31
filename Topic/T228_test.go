package main

import (
	"fmt"
	"testing"
)

/*
给定一个无重复元素的有序整数数组 nums 。

返回 恰好覆盖数组中所有数字 的 最小有序 区间范围列表。也就是说，nums 的每个元素都恰好被某个区间范围所覆盖，并且不存在属于某个范围但不属于 nums 的数字 x 。

列表中的每个区间范围 [a,b] 应该按如下格式输出：

"a->b" ，如果 a != b
"a" ，如果 a == b

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/summary-ranges
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func summaryRanges(nums []int) []string {
	var result []string
	for i := 0; i < len(nums); {
		index := i
		for j := i + 1; j < len(nums); j++ {
			if nums[index]+1 == nums[j] {
				index++
				continue
			}
			break
		}
		if index == i {
			num := fmt.Sprintf("%d", nums[i])
			result = append(result, num)
			i++
		} else {
			num := fmt.Sprintf("%d->%d", nums[i], nums[index])
			result = append(result, num)
			i = index + 1
		}
	}
	return result
}

func Test228(t *testing.T) {
	fmt.Println(summaryRanges([]int{1, 2, 4, 5, 6, 10}))
}
