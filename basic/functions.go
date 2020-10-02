package basic

import "fmt"

func add(x int, y int) int {
	return x + y
}

func add2(x, y int) int {
	return x + y
}

func Functions() {
	fmt.Println("== Functions ==")
	fmt.Println(add(42, 13))
	fmt.Println(add2(43, 14))
}
