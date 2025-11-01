package main

import "fmt"

func main() {

	// Integer variable short form
	intNum := 32767 + 1
	fmt.Println(intNum)

	// Integer variable explicit type
	var intTwo int32 = 2
	fmt.Println(intTwo)

	// Integer variable with var keyword and no type
	var intThree = 5
	fmt.Println(intThree)

	var name string
	// var age int32
	// var salary float32
	// var bool bool

	name = "yacquub"

	fmt.Println(len(name))

	// Constant variables unchanging NEED to set a value when defining
	const myConst string = "const value"
	const pi float32 = 3.1415

	fmt.Println(myConst, pi)

}
