package main

import (
	"fmt"
)

func main() {
	//a := 10
	//b := true
	//a, b := 10, true //可以这样一行简写多个变量

	if a, b := 10, true; a > 5 { // 可以简写为 if a, b := 10, true  ;  a >5   --->分号之前进行初始化
		fmt.Println("a > 5")
	} else if b {
		fmt.Println("b => true")
	} else {
		fmt.Println("b => false")
	}

}
