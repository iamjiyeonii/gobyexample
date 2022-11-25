package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check2(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("/Users/jiyeonpark/GoProject/gobyexample/dat1.txt", d1, 0644)
	check2(err)

	f, err := os.Create("/Users/jiyeonpark/GoProject/gobyexample/dat2.txt")
	check2(err)

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check2(err)
	fmt.Printf("Wrote %d byte\n", n2)

	n3, err := f.WriteString("writes\n")
	fmt.Printf("Wrote %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("Wrote %d byte\n", n4)

	w.Flush()
}