package main

import (
	"errors"
	"fmt"
)

func main() {
	// if a slice is passed to the function, then if the function modifies the input variable,
	// it affects the slice. Because slice is essentially a pointer to the underlying array.
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)
	modifySlice(slice)
	fmt.Println(slice)

	// if a pointer to an array is passed to the function, then if the function modifies the input variable,
	// it affects the array. The reason is same as above.
	arrWithPointer := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arrWithPointer)
	modifyArrPointer(&arrWithPointer)
	fmt.Println(arrWithPointer)

	// if an array is passed to the function, then if the function modifies the input variable,
	// it doesn't affects the array. The reason is same as above.
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	modifyArray(arr)
	fmt.Println(arr)

	// To modify the array with the help of the function, set the return type of the function same as the input array
	// Then set the array variable with the result of the function.
	arr = modifyArray(arr)
	fmt.Println(arr)

	if value, err := lookupIfPossible(arr[0:], 10); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("The looked up element is %d", value)
	}

	if value, err := lookupIfPossibleWithNamedReturns(arr[0:], -1); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("The looked up element is %d\n", value)
	}

	// anonyomous fuction - functions as a reference types

	var funcStore1 []func() // array to store functions
	for _, ele := range arr {
		local := ele
		funcStore1 = append(funcStore1, func() {
			fmt.Println(local)
		})
	}

	for _, v := range funcStore1 {
		v()
	}

	var funcStore2 []func() // array to store functions
	for _, ele := range arr {
		funcStore2 = append(funcStore2, func() {
			fmt.Println(ele)
		})
	}

	for _, v := range funcStore2 {
		v()
	}

	// variadic functions - function that can take many input variables. Similar to [...args] in javascript
	fmt.Println("variadic sum", variadicSum(slice...))

	deferExample()

	// example of panic and recover
	divide(4, 0)

	fmt.Println(noReturn())

}

func modifyArray(array [5]int) [5]int {
	for i := 3; i < len(array); i++ {
		array[i] *= 20
	}
	return array
}

func modifySlice(pointer []int) {
	for i := 3; i < len(pointer); i++ {
		pointer[i] *= 2
	}
}

func modifyArrPointer(pointer *[5]int) {
	for i := 3; i < len(pointer); i++ {
		pointer[i] *= 10
	}
}

func lookupIfPossible(slice []int, index int) (int, error) {
	if index < 0 || index >= len(slice) {
		return 0, errors.New("the index is out of bounds")
	}

	return slice[index], nil
}

func lookupIfPossibleWithNamedReturns(slice []int, index int) (value int, err error) {
	if index < 0 || index >= len(slice) {
		err = fmt.Errorf("the index %d is out of bounds for the slice being attempted to lookup", index)
	} else {
		value = slice[index]
	}
	return
}

func variadicSum(values ...int) int {
	fmt.Println("Summing up", values)
	var total int
	for _, v := range values {
		total += v
	}
	return total
}

func deferExample() {
	fmt.Println("Line 1")

	// defer functions will execute only after execution of normal functions
	// defer functions will be executed in Last In First Out fashion
	// So in this case "Defered line 1" is the last in. So will execute first
	defer fmt.Println("Defered line 2")
	defer fmt.Println("Defered line 1")

	fmt.Println("Line 2")
}

func divide(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Error in division: %v\n", r)
		}
	}()
	if b == 0 {
		panic("division by zero!")
	}
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
}

func noReturn() (result int) {
	defer func() {
		if r := recover(); r != nil {
			// Modify the named return value in the deferred function
			result = r.(int)
		}
	}()

	// Panic with the value we want to "return"
	panic(42)
}
