package main

import (
	"fmt"
	"strings"
)

func isValid(s string) bool  {
	mapping := map[string]string{")":"(","]":"[","}":"{"}
	var stack []string

	for _, c:= range s {
		str := string(c)
		fmt.Println(str)
		if ms, ok := mapping[str]; !ok {
			fmt.Println(ms,ok)
			stack = append(stack, str)
		} else {
			fmt.Println(ms,ok)
			if len(stack) == 0 {
				return  false
			} else {
				if ss := stack[len(stack) - 1]; ss != ms {
					return false
				}
				stack = stack[:len(stack) - 1]
			}
		}
	}
	return len(stack) == 0
}

/**
核心思想： 利用栈的性质，先进后出，遇到左括号则压入栈，遇到右括号则与栈顶元素匹配，若匹配成功则将栈顶元素弹出，反之返回false。

匹配方法： 利用map建立哈希表，实现括号一一对应关系。
链接：https://leetcode-cn.com/problems/two-sum/solution/you-xiao-de-gua-hao-by-gpe3dbjds1/
*/


func isValid2(s string) bool {
	for s != "" {
		num := len(s)
		s = strings.Replace(s, "{}", "", -1)
		s = strings.Replace(s, "[]", "", -1)
		s = strings.Replace(s, "()", "", -1)
		if num == len(s) {
			return false
		}
	}
	return true
}

func main()  {
	str := "(){}[()]"
	fmt.Println(isValid(str))
}