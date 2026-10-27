package main

import (
	"bufio"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimRight(text, "\r\n")

	data := []byte(text)

	// 1. Binary representation (each byte formatted to 8 bits)
	var binaryStr strings.Builder
	
	for _, b := range data {
		binaryStr.WriteString(fmt.Sprintf("%08b ", b))
	}

	// 2. Hexadecimal representation

	hexStr := hex.EncodeToString(data)

	// 3. Base64 representation

	base64Str := base64.StdEncoding.EncodeToString(data)

	fmt.Printf("\nOriginal:    %s\n", text)
	fmt.Printf("Binary:      %s\n", binaryStr.String())
	fmt.Printf("Hexadecimal: %s\n", hexStr)
	fmt.Printf("Base64:      %s\n", base64Str)
}