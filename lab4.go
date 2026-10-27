package main

import "fmt"

func add(a, b float64) float64 { return a + b }
func sub(a, b float64) float64 { return a - b }
func mul(a, b float64) float64 { return a * b }

func div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func mod(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a % b, nil
}

func main() {
	for {
		fmt.Println("\n===== Mini Calculator =====")
		fmt.Println("1) Add  2) Sub  3) Mul  4) Div  5) Mod  6) Exit")
		fmt.Print("Choose: ")

		var choice int
		fmt.Scan(&choice)

		if choice == 6 {
			fmt.Println("Exiting program.")
			break
		}

		if choice < 1 || choice > 6 {
			fmt.Println("Error: Invalid option chosen.")
			continue
		}

		if choice == 5 {
			var a, b int
			fmt.Print("Enter a: ")
			fmt.Scan(&a)
			fmt.Print("Enter b: ")
			fmt.Scan(&b)

			res, err := mod(a, b)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
			} else {
				fmt.Printf("Result: %d\n", res)
			}
			continue
		}

		var a, b float64
		fmt.Print("Enter a: ")
		fmt.Scan(&a)
		fmt.Print("Enter b: ")
		fmt.Scan(&b)

		switch choice {
		case 1:
			fmt.Printf("Result: %g\n", add(a, b))
		case 2:
			fmt.Printf("Result: %g\n", sub(a, b))
		case 3:
			fmt.Printf("Result: %g\n", mul(a, b))
		case 4:
			res, err := div(a, b)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
			} else {
				fmt.Printf("Result: %g\n", res)
			}
		}
	}
}