package more

import "fmt"

type Vertex struct {
	X int
	Y int
}

func Struct() {
	fmt.Println("== Struct ==")
	fmt.Println(Vertex{1, 2})
}

func StructField() {
	fmt.Println("== StructField ==")
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}

func StructPointer() {
	fmt.Println("== StructPointer ==")
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}

var (
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	pX = &Vertex{1, 2} // has type *Vertex
)

func StructLiteral() {
	fmt.Println("== StructLiteral ==")
	fmt.Println(v1, pX, v2, v3)
}
