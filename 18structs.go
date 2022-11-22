package main

import "fmt"

type person struct {
	name string
	age int
}

func main() {
	yundream := person{name:"yundream", age:25}
	jackson := person{name:"jackson", age:31}
	fmt.Println(yundream)
	fmt.Println(jackson)

	sp := &jackson
	sp.name = "jackson json"
	sp.age = 48
	fmt.Println(sp.age)
	fmt.Println(jackson)
}