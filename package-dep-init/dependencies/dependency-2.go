package dependencies

import (
	"fmt"
)

var dep2Val int

func init() {
	fmt.Println("In dependency 2 init, multiplying value by 4")
	dep2Val = GetDep1Val() * 4
}

func GetDep2Val() int {
	fmt.Println("In get 2")
	return dep2Val
}
