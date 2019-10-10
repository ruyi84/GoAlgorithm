package main

import (
	"fmt"
)

//一边hashmap
func twoSum1(nums []int, target int) []int {
	nes := make([]int, 2)
	m1 := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		num := target - nums[i]
		if _, flag := m1[num]; !flag {
			m1[nums[i]] = i
		} else {
			nes[0], nes[1] = m1[num], i
		}
	}
	return nes
}

func main() {
	cc := []int{2, 7, 11, 15}
	fmt.Println(twoSum1(cc, 9))
}

/*
第一题也是个小神仙题，第一次做的时候是按照T1中的双重for循环，当时感觉题目很像在校写的一样。
后来，看到后面的一个题目，别人用了一个map存取，两次for循环，我觉得这个很棒，感觉会快很多。
然后我就尝试重写，但是发现一个问题，他是取的值，我是要取数组下标。
go里，我并不知道怎么map的key和value互取，于是我就想到存两个map，k和v互反，这样就解决了。
但是接着又出现一个问题，当出现3+3=6这种情况，他会取重，比如[1,1]
我反思，想起来一个方式，既然确定了只会存在一对，那么我完全可以一次循环解决这个问题：
第一步，计算出下标对应的目标值
第二步，判断目标值是否存在，不存在就把当前的值存进map，若存在及取出目标值
第三步，返回结果

感觉这个方式很棒，空间复杂度不是很懂，但是利用leetCode里：
解决方法	执行耗时	内存消耗
T1			8ms			3
T1_2		80ms		3.8
时间是快了，但是关于怎么计算的内存消耗还是存在不明白的地方
*/
