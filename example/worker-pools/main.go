package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int)  {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main()  {

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j<= 99; j++ {
		jobs <- j
	}

	for a := 1; a <= 9; a++ {
		<-results
	}
}

/*
使用协程和通道实现的一个工作池
这些执行者将从 jobs 通道接收任务，并且通过 results 发送对应的结果。我们将让每个任务间隔 1s 来模仿一个耗时的任务
*/