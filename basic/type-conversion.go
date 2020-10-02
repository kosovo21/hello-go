package basic

import (
	"fmt"
	"math"
)

func TypeConvert() {
	fmt.Println("== TypeConvert ==")
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}

func TypeInterface() {
	fmt.Println("== TypeInterface ==")
	v := "X" // change me!
	fmt.Printf("v is of type %T\n", v)
}
