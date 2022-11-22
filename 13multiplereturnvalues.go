package main

import (
	"errors"
	"fmt"
)

func vals() (int, int) {
	return 3, 7
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("분모로 0을 사용할 수 없습니다.")
	}

	return a/b, nil
}

func main() {
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)

	a, c = vals()
	fmt.Println(a, c)

	//res, err := divide(4, 0)
	res, err := divide(4, 2)
	if err != nil {
		fmt.Println("ERROR: ", err.Error())
	} else {
		fmt.Println("4/2=", res)
	}
}