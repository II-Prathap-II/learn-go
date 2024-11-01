package dependencies

import "fmt"

var dep1Val = 1

func init() {
	fmt.Println("In dependency 1 init, multiplying value by 2")
	dep1Val *= 2
}

func GetDep1Val() int {
	fmt.Println("In get 1")
	return dep1Val
}
