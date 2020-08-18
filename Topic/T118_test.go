package main

import (
	"fmt"
	"testing"
)

/*
给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。
*/
/*
杨辉三角的概念，每一行的数目比前一行多以，每一个格子的数值为前一行的两个肩膀数值的和
*/
func generate2(numRows int) [][]int {
	ints := [][]int{}
	if numRows == 0 {
		return ints
	}
	ints = append(ints, []int{1})
	if numRows == 1 {
		return ints
	}
	for i := 1; i < numRows; i++ {
		ints = append(ints, []int{})
		ints[i] = append(ints[i], 1)
		if i > 1 {
			for j := 1; j <= i-1; j++ {
				ints[i] = append(ints[i], ints[i-1][j-1]+ints[i-1][j])
			}
		}
		ints[i] = append(ints[i], 1)
	}

	return ints
}

func generate(numRows int) [][]int {
	ints := [][]int{}
	if numRows == 0 {
		return ints
	}

	for i := 0; i < numRows; i++ {

	}
}

func Test_T118(t *testing.T) {
	ints := generate(5)
	for k := range ints {
		fmt.Println(ints[k])
	}
}
