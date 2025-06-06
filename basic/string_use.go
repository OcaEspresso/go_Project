package main

import (
	"fmt"
)

func main() {
	s := "golang"
	fmt.Println(s)

	// 字符串本质是一个不可修改的byte数组
	// golang的遍历根据 for value1, value2 := range arr 来做
	//for i, ele := range s {
	//	fmt.Printf("%d   %c\n", i, ele)
	//}

	fmt.Println(len(s))
	a := "golang你好" //一个汉字占3个byte
	fmt.Println(len(a))
	// 强制转换为rune切片
	arr := []rune(a)
	fmt.Println(len(arr)) //8

	for _, ele := range arr {
		fmt.Printf("%c", ele)
	}

	/**
		``反引号原封不动的打印内容
	`xxx
	   xx`  包括这些换行符
	**/
}

//字符串拼接  +即可
