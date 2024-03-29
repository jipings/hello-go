
## 语言结构
go的语言的基础组成有以下几个部分
* 包声明
* 引入包
* 函数
* 变量
* 语句 & 表达式
* 注释

## go 派生类型
* 指针类型
* 数组类型
* 结构化类型
* Channel类型
* 函数类型
* 切片类型
* 接口类型（interface）
* Map类型

## iota
iota，特殊常量，可以认为是一个可以被编译器修改的常量。
iota 在 const 关键字时将被重置为 0 (const内部的第一行之前)，const 中每新增一行常量声明使iota计数一次（iota可理解为const语句块中的行索引）。

* 第一个 iota 等于 0，每当 iota 在新的一行被使用时，它的值都会自动加1；

## 无类型常量
许多常量并不从属某一具体类型。编译器将这些从属类型待定的常量表示成某些值，这些
值比其他基本类型的数字精度更高，且算术精度高于原生的机器精度。可认为其精度达到了256位，从属类型待定的常量共6种，分别是无类型布尔、无类型整数、无类型文字符号，无类型浮点数，无类型复数，无类型字符串

## 特殊运算符

* & 返回变量存储的内存地址
* * 指针变量 *a; 是一个指针变量

## 语句

* select 语句
1) 每个case 都必须是一个通信
2) 所有channel表达式都会被求值
3) 所有被发送的表达式都会被求值
4) 如果任意某个通信可以进行，它就被执行；其他忽略
5) 如果有多个case都可以运行，Select会随机公平地选出一个执行，其他不会执行，否则：如果有default子句，则执行该语句，如果没有default子句，select将阻塞，直到某个通信可以运行；go不会重新对channel或表达式进行求值

* goto 语句

goto 语句可以无条件地转移到过程中制定的行

## 函数
Go 语言最少有个`main()`函数
```go
    func function_name ([ parameter list ]) [ return_types ] {
        // 函数体
    }
```
函数的定义解析
1) func: 函数由func开始声明
2) function_name: 函数名称，函数名和参数列表一起构成了函数签名
3) parameter list： 参数列表，参数就像一个占位符，当函数被调用时，你可以将值传递给参数，这个值被称为实际参数。参数列表指定的是参数类型、顺序、及参数个数。参数是可选的，也就是说函数也可以不包含参数。
4) return_types: 返回类型，函数返回一列直。return_types 是该列值得数据类型。有些功能不需要返回值，这种情况下 return_types 不是必须的.5) 函数体： 函数定义的代码集合
6) 函数可以返回多个值

## 数组

数组是具有相同唯一类型的一组已编号长度固定的数据项序列，这种类型可以是任意的原始类型，如整形，字符串或者自定义类型。

* 数组的声明
```js
    var variable_name [SIZE] variable_type
```

## go 指针
* 变量是一种使用方便的占位符，用于引用计算机内存地址，go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。

* 指针，一个指针变量指向一个值得内存地址。使用前同样需要声明指针
```go
var var_name *var-type
```
var-type 为指针类型，var_name 为指针变量， * 符号用于指定变量是作为一个指针。
```go
var ip *int // 指向整型的指针
var fp *float32 // 指向浮点型
```
* 如何使用指针
1) 定义指针变量
2) 为指针变量赋值
3) 访问指针变量中指向地址的值

* 空指针
当一个指针被定义后没有分配到任何变量时，它的值为 `nil`
nil在概念上和其他语言的 null,none,nil一样

## go 语言结构体
go 语言中数组可以存储同一类型的数据，但在结构体中我们可以为不同项定义不同的数据类型。
结构体是由一系列具有相同类型或不同类型的数据构成的数据集合

* 定义结构体
定义结构体需要使用type 和 struct 语句。 struct 语句定义一个新的数据类型，结构体有中有一个或多个成员。type语句设定了结构体的名称。
```go
type struct_variable_type struct {
    member definition
    member definition
    ...
}
```

## go 语言切片( Slice )

go语言的切片是数组的抽象
go 数组的长度不可改变，在特定场景中这样的集合就不太适用，go 中提供了一种灵活，功能强悍的内置类型切片（”动态数组“），与数组相比切片的长度
是不固定的，可以追加元素，在追加时可能使切片的容量增大

* 定义切片
```go
var identifier []type
```
切片不需要说明长度
或适用make() 函数来创建切片：
```go
var slice1 []type = make([]type, len)
// 简写
slice2 := make([]type, len)
// 也可以指定容量，其中capacity为可选参数
make([]type, len ,capacity);

// 切片初始化
s := [] int {1,2,3}
s := arr[startIndex: endIndex]

```
* len() 和 cap() 函数
切片是可索引的，并且可以由 len() 方法获取长度
切片提供了计算容量的方法cap() 可以测量切片最长可以达到多少

* 空(nil)切片

* 切片截取

可以通过设置下限及上限来设置截取切片 [lower-bound:upper-bound]

* append() 和 copy() 函数
如果想增加切片容量，我们必须创建一个新的更大的切片并把原切片的内容都拷贝过来

## 语言范围(Range)
go语言中 range 关键字用于 for 循环中迭代数组（array）、切片(slice)、通道(channel) 或者集合（map）的元素。在数组和切片中它返回元素的索引和索引对应的值，在集合返回 key-value 的key值

## Go 语言集合 Map
Map 是一种无序的键值对的集合。Map 最重要的一点是通过key来快速检索数据，key类似于索引，指向数据的值
Map 是一种集合，所以我们可以向迭代数组和切片那样迭代它。不过，Map是无序的，我们无法决定它的返回顺序，这是因为`Map是使用hash`表来实现的。

定义Map
可以使用内建函数 make 也可以使用 map 关键字来定义 Map

```go
// 声明变量，默认 map 是 nil
var map_variable map[key_data_type] value_data_type
// 使用 make 函数
 map_variable := make(map[key_data_type]value_data_type)
```
如果初始化 map， 那么就会创建一个 nil map。nil map 不能用来存放键值对

## 递归
 递归就是在运行过程中调用自己
* 递归阶乘
```go
package main

import "fmt"

func main() {
	var i int = 15
    fmt.Printf("%d 的阶乘是 %d\n", i, Factorial(uint64(i)))
}

func Factorial(n uint64)(result uint64) {
	if(n>0) {
		result = n * Factorial(n - 1);
		return result
	}
	return 1
}
```
* 打印斐波那契数列
```go
package main

import "fmt"

func main() {
	var i int;
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}
```
## go 语言接口
go 语言提供了另外一种数据类型及接口，它把所有的具有共性的方法定义在一起，任何其他类型只要实现了这些方法就是实现了这个接口

```go
// 定义接口
type interface_name interface {
    method_name1 [return_type]
    method_name2 [return_type]
    ...
    method_namen [return_type]
}
// 定义结构体
type struct_name struct {
    // variables
}
// 实现接口方法
func (struct_name_variable struct_name) method_name1() [return_type]{
    // 方法实现
}
```
实例
```go
package main

import "fmt"

type Phone interface {
	call()
}
type NokiaPhone struct {

}
func (nokiaPhone NokiaPhone) call() {
    fmt.Println("I am Nokia, I can call you!")
}
type Iphone struct {

}
func (iPhone Iphone) call() {
	fmt.Println("I am iPhone, I can call you!")
}
func main()  {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()
	phone = new(Iphone)
	phone.call()
}
```
在上面的例子中，我们定义了一个接口 Phone，折扣里面有一个call()。然后我们在main函数里面定义了一个Phone类型变量，并分别为之赋值为NokiaPhone和IPhone

## 错误处理
go 语言通过内置的错误接口提供了非常简单的错误处理机制
error 类型是一个接口类型，这是它的定义
```go
type error interface {
    Error() string
}
```
我们可以在编码中通过实现 error 接口类型来生成错误信息。
函数通常在最后的返回值中返回错误信息。使用errors.New 可返回一个错误信息：
```go
func Sqrt(f float64) (float64,error) {
    if f < 0 {
        return 0, errors.New("math: square root of negative number")
    }
} 
```