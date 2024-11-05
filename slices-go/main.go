package main

import "fmt"

func main() {
	months := [...]string{"Janurary", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}
	months1 := [...]string{"Janurary", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	fmt.Println("arrays are comparable", months == months1)

	q2 := months1[4:7]
	fmt.Println("the slice of months1 array from 4 to 7 which means we got indexes 4,5,6", q2)
	// fmt.Println("slices are not comparable", q1 == q2)
	fmt.Println("slices are only comparable with nil", q2 != nil)

	fmt.Println("capacity of months", cap(months))

	//slices are only pointers to the underlying array
	//Note how even though the slice q1 is reversed, the underlying array months is also modified
	q1 := months[:3]
	fmt.Println("the slice of months array from 0 to 3 which means we got indexes 0,1,2", q1)
	reverse(q1)
	fmt.Println("the slice q1 is reversed", q1)
	fmt.Println("modifying q1, also modified the underlying array months", months)

	//Note how if the underlying array months1 is modified, the slice q2 is also modified
	fmt.Println("the slice q2 before modfiying months1 array", q2)
	months1[4] = "May-Month"
	months1[5] = "June-Month"
	months1[6] = "July-Month"
	fmt.Println("the slice q2 after modfiying months1 array", q2)

	// in order to append to the array, the capacity is doubled before adding a entity

	ints := []int{0, 1, 2, 3, 4}
	fmt.Println("appending using custom append method; before appending", ints)
	ints = appendTenMore(ints)
	fmt.Println("after appending", ints)

	intMax := ints[len(ints)-1] + 5
	for i := ints[len(ints)-1] + 1; i < intMax; i++ {
		ints = append(ints, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(ints), ints)
	}
	fmt.Println("after appending using built in append", ints)

}

func appendTenMore(x []int) []int {
	var y []int
	y = x
	for i := x[len(x)-1] + 1; i < 10; i++ {
		y = appendInt(y, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
	}
	return y
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		z = make([]int, zlen, 2*len(x))
		copy(z, x)
	}

	z[len(x)] = y
	return z
}

func reverse(s []string) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
