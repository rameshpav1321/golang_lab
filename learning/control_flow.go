// package main

// func main() {
// 	number := 50
// 	guess := 500
// 	if guess < number {
// 		println("It's low")
// 	} else if guess > number {
// 		println("It's high")
// 	} else {
// 		println("You got it")
// 	}

// }
package main

// func main() {
// 	i := 10
// 	switch {
// 	case i <= 10:
// 		fallthrough
// 	case i <= 20:
// 	default:
// 	}
// }

// ----- or ------

func main() {
	var i interface{} = 1
	switch i.(type) {
	case int:
	case float64:
	case string:
	default:

	}
}
