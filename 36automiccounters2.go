package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
		runtime.Gosched()
	}
}

func main() {
	go say("world")
	say("hello")
}
