package main

import (
	"fmt"
	"time"
)

func main() {

	jobs := make(chan int,5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <- jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("reveived all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		time.Sleep(time.Second)
		jobs <- j
		fmt.Println("sent jobs", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	<-done
}

/*
关闭一个通道意味着不能再向这个通道发送值了。这个特性可以用来给这个通道的接受方法传达工作已经完成的信息
*/