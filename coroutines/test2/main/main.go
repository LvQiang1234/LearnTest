package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func say(s string) {
	fmt.Println(s)
	wg.Done()
}

func main() {
	wg.Add(1)
	go say("Hello World")
	wg.Wait()
}


