package main

import (
	"fmt"
	"sort"
)

type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap (i, j int)  {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i,j int) bool {
	return len(s[i]) < len(s[j])
}

func main()  {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}


// 有时候我们想使用和集合的自然排序不同的方法对集合进行排序。
// 例如，我们想按照字母的长度而不是首字母顺序对字符串排序。这里是一个 Go 自定义排序的例子。
/*
我们在类型中实现了 sort.Interface 的 Len，Less和 Swap 方法，这样我们就可以使用 sort 包的通用Sort 方法了，
Len 和 Swap 通常在各个类型中都差不多，Less 将控制实际的自定义排序逻辑。
在我们的例子中，我们想按字符串长度增加的顺序来排序，所以这里使用了 len(s[i]) 和 len(s[j])。
*/