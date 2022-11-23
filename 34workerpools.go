package main

import (
	"fmt"
	"time"
)

func worker3(id int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Println("워커 ", id, "작업처리 중 : ", j)
		time.Sleep(time.Millisecond * 100)
		result <- j * 2
	}
}

func main() {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker3(w, jobs, results)
	}

	for j := 1; j <= 9; j++ {
		jobs <- j
	}

	close(jobs)
	for a := 1; a <=9; a++ {
		<- results
	}
}