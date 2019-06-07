package main

import (
	"fmt"
	"strings"
)

func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func Include(vs []string, t string) bool {
	return Index(vs, t) >0
}

func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return  true
		}
	}
	return false
}

func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main()  {
	var strs = []string{"peach", "apple", "pear", "plum"}
	fmt.Println(Index(strs, "pear"))
	fmt.Println(Include(strs, "grape"))
	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))
	fmt.Println(Map(strs, strings.ToUpper))

}
/*

我们经常需要程序在数据集上执行操作，比如选择满足给定条件的所有项，或者将所有的项通过一个自定义函数映射到一个新的集合上

在某些语言中，会习惯使用泛型。Go 不支持泛型，在 Go 中，当你的程序或者数据类型需要时，通常是通过组合的方式来提供操作函数。

*/
