package main

import (
	"fmt"
)

// func main() {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(i)
// 	}

// }

// looping for loop over collections

func main() {
	arr := []int{4, 5, 6}
	for k, v := range arr {
		fmt.Println(k, v)
	}
}
