package main

import "fmt"

/*
给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。

示例 1:

输入: s = "anagram", t = "nagaram"
输出: true
示例 2:

输入: s = "rat", t = "car"
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-anagram
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
看到题目的第一时间想到了排序直接做对比，但是秉着第一想法不是最佳实现的想法，而且相对没啥意思就继续向其他方法

之后想到，是否可以存成数组，直接对比数组内容，于是开始code
code中遇到一个问题，因为我开始定义数组的问题，出现a，b返回true，之后排查中突然惊醒一个问题
	如果单纯的分别取出放入到数组，然后两次遍历对比a数组的部分是否出现在b数组中，会出现的一个问题：
		万一，a中出现两次，b中仅出现一次，那么仍旧会返回true，虽然有方法解决，但是复杂了一些，于是及时止损。

紧接着我思考了一下两个都包含，但是次数出现不一样的解决方法，区别在差值，bingo
map恰好解决，把s字符串计算每个字符出现的次数，存入map value值++，
然后再把t字符串中取出vaule- - ，然后判断map中是否value都为0，完美解决
两次for循环。
*/

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	strmap := make(map[interface{}]int, len(s)+1)

	for i := 0; i < len(s); i++ {
		strmap[s[i]]++
		strmap[t[i]]--
	}

	for _, v := range strmap {
		if v != 0 {
			return false
		}
	}
	return true
}

func main() {
	s := "a"
	t := "b"
	fmt.Println(isAnagram(s, t))
}
