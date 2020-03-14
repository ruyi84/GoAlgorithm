package pointer

import (
	"fmt"
	"testing"
)

func Test_Assignment(t *testing.T) {
	var a int = 2
	var b int = 2
	//定义一个类型为int型指针的指针变量pa，并且将地址指向a的地址
	var pa *int = &a
	fmt.Println(pa)
	pa = &b
	fmt.Println(a)
}

type Person struct {
	Age int
}

//测试关于引用类型
//go中存在引用类型包括slice、map、channel。引用类型存储的是地址
func Test_Pointer(t *testing.T) {
	var arr [10]int
	arr[0] = 11
	arrS := make([]int, 10)
	arrS[0] = 11
	Bob := Person{Age: 10}

	arr2 := arr
	arrS2 := arrS

	fmt.Println("old", arr, arrS, Bob)
	Changer(arr)
	Changer(arrS)
	arr2[1] = 2
	arrS2[1] = 2
	Changer(Bob)
	fmt.Println("After changer", arr, arrS, Bob)

}

//为数组或切片的0下标处重新赋值
func Changer(x interface{}) {
	switch x := x.(type) {
	case [10]int:
		x[0] = 0
	case []int:
		x[0] = 0
	case Person:
		x.Age = 0
	}

}
