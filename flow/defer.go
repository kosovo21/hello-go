package flow

import "fmt"

func Defer() {
	fmt.Println("== Defer ==")
	defer fmt.Println("world")

	fmt.Println("hello")
}

func DeferStack() {
	fmt.Println("== DeferStack ==")
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
