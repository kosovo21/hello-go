package basic

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func MultiResult() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
