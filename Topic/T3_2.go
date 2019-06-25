package main

import "fmt"

func lengthOfLongestSubstring2(s string) int {
	sli := make([]string, 0)
	now, max := 0, 0 //当前存储字符串长度，历史最大长度
	boo := false     //用来记录当前字是否存在
	var str string   //存储当前字符

	for _, v1 := range s {
		str = string(v1)
		for k2, v2 := range sli {
			if str == string(v2) {
				sli = sli[k2+1:]
				sli = append(sli, str)
				now = len(sli)
				boo = true
				break
			}
		}
		if boo == false {
			sli = append(sli, str)
			now = len(sli)
		}
		if now >= max {
			max = now
		}
		boo = false
	}
	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring2("pwwkew"))
}

/*
	1、遍历字符串，依次取出字母存储计算长度。
	2、定义两个数字，一个是当前存储字符串长度，另一个是历史最大长度
	3、将取出字母在切片中检查是否存在，如果不存在就存进去，当前长度和最大长度均加一
	4、如果存在，则找寻位置，并删除已存在位置及其以前的字符串，返回当前长度
	5、每次长度比较，如果当前长度大于最大长度，则最大长度复制当前长度值

	反思，可以使用滑动窗口
	1、使用一个

*/
