package main

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"
)

/*
题目
给定一个 24 小时制（小时:分钟）的时间列表，找出列表中任意两个时间的最小时间差并已分钟数表示。


输入: ["23:59","00:00"]
输出: 1

备注:

列表中时间数在 2~20000 之间。
每个时间取值在 00:00~23:59 之间。

*/

/*
思路：
	将时间换算成分钟，然后直接进行计算。
	可以判断，两个时间的最大差值为十二小时，因为时间是双向的。

解：
	先将时间转换成数字，使用正则匹配，并将匹配结果转换成int值
	然后进行换算，换算结果=小时*60+分钟

笔记：
	考虑情况没考虑全，忘记计算负数问题
*/

func Test_539(t *testing.T) {
	arr := []string{"00:00", "23:59", "00:00"}
	//arr := []string{"01:01","02:01"}
	fmt.Println(findMinDifference(arr))
}

func findMinDifference(timePoints []string) int {
	num := 10000

	var timeArr []int

	reg := regexp.MustCompile(`[\d]+`)

	for _, v := range timePoints {
		numArr := reg.FindAllString(v, -1)

		num1, _ := strconv.Atoi(numArr[0])
		num2, _ := strconv.Atoi(numArr[1])

		value := num1*60 + num2
		timeArr = append(timeArr, value)
	}

	for i := 0; i < len(timeArr); i++ {
		for j := i + 1; j < len(timeArr); j++ {
			value := timeArr[i] - timeArr[j]
			if value < 0 {
				value = 0 - value
			}
			if value > 720 {
				value = 1440 - value
			}
			if value < num {
				num = value
			}
		}
	}

	return num
}
