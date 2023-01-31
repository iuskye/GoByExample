package main

import (
	"fmt"
	"time"
)

func main() {
	// 在这个例子中，我们将从两个通道中选择。
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// 我们使用 select 关键字来同时等待这两个值，并打印各自接收到的值。
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}

// 跟预期的一样，我们首先接收到值“one”，然后是“two”。
