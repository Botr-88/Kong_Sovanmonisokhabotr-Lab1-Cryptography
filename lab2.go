package main

import "fmt"

func main() {
	var a, b int

	// User input
	fmt.Print("Enter first integer (a): ")
	// fmt.Scan(&a)

	fmt.Print("Enter second integer (b): ")
	fmt.Scan(&b)

	fmt.Printf("\nEvaluating inputs: a = %d, b = %d\n\n", a, b)

	// 1. Both positive (&&)
	
	bothPositive := (a> 0) && (b> 0)
	fmt.Printf("Both are positive (a> 0 && b> 0): %t\n", bothPositive)

	// 2. One greater than the other (||)
	findGreater := (a> b) || (b> a)
	fmt.Printf("One is greater than the other (a > b || b > a): %t\n", findGreater)

	// 3. Not equal (!)
	notEqual := !(a== b)
	fmt.Printf("Not equal (!(a == b)): %t\n", notEqual)
}