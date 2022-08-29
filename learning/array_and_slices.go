package main

import "fmt"
func main(){

	// grades:=[3]int{1,2,3}
	// grades:=[...]int{1,2,3}
	var students [3]string
	students[0]="lisa"
	students[1]="hello"
	students[2]="world"
	fmt.Printf("Grades: %v\n",students)
	fmt.Printf("array length: %v\n",len(students))

} 