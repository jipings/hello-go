package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {
	var ops uint64 = 0

	for i := 0; i< 50; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1)

				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
/*
go 中最主要的状态管理方式是通过通道间的沟通完成的，但是还是有一些其他的方法来管理状态的。
这里我们将看看如何使用 sync/atomic包在多个 Go 协程中进行 原子计数 。
*/
