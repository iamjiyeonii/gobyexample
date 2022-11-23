package main

import (
	"fmt"
	"time"
)

type info struct {
	name string
}

func main() {
	channel := make(chan info, 1)

	go func() {
		time.Sleep(time.Second * 1)
		info1 := info{"first goroutine"}
		channel <- info1
	}()

	go func() {
		time.Sleep(time.Second * 2)
		info2 := info{"second goroutine"}
		channel <- info2
	}()

	for i := 0; i < 2; i++ {
		select {
		case info := <-channel:
			fmt.Println("Received: ", info)
		}
	}
}
