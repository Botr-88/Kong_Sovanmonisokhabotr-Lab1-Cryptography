package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
	var a, b int

	// Prompt user for input
	fmt.Print("Enter first integer (a): ")
	fmt.Scan(&a)

	fmt.Print("Enter second integer (b): ")
	fmt.Scan(&b)

	fmt.Printf("\nInitial values: a = %d, b = %d\n\n", a, b)

	// Simple assignment (=)
	result := a
	fmt.Printf("Simple Assignment (result = a): %d\n", result)

	// Addition assignment (+=)
	result += b
	fmt.Printf("Add and assign (result += b): %d\n", result)

	// Subtraction assignment (-=)
	result -= b
	fmt.Printf("Subtract and assign (result -= b): %d\n", result)

	// Multiplication assignment (*=)
	result *= b
	fmt.Printf("Multiply and assign (result *= b): %d\n", result)

	// Division and Modulus (with zero-division check)
	if b != 0 {
		// Division assignment (/=)
		result /= b
		fmt.Printf("Divide and assign (result /= b): %d\n", result)

		// Modulus assignment (%=)
		result %= b
		fmt.Printf("Modulus and assign (result %%= b): %d\n", result)
	} else {
		fmt.Println("Division and modulus operations skipped (division by zero is not allowed).")
	}
}