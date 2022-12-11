package main

import (
	"fmt"
)

func main() {
	a := 42
	var b *int = &a
	fmt.Println(a, b)
	fmt.Println(a, *b)
}
