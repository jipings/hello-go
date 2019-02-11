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

## 4.3 Map
```go
_ = &ages["bob"] // compile error: cannot take address of map element
```
* 无法获取map 元素的地址的一个原因是 map 的增长可能会导致已有元素被重新散列到新的存储位置，这样就可能使获取的地址无效
* 散列的存储是无序的，如果要按顺序遍历key/value对，我们必须显式地对key进行排序，可以使用sort包的Strings函数对字符串slice进行排序。
* 要判断两个map是否包含相同的key和value，我们必须通过一个循环实现
```go
func equal(x, y map[string]int) bool {
    if(len(x) != len(y)) {
        return false
    }
    for k, xv := range x {
        if yv, ok := y[k]; !ok || yv != xv {
            return false
        }
    }
    return true
}
```


## 5 27 *
## 6 12
## 7 36 *
## 8 30 *
## 9 20
## 10 8
## 11 20
## 12 20
## 13 10