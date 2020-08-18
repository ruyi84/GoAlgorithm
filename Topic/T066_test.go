package main

import (
	"fmt"
	"testing"
)

/*
给定一个由整数组成的非空数组所表示的非负整数，在该数的基础上加一。

最高位数字存放在数组的首位， 数组中每个元素只存储单个数字。

你可以假设除了整数 0 之外，这个整数不会以零开头。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/plus-one
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
第一时间读题，出现偏差，以为只给最后一位添加就可以，结果提交出现了最后一位是9的问题，啪啪打脸。
第二时间还是没当真，就做个遍历，如果当前位加完后出现了10就进一位1，结果忘记了全是9，需要给数组头部添加的问题了，又一次打脸。
最后写了如下代码，如果当前位计算完为10，就需要进1，且有num作为标志位，如果循环结束，num仍为1，则需要在原数组上进一位，之后又新增了每次循环开始判断是否为0，若为0则说明不再需要进位计算，直接跳出，节省了一定的时间消耗。
*/

func plusOne2(digits []int) []int {
	if len(digits) == 0 {
		return []int{1}
	}

	num := 1
	for k := range digits {
		if num == 0 {
			return digits
		}
		digits[len(digits)-k-1] += num
		num = 0
		if digits[len(digits)-k-1] == 10 {
			digits[len(digits)-k-1] = 0
			num = 1
		}

	}
	if num == 1 {
		return append([]int{1}, digits...)
	}
	return digits
}

// 后来考虑，其实如果当前值不加为0可以直接跳出，继续修改为倒叙遍历，当前值为9则置为0进位运算，如果不为9就++返回，如果一直未返回则需要进位在头部+1

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return digits
	}

	for i := len(digits); i > 0; i-- {
		if digits[i-1] != 9 {
			digits[i-1]++
			return digits
		} else {
			digits[i-1] = 0
		}
	}
	return append([]int{1}, digits...)
}

func Test_T066(t *testing.T) {
	fmt.Println(plusOne([]int{0}))
}
