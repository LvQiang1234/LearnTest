package main

import (
	"fmt"
	"strconv"
)

func say(s string, c chan string) {
	c <- s
}

func main() {
	var result = make(chan string)
	for i := 0; i < 5; i++ {
		go say("Hello World: " + strconv.Itoa(i), result)
	}
	for s:= range result {
		fmt.Println(s)
	}
}
