package flow

import "fmt"

func For() {
	fmt.Println("== For ==")
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func ForAsWhile() {
	fmt.Println("== ForAsWhile ==")
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func Forever() {
	fmt.Println("== Forever ==")
	for {
	}
}
