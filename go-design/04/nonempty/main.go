// slice 的就地修改算法
package main
import "fmt"
// nonempty 返回一个新的slice，slice中的元素都是非空字符串
// 在函数调用过程中，底层数组的元素发生了改变
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s;
			i++
		}
	}
	return strings[:i]
}
/* 这里有一点是输入的slice和输出的slice拥有相同的底层数组，这样就避免了内部
重新分配一个数组。当然这种情况下，底层数组的元素只是部分被修改
*/
// 使用append函数
func nonempty2(strings []string) []string {
	out := strings[:0] // 引用原始slice的新的零长度
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out;
}

func remove(slice []int, i int) []int {
    copy(slice[i:], slice[i+1:])
    return slice[:len(slice)-1]
}
// 调换位置，但是顺序会打乱
func remove2(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func main() {
	data := []string{"1","4","3"}
	fmt.Printf("%q\n", nonempty(data))
	fmt.Printf("%q\n",data)

	var a = []string{"1","","2","3"}
	a = nonempty(a)
	fmt.Println(a)
	
	s := []int{5,6,7,8,9}
	// fmt.Println(s[2:], s[3:])
	
	fmt.Println(remove(s,2))
}
