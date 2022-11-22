package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iprt *int) {
	*iprt = 0
}

func main() {
	i := 100
	fmt.Println("초기값 : ", i)

	zeroval(i)
	fmt.Println("zeroval : ", i)

	zeroptr(&i)
	fmt.Println("zeroptr : ", i)

	fmt.Println("포인터 : ", &i)
}