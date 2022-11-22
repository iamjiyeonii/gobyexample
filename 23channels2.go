package main

import (
	"fmt"
	"time"
)

func ff(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	messages := make(chan string)

	ff("direct")

	go ff("goroutine")

	go func(msg string) {
		time.Sleep(time.Second * 1)	
		fmt.Println(msg)
	}("going")

	go func() {
		time.Sleep(time.Second * 2)
		messages <- "done"
	} ()

	msg := <- messages
	fmt.Println(msg)
}