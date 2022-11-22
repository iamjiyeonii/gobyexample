package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("goroutine")

	go func(msg string) {
		time.Sleep(time.Second * 1)	
		fmt.Println(msg)
	}("going")

	//time.Sleep(time.Second * 2)
	fmt.Println("done")
}