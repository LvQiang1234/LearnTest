package main

import (
	"fmt"
)

func main() {
	go say("Hello World")
	///time.Sleep(time.Second * 1)
}

func say(s string) {
	fmt.Println(s)
}