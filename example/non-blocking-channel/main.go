package main

import (
	"fmt"
	"time"
)

func main()  {
	messages := make(chan string)
	signals := make(chan bool)

	go func() {
		// time.Sleep(time.Second*1)
		signals <- true
		messages <- "12"

	}()



	select {
	case msg := <-messages:
		fmt.Println("received message1", msg)
	case <-time.After(time.Second*2):
		fmt.Println("timeout")
	//default:
	//	fmt.Println("no message received")
	}

	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-messages:
		fmt.Println("received message2", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

}













/*
常规的通过通道发送和接受数据是阻塞的。然而，我们可以使用带一个default子句的select来实现非阻塞的发送、接受，甚至是非阻塞的多路select


*/
