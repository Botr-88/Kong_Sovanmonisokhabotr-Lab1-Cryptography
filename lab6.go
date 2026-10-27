package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

// Single byte XOR encryption/decryption function
func xorEncrypt(text string, key byte) string {
	bytes := []byte(text)
	result := make([]byte, len(bytes))

	for i := 0; i < len(bytes); i++ {
		result[i] = bytes[i] ^ key
	}

	return string(result)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter plaintext: ")
	plaintext, _ := reader.ReadString('\n')
	plaintext = strings.TrimRight(plaintext, "\r\n")

	var keyChar string
	fmt.Print("Enter a single key character: ")
	fmt.Scan(&keyChar)

	if len(keyChar) == 0 {
		fmt.Println("Error: Key cannot be empty.")
		return
	}
	key := keyChar[0]

	// Encrypt
	ciphertext := xorEncrypt(plaintext, key)

	// Decrypt using the same function
	decrypted := xorEncrypt(ciphertext, key)

	fmt.Println("\n--- Results ---")
	fmt.Printf("Plaintext:          %s\n", plaintext)
	fmt.Printf("Key:                %c (0x%02x)\n", key, key)
	// Ciphertext may contain non-printable control characters, so printing hex is standard:
	fmt.Printf("Ciphertext (hex):   %s\n", hex.EncodeToString([]byte(ciphertext)))
	fmt.Printf("Decrypted text:     %s\n", decrypted)
}