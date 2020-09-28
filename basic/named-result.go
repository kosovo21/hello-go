package basic

import "fmt"

func NamedResult() {
	fmt.Println(split(17))
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
