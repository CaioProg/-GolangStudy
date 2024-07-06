package main

import "fmt"

var globalVariable = "Hello, I'm a global variable"

func main() {

	// Explicit statement

	var name1 string = "Caio"
	var age1 int = 21
	var height1 float64 = 1.82
	var isDeveloper1 bool = true

	fmt.Println(name1, age1, height1, isDeveloper1)

	// Implied statement

	name2 := "Caio"
	age2 := 21
	height2 := 1.82
	isDeveloper2 := true

	fmt.Println(name2, age2, height2, isDeveloper2)

	// Uninitialized variables

	var name3 string
	var age3 int
	var height3 float64
	var isDeveloper3 bool

	fmt.Println(name3, age3, height3, isDeveloper3)

	// Multiple variables

	var x, y int = 1, 2
	var a, b, c = 3, "Hello", 4.5

	fmt.Println(x, y, a, b, c)

	// Global variable

	fmt.Println(globalVariable)

	// Constants

	const Pi = 3.14
	const Truth = true

	fmt.Println(Pi, Truth)
}
