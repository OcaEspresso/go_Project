package main

import "fmt"

func main() {

	// 变量的声明  (指定变量的类型)

	var a int
	var b int
	var c int

	var age byte //添加一个虚拟分配 var _ byte
	var sex bool
	var price float64

	//变量的赋值
	a = 10
	b = 20
	c = a + b

	//简写形式  :=  声明和赋值同时完成
	f := 40.0 // go会自动推断f是整型
	g := c - a

	var h int = 4
	var m = 4 //声明和赋值一行完成，但类型也是自动推断出来的

	fmt.Printf("Hello World\n") //println 每行末尾追加换行符
	//fmt.Printf("%d \n", c)       // printf  f是format
	fmt.Println(age, sex, price) // 未赋值变量有自己的默认值
	fmt.Printf("%d   %d   %t   %.2f\n", c, age, sex, price)
	fmt.Printf("%T   %T", f, g) // %T 推断类型
}
