package more

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m1 map[string]Vertex

func Map() {
	fmt.Println("== Map ==")
	m1 = make(map[string]Vertex)
	m1["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m1["Bell Labs"])
}

var m2 = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

var m3 = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func MapLiteral() {
	fmt.Println("== MapLiteral ==")
	fmt.Println(m2)
	fmt.Println(m3)
}

func MapAction() {
	fmt.Println("== MapAction ==")
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, okz := m["Answer"]
	fmt.Println("The value:", v, "Present?", okz)
}
