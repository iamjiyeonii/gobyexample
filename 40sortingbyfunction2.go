package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Id   int
}

type Persons []Person

func (self Persons) Len() int {
	return len(self)
}

func (self Persons) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

func (self Persons) Less(i, j int) bool {
	return self[i].Id < self[j].Id
}

func (self Persons) Sort() {
	sort.Sort(Persons(self))
	fmt.Println(self)
}

func main() {
	data := Persons{
		{"ab", 33},
		{"hong", 21},
		{"kim", 38},
		{"Tei", 29},
		{"Tom", 28},
	}
	data.Sort()
}