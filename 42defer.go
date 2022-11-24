package main

import (
	"fmt"
	"os"
)

func main() {
	f := createFile("/Users/jiyeonpark/GoProject/gobyexample/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("Creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("Writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("Closeing")
	f.Close()
}