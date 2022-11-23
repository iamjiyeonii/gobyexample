package main

import "fmt"

func main() {
	message := make(chan string)
	signals := make(chan bool)

	select {
	case msg := <-message:
		fmt.Println("Received message ", msg)
	default:
		fmt.Println("no message received")
	}

	msg := "hi"

	select {
	case message <- msg:
		fmt.Println("Send message ", msg)
	default:
		fmt.Println("no message sent")
	}

	select {
	case msg := <-message:
		fmt.Println("Received message ", msg)
	case sig := <-signals:
		fmt.Println("Received signal ", sig)
	default:
		fmt.Println("no activity")
	}
}
