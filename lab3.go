package main

import "fmt"

func myAND(a, b int) {
	fmt.Printf("Bitwise AND (%d & %d) = %d (binary: %08b)\n", a, b, a&b, a&b)
}

func myOR(a, b int) {
	fmt.Printf("Bitwise OR  (%d | %d) = %d (binary: %08b)\n", a, b, a|b, a|b)
}

func myXOR(a, b int) {
	fmt.Printf("Bitwise XOR (%d ^ %d) = %d (binary: %08b)\n", a, b, a^b, a^b)
}

func myNOT(a, b int) {
	fmt.Printf("Bitwise NOT (^%d) = %d\n", a, ^a)
	fmt.Printf("Bitwise NOT (^%d) = %d\n", b, ^b)
}

func myShifts(a int, shift uint) {
	fmt.Printf("Left shift  (%d << %d) = %d\n", a, shift, a<<shift)
	fmt.Printf("Right shift (%d >> %d) = %d\n", a, shift, a>>shift)
}

func demonstrateAssignmentOperators(a int) {
	fmt.Printf("\n--- Assignment Operators Demo on %d ---\n", a)
	temp := a
	temp &= 3
	fmt.Printf("temp &= 3  -> %d\n", temp)

	temp = a
	temp |= 3
	fmt.Printf("temp |= 3  -> %d\n", temp)

	temp = a
	temp ^= 3
	fmt.Printf("temp ^= 3  -> %d\n", temp)

	temp = a
	temp <<= 1
	fmt.Printf("temp <<= 1 -> %d\n", temp)

	temp = a
	temp >>= 1
	fmt.Printf("temp >>= 1 -> %d\n", temp)
}

func main() {
	var a, b int
	fmt.Print("Enter first integer (a): ")
	fmt.Scan(&a)
	fmt.Print("Enter second integer (b): ")
	fmt.Scan(&b)

	fmt.Printf("\nInputs: a = %d (%08b), b = %d (%08b)\n\n", a, a, b, b)

	myXOR(a, b)
	myNOT(a, b)
	myOR(a, b)
	myAND(a, b)
	myShifts(a, 2)

	demonstrateAssignmentOperators(a)
}