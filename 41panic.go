package main

import "os"

func main() {
	panic("A problem!!!")

	_, err := os.Create("/Users/jiyeonpark/GoProject222/file01")
	if err != nil {
		panic(err)
	}
}