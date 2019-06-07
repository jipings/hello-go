package main
// 当使用通道作为函数的参数时，你可以指定这个通道是不是
// 只用来发送或者接受值。这个特性提升了程序的类型安全
import (
	"fmt"
)

func ping(pings chan<- string, msg string)  {
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main()  {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}