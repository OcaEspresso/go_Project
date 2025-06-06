package main

/**
slice 切片在go中很重要
**/

import (
	"fmt"
)

func main() {

	// 声明数组 数组容量是固定的所以后续动态添加是不行的，所以发明了slice
	// 使用 make运算符创建slice
	//var arr [5]int
	//arr = [5]int{2, 9, 7, 3, 5}

	arr := make([]int, 3, 5)
	arr[0], arr[1], arr[2] = 2, 9, 7
	brr := append(arr, 8)
	brr = append(brr, 8)
	arr[0] = 4
	fmt.Println(brr[0]) //会输出4

	brr = append(brr, 8) // 超过5了，发生扩容，会申请新的内存空间过去，这里开始不共享内存空间了
	arr[1] = 5           // arr更改，brr不受影响了
	fmt.Println(brr[1])

	fmt.Println(len(brr), cap(brr))
	fmt.Println(len(arr), cap(arr))
}

/**

	type slice struct {

	array unsafe.Pointer
	len int  ---> 3
	cap int  -----> 5

	*扩容后  arr,brr 就不共享内存空间了
}
**/
