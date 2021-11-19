package main

import "fmt"

//select 的功能和 select, poll, epoll 相似， 就是监听 IO 操作，当 IO 操作发生时，触发相应的动作。

func main() {
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)

	ch2 <- 2

	ch3 <- 3

	//
	select {
	case <-ch2:
		fmt.Println("ch2 pop one element")
		break
		fmt.Println("execute break")
	case <-ch3:
		fmt.Println("ch3 pop one element")
		break
		fmt.Println("execute break")
	}
}
