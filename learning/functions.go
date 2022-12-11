package main

import "fmt"

func main() {
	g := greeter{
		greeting: "Let's",
		name:     "Go!",
	}
	g.greet()
}

type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}

// func main() {
// 	s := sum(1, 2, 3, 4, 5)
// 	fmt.Println("The sum is:", s)
// }

// func sum(values ...int) int {
// 	fmt.Println(values)
// 	result := 0
// 	for _, v := range values {
// 		result += v
// 	}
// 	return result
// }
