package pointer

import (
	"fmt"
	"testing"
)

/*
测试golang中对象作为参数的参数引用
*/

type cache struct {
	name string
	age  int
}

func f(cache *cache) {
	cache.age = cache.age + 1
}

func Test_parameter(t *testing.T) {
	var cache *cache
	cache.age = 11
	f(cache)
	fmt.Println(cache.age)
}
