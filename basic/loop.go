package main

import "fmt"

func main() {
	// i :=0
	// for ; i <3; i++  这样也可以
	//for i := 0; i < 3; i++ {
	//	fmt.Println(i)
	//}
	i := 0
	for {
		i++
		if i%2 == 0 {
			continue
		}

		fmt.Println(i)

		if i > 10 {
			break
		}
	}

}
