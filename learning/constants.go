package main

import (
	"fmt"
)
const(
	a=iota
	b 
	c
)
func main(){

	const myConst int=42
	fmt.Printf("%v,%T\n",myConst,myConst)
	fmt.Printf("%v\n",a)
	fmt.Printf("%v\n",b)
	fmt.Printf("%v\n",c)

}