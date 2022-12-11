// package main

// import "fmt"
// func main(){
// 	statePopulations:=map[string]int{
// 		"california":39284234,
// 		"texas":23435,
// 		"new york":23433423,
// 	}
// 	fmt.Println(statePopulations)
// }

// to delete : delete(statePoulations,"delhi")
// , ok syntax is used to verify whether a key is in map or not
// as by default go returns 0 value to non-existing keys

// ============================================================================

// package main

// import "fmt"

// type Doctor struct{
// 	number int
// 	actorName string
// 	companions []string

// }
// func main(){
// 	aDoctor:=Doctor{
// 		number: 3,
// 		actorName: "John",
// 		companions: []string{
// 			"Lisa",
// 			"Adam",
// 			"Bruce",
// 		},
// 	}
// 	fmt.Println(aDoctor.actorName)
// }

// =====================================================================

package main

type Human struct{
	name string
	origin string
}

type Doctor struct{
	Human
	licNumber int
	actorName string
	companions []string

}

func main(){
	newDoctor:=Doctor{}
	newDoctor.name="shankar dada"
	newDoctor.origin="hyderabad"
	newDoctor.licNumber=1235730
	newDoctor.actorName="chiranjeevi"
	newDoctor.companions=[]string{
		"sonali bindre",
		"srikanth",
		"paresh rawal",
	}
	println(newDoctor.companions[1])
}