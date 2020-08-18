package DynamicProgramming

import (
	"fmt"
	"testing"
)

func Fibonacci(n int) int {
	if n <= 0 {
		return n
	}

	ints := make([]int, n+1)
	for i := 0; i <= n; i++ {
		ints[i] = -1
	}
	return fib(n, ints)
}

func fib(n int, ints []int) int {
	if ints[n] != -1 {
		return ints[n]
	}

	if n <= 2 {
		ints[n] = 1
	} else {
		ints[n] = fib(n-1, ints) + fib(n-2, ints)
	}
	return ints[n]
}

func Test_Fib(t *testing.T) {
	fmt.Println(Fibonacci(6))
	fmt.Println(fib2(6))
	fmt.Println(fib3(6))
}

func fib2(n int) int {
	if n <= 0 {
		return n
	}

	ints := make([]int, n+1)
	ints[0] = 0
	ints[1] = 1
	for i := 2; i <= n; i++ {
		ints[i] = ints[i-1] + ints[i-2]
	}
	return ints[n]
}

func fib3(n int) int {
	if n <= 1 {
		return n
	}

	arr2 := 0
	arr1 := 1
	arr0 := 1
	for i := 2; i <= n; i++ {
		arr0 = arr2 + arr1
		arr2 = arr1
		arr1 = arr0
	}
	return arr0
}
