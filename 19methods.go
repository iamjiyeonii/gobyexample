package main

import "fmt"

type rect struct {
	width, height int
}

func (self *rect) area() int {
	return self.width * self.height
}

func (self rect) perim() int {
	return 2 * self.width + 2 * self.height
}

type ByteSlice []byte

func (self ByteSlice) Append(data []byte) []byte {
	return append(self, data...)
}

func (self *ByteSlice) MAppend(data []byte) {
	*self = append(*self, data...)
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("Area: ", r.area())
	fmt.Println("perim: ", r.perim())

	rp := &r
	fmt.Println("Area: ", rp.area())
	fmt.Println("perim: ", rp.perim())

	var b ByteSlice
	c := b.Append([]byte("Hello World"))
	fmt.Println("value: ", string(c))

	var f ByteSlice
	f.MAppend([]byte("Hello World"))
	fmt.Println("pointer: ", string(f))
}