package main

import (
	"fmt"
	"runtime"
	"strconv"
)

func main() {
	runtime.GOMAXPROCS(0)
	for i := 0; i < 5; i++ {
		go say("Hello World: " + strconv.Itoa(i))
	}
	for {

	}
}

func say(s string) {
	fmt.Println(s)
}
