package main

import (
	"fmt"
	"strconv"
	"sync"
)

var wg = sync.WaitGroup{}

func say(s string) {
	fmt.Println(s)
	wg.Done()
}

func main() {
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go say("Hello World: " + strconv.Itoa(i))
	}
	wg.Wait()
}
