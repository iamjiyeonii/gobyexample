package main

import (
	"fmt"
	"sort"
)

func sum(nums ...int) {
	fmt.Print(nums, "\n")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func min(nums ...int) {
	fmt.Print(nums, "\n")
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	fmt.Println(nums[0])
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4, 5}
	sum(nums...)

	//nums = []int{5, 1, 7, 8, 3, 9}
	nums = []int{25, 21, 57, 18, 23, 9}
	min(nums...)
}
