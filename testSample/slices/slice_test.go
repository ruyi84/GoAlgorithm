package slices

import (
	"fmt"
	"testing"
)

/*
可以理解slice（切片）是数组的视图
即，slice本身并不存储值，是对arr的一个view
*/
func Test_Slice(t *testing.T) {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	s1 := arr[2:5]
	fmt.Println(s1)
	s2 := s1[3:7]
	fmt.Println(s2)
}
