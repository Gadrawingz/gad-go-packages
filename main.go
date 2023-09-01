package main

import (
	"fmt"
	"goPackages/calculator"
	"goPackages/simpleInterest"
)

// I am using blank identifier (_) for
// Unused packages
var _ = "goPackages/unusedPackage"

// Above line mutes the error!

func init() {
	fmt.Println("Simple interest package has been initialized!")
}

func main() {
	fmt.Println("The program for simple interest!")
	p := 5000.0
	r := 10.0
	t := 1.0

	n1, n2 := 55, 5

	simpleInt := simpleInterest.CalculateInterest(p, r, t)
	fmt.Println("The simple interest is: ", simpleInt)

	addition := calculator.AddNumbers(float64(n1), float64(n2))
	fmt.Printf("The addition gives: %f \n", addition)

	multiply := calculator.MultiplyNumbers(float64(n1), float64(n2))
	fmt.Printf("The multiplication gives: %f \n", multiply)

}
