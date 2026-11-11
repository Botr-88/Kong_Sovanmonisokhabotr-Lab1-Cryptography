
package main

import (
	"fmt"
	"passwordcracking/utils/cracks"
)

func main() {

	fmt.Println("======== MD5 Password Cracking Program ========")

	// Target MD5 hash from the lab
	targetHash := "6a85dfd77d9cb35770c9dc6728d73d3f"

	// Wordlist location
	wordlistPath := "wordlists/nord_vpn.txt"

	// Verbose output location
	verbosePath := "verbose.txt"

	fmt.Println("Target Hash:", targetHash)
	fmt.Println("Wordlist:", wordlistPath)

	// Call the MD5 cracking function
	password, found, err := crack.CrackMD5(
		targetHash,
		wordlistPath,
		verbosePath,
	)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if found {
		fmt.Println("Cracked Password:", password)
	} else {
		fmt.Println("Try another wordlist.")
	}

	fmt.Println("Verbose output saved to:", verbosePath)
}
