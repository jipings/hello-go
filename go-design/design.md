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
## 5.2 递归
大部分编程语言使用固定大小的函数调用栈，常见的大小从64kb 到 2mb不等。固定的大小栈会限制递归的深度，当你用递归处理大量数据时，需要避免栈溢出；除此之外还会导致安全性问题。与此相反，go 使用可变栈，栈的大小按需增加(初始时很小)。这使得我们使用递归时不必考虑溢出和安全问题

## 5.8 defer 函数

当defer语句被执行时，跟在defer后面的函数会被延迟执行。直到包含该defer语句的函数执行完毕时，跟在defer后面的函数会被延迟执行。直到包含该defer语句的函数执行完毕时，defer后的函数才会被执行，不论包含defer语句的函数是通过return 正常结束，还是由于panic导致的异常结束。你可以在一个函数中执行多条defer语句，它们执行的顺序和声明顺序相反

## 6.1 方法声明
封装和组合

## 6.2 指针接收者的方法
由于主调用函数会复制每一个实参变量，如果函数需要更新一个变量，或者如果一个实参数太大而我们希望避免复制整个实参，
因此我们必须使用指针来传递变量地址。这同样适用于更新接收者

## 6.5 位向量
位向量是理想的数据结构，使用一个无符号整型值的slice，每一位代表集合中的一个元素。如果设置第 i 位的元素，则认为集合包含 i。
## 6.6 封装
如果变量或者方法不能通过对象访问到的，这称作封装的变量或者方法。封装（有时候称作数据隐藏）是面向对象编程中重要的一方面

## 8.4 通道
如果说 goroutine 是 Go 程序并发执行体，通道就是它们之间的连接。通道是可以让一个 goroutine 发送特定值到另一个 goroutine 的通信机制。每一个通道是一个具体类型的导管，叫作通道的元素类型。一个有 int 类型元素的通道写为 chan int
```go
ch := make(chan int)
```
使用简单的make调用创建的通道叫 无缓冲（unbuffered）通道，但make还可以接受第二个可选参数，一个表示通道容量的整数。如果容量是0，make创建一个无缓冲通道：
```go
ch = make(chan int) // 无缓冲通道
ch = make(chan int ,0) // 无缓冲通道
ch = make(chan int , 3) // 容量为 3 的缓冲通道
```
无缓冲通道上的发送操作将会阻塞, 无缓存通道也称为 同步通道。

* 通道可以用来连接 goroutine，这样一个的输出是另一个的输入。这个叫管道(pipeline)

## 9 20
## 10 8
## 11 20
## 12 20
## 13 10