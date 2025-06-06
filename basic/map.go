package main

import (
	"fmt"
)

func main() {

	var m map[string]int
	m = make(map[string]int, 100) // map扩容比slice扩容 影响性能，所以一开始要估计好

}
