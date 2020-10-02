package exercise

import "fmt"

func Sqrt(x float64) float64 {
	i, z := 1.0, 1.0
	for {
		tmp := z
		z -= (z*z - x) / (2 * z)
		if z == tmp || i == 10 {
			return z
		}
		i++
	}
}

func Loop() {
	fmt.Println("== Loop ==")
	fmt.Println(Sqrt(3))
}
