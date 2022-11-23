package main

import "fmt"

func main() {
	jobs := make(chan int)
	done := make(chan bool)

	go func() {
		for {
			j, more := <- jobs
			if more {
				fmt.Println("작업 수행 : ", j)
			} else {
				fmt.Println("더 이상 작업이 없음")
				done <- true
				return
			}
		}
	} ()

	for j := 0; j <= 3; j++ {
		jobs <- j
		fmt.Println("작업 요청 : ", j)
	}
	close(jobs)
	fmt.Println("모든 작업 마침")
	<- done
}