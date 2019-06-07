package main

import (
	"fmt"
	"time"
)

func main()  {
	timer1 := time.NewTicker(time.Second * 2)

	<-timer1.C
	fmt.Println("Timer 1 expired")

	timer2 := time.NewTicker(time.Second)

	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	timer2.Stop()
	//if stop2 {
	//	fmt.Println("Timer 2 stopped")
	//}
}

/*

我们常常需要在后面一个时刻运行 Go 代码，或者在某段时间间隔内重复运行。Go 的内置 定时器 和 打点器 特性让这些很容易实现
*/