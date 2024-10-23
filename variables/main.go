package main

import "fmt"

//RandomGlobal := "random";
var RandomGlobal = "random" // Convention is that Capitalized variable name denotes it is a public variable.

func main() {
	var number int = 4
	fmt.Println("number is", number)

	// the variable can be declared without specifying the type and still work
	// but the once a value of certain type is assigned to the variable, the type cannot be changed
	// the compiler would throw an error. Try to uncomment line 15 to check.

	var unInitNum = 5
	fmt.Println("unInitNum", unInitNum)
	// unInitNum = 4.3

	// the walrus operator can be used to assign a variable without specifying its type and without using var keyword
	// note the this can be used only inside the methods and not outside it.
	// the compiler would throw an error. Try to uncomment line 5 to check.

	randomString := "fassfga"
	fmt.Println("randomString", randomString)

	fmt.Println("RandomGlobal", RandomGlobal)

}
