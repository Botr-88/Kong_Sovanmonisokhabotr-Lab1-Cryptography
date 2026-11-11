
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"passwordcracking/utils/crack"
)

func main() {

	// Target SHA-512 hash from Lab #3
	targetHash := "485f5c36c6f8474f53a3b0e361369ee3e32ee0de2f368b87b847dd23cb284b316bb0f026ada27df76c31ae8fa8696708d14b4d8fa352dbd8a31991b90ca5dd38"

	// Open the password wordlist
	file, err := os.Open("nord_vpn.txt")

	if err != nil {
		fmt.Println("Error opening wordlist:", err)
		return
	}

	defer file.Close()

	// Create scanner to read file line by line
	scanner := bufio.NewScanner(file)

	fmt.Println("================================")
	fmt.Println("    SHA-512 PASSWORD CRACKER")
	fmt.Println("================================")
	fmt.Println("Target Hash:", targetHash)
	fmt.Println("Starting password cracking...")
	fmt.Println()

	attempts := 0
	found := false

	// Read every password in the wordlist
	for scanner.Scan() {

		// Get password from current line
		password := strings.TrimSpace(scanner.Text())

		if password == "" {
			continue
		}

		attempts++

		// Generate SHA-512 hash
		generatedHash := crack.SHA512Hash(password)

		// Display verbose output
		fmt.Printf("[Attempt %d] Password: %s | SHA512: %s\n",
			attempts, password, generatedHash)

		// Compare generated hash with target hash
		if generatedHash == targetHash {

			fmt.Println()
			fmt.Println("================================")
			fmt.Println("        PASSWORD FOUND!")
			fmt.Println("================================")
			fmt.Println("Password:", password)
			fmt.Println("SHA-512 Hash:", generatedHash)
			fmt.Println("Total Attempts:", attempts)

			found = true
			break
		}
	}

	// Check for file reading errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading wordlist:", err)
		return
	}

	// If password was not found
	if !found {
		fmt.Println()
		fmt.Println("Password not found in wordlist.")
		fmt.Println("Total Attempts:", attempts)
	}

	fmt.Println()
	fmt.Println("Password cracking completed.")
}

