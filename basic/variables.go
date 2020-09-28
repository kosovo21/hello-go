package basic

import "fmt"

var c, python, java bool

var i, j int = 1, 2

func Variables() {
	var x int
	fmt.Println(x, c, python, java)
}

func VarInit() {
	var c, python, java = true, false, "no!"
	fmt.Println(i, j, c, python, java)
}

func ShortVar() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
