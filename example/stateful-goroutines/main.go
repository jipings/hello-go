package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key int
	resp chan int
}

type writeOp struct {
	key int
	val int
	resp chan bool
}

func main()  {
	var rOps int64
	var wOps int64

	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
				// fmt.Println(state[read.key], read.key)
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	for r := 0; r < 1000; r++ {
		go func() {
			read := &readOp{
				key: rand.Intn(5),
				resp: make(chan int)}
			reads <- read
			<-read.resp
			atomic.AddInt64(&rOps, 1)
		}()
	}

	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOp{
					key: rand.Intn(5),
					val: rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddInt64(&wOps, 1)
			}
		}()
	}

	time.Sleep(time.Second)

	rOpsFinal := atomic.LoadInt64(&rOps)
	wOpsFinal := atomic.LoadInt64(&wOps)

	fmt.Println("ops:",rOpsFinal, wOpsFinal)
}

/*
go 状态协程 通过通讯来共享内存
我们用互斥锁进行了明确的锁定来让共享的state 跨多个 Go 协程同步访问。
另一个选择是使用内置的 Go协程和通道的的同步特性来达到同样的效果。
这个基于通道的方法和 Go 通过通信以及 每个 Go 协程间通过通讯来共享内存，
确保每块数据有单独的 Go 协程所有的思路是一致的。
*/