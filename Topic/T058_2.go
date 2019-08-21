package main

import "fmt"

/*
感觉之前的解决方式有待提高，感觉第一可以通过正则进行计算，直接匹配最后一个单词内容，使用len（）取出即可
后来看了一下题解
打开的第一时间看到从后向前遍历，发现也是一个不错的方式
*/
func lengthOfLastWord2(s string) int {
	Len1 := 0
	Len := 0
	for _, v := range s {

		if v == 32 {
			Len = 0
			continue
		}
		Len1++
		if Len == 0 {
			Len1 = 1
		}
		Len = Len1
	}
	return Len1
}

func main() {
	str := "Hello "

	fmt.Println(lengthOfLastWord(str))
}
