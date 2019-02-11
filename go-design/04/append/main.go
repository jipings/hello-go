// 为 [] int 数组 slice 定义的方法 appendInt
package main

import (
	"fmt"
)

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// slice 仍有增长空间，拓展slice内容
		z = x[:zlen]
	} else {
		// slice 已无空间，为它分配一个新的底层数组
		// 为了达到分摊线性复杂性，容量扩展一倍
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2*len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z,x)
	}
	z[len(x)] = y
	return z;
}

func main() {
	var x,y []int
	for i := 0; i< 10; i++ {
		y = appendInt(x,i)
		fmt.Printf("%d cap=%d\t%v\n",i,cap(y),y)
		x = y
	}
}