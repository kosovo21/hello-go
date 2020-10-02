package more

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func Range() {
	fmt.Println("== Range ==")
	wc.Test(WordCount)
}

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	splitString := strings.Fields(s)
	for _, v := range splitString {
		m[v] = m[v] + 1
	}
	return m
}

func RangeSimple() {
	fmt.Println("== RangeSimple ==")
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
