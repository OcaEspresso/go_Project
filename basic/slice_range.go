package main

import (
	"fmt"
)

func main() {
	arr := make([]int, 3, 5)
	arr[0], arr[1], arr[2] = 2, 7, 9
	for i, ele := range arr {
		fmt.Println(i, ele)
	}
}
