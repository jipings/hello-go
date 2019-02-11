## 3.3 复数
* go 具备两种大小的复数 `complex64` 和 `complex128`, 二者分别由 float32
和 float64构成

## 4.2.2 Slice 内存技巧
```go
// slice 尾部追加值
stack = append(stack, v)
// 栈的顶部是最后一个元素
top := stack[len(stack)-1]
// 通过弹出最后一个元素缩减栈
stack = stack[:len(stack)-1] // pop
// 删除中间某个元素
func remove(slice []int, i int) []int {
    copy(slice[i:], slice[i+1:])
    return slice[:len(slice)-1]
}
```