package main

import "fmt"

var Global *int

// TO - DO - write comments explaining how garbage collection works

func main() {
	assignToGlobal()
	localVariables()
}

func assignToGlobal() {
	// TO - DO - write comments why x is preserved in the memory
	x := 5
	Global = &x
}

func localVariables() {
	// TO - DO - write comments why y is not preserved in the memory
	y := 10
	fmt.Println("print local variable", y)
}
