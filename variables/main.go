package main

import "fmt"

//RandomGlobal := "random";
var RandomGlobal = "random" // Convention is that Capitalized variable name denotes it is a public variable.

//rune = int32
//byte = uint8 - signifies a piece of raw data

func main() {
	var number int = 4
	fmt.Println("number is", number)

	// the variable can be declared without specifying the type and still work
	// but the once a value of certain type is assigned to the variable, the type cannot be changed
	// the compiler would throw an error. Try to uncomment line 15 to check.

	var unInitNum = 5
	fmt.Println("unInitNum", unInitNum)
	// unInitNum = 4.3;

	// the walrus operator can be used to assign a variable without specifying its type and without using var keyword
	// note the this can be used only inside the methods and not outside it.
	// the compiler would throw an error. Try to uncomment line 5 to check.

	rrandomString := "fassfga"
	fmt.Println("randomString", rrandomString)
	fmt.Println("RandomGlobal", RandomGlobal)

	// note that the short variable declaration doesn't work for assignment of an old variable
	// the compiler would throw an error. Try to uncomment line 30 to check.

	//rrandomString := "aasfdf"

	// assigning multiple variables of same time can be done with a single assingnment opertor. look line 36
	// swapping variable values can be done with assignment operator ('='). Look at line 38.

	var i, j int = 5, 10
	fmt.Println(i, j)
	j, i = i, j // this is called Tuple assignment
	fmt.Println(i, j)

	// when assigning multiple values with short declaration operator,
	// atleast one of the variable on the LHS, should be a new variable
	// the compiler would throw an error. Try to uncomment line 47 to check.

	z, i := 15, 20
	fmt.Println(z, i)
	//i,z := 20,15;

	// Pointers are the address at which the variables are stored.
	// To get the pointer of a variable use &z. This returns type *int (pointer to int) - because z is a pointer here.

	point := &z
	fmt.Println(*point)
	fmt.Println(point == &z)

	// new function to create an unnamed variable of any type.
	// returns *type - returns the address at which the variable lives
	// its value can be obtains by points to operator - check line 62

	unnamed := new(int)
	fmt.Println(unnamed)
	fmt.Println(*unnamed)

	arr := []string{"gold", "silver", "bronze"}
	fmt.Println(arr[0])
	fmt.Println(arr[1])
	fmt.Println(arr[2])
	arr[2] = "Pithalai"
	fmt.Println(arr[2])
	//arr[1] = 4            - for assignability : types on both sides should match
	// if(arr[0] != 3) {    - for comparability : types on both sides should match
	// 	arr[2] = "gold"
	// }

	// If the result of an arithmetic operation, whether signed or unsigned,has more bits than can be
	// represented in the result type, it is said to overflow. The high-order bits that do not fit are
	// silently discarded.
	// If the original number is asigned type, the result could be negative if the
	// leftmost bit is a 1, as in the int8 example here:

	var unsignedInt uint8 = 255
	fmt.Println(unsignedInt, unsignedInt+1, unsignedInt*unsignedInt) // "255 0 1"
	var signedInt int8 = 127
	fmt.Println(signedInt, signedInt+1, signedInt*signedInt) // "127 -128 1

}
