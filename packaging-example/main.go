package main

import (
	"example-module/conversions"
	"fmt"
)

func main() {
	// Note that the type Celcius is being imported from conversions package
	// Note that the method CToF is being imported from conversions package. Note the method name being capitalized.

	// write a program which converts local variable in celcius to Farenheit and vice versa
	celcius := 123.64
	//cToF(celcius) - throws error - cannot use celcius (variable of type float64) as Celcius value in argument to cToF
	fmt.Println(conversions.CToF(conversions.Celcius(celcius))) // the argument is type casted before passing into the method
}
