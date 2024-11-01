package main

import (
	"fmt"
	"package-dep-init/dependencies"
)

func init() {
	fmt.Println("In main init")
}

func main() {
	fmt.Println("Value dependent is", dependencies.GetDep2Val())
}
