package basic

import (
	"fmt"
	"math/rand"
)

func Package() {
	fmt.Println("== Package ==")
	fmt.Println("My favorite number is", rand.Intn(10))
}
