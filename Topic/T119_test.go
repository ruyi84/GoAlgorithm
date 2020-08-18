package main

import (
	"fmt"
	"testing"
)

// 给定一个非负索引 k，其中 k ≤ 33，返回杨辉三角的第 k 行。

/*
反向计算，根据当前k值计算前一行的值，如果
*/

func getRow2(rowIndex int) []int {
	ints := []int{}

	for i := 0; i <= rowIndex; i++ {
		if i == 0 || i == rowIndex {
			ints = append(ints, 1)
			continue
		}
		if i > rowIndex/2 {
			ints = append(ints, ints[rowIndex-i])
			continue
		}
		ints = append(ints, t119(i, rowIndex+1))
	}
	return ints
}

// 用于递归计算
func t119(rowIndex, nowIndex int) int {
	if rowIndex == 0 || rowIndex == nowIndex-1 {
		return 1
	}

	return t119(rowIndex-1, nowIndex-1) + t119(rowIndex, nowIndex-1)
}

// 在注释里看到的计算方式，倒着计算，不是很能理解
/*
遍历，根据长度在末尾插入1
然后计算前几位，每一位的实际值应该是之前的值加前一个的值
*/
func getRow(rowIndex int) []int {
	nums := []int{1}
	for i := 1; i <= rowIndex; i++ {
		nums = append(nums, 1)
		for j := i - 1; j > 0; j-- {
			nums[j] += nums[j-1]
		}

	}
	return nums
}

func Test_T119(t *testing.T) {
	fmt.Println(getRow(5))
}
