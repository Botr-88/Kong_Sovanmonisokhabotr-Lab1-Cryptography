
package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"os"

	"crypto/sha3"
)

func proofMe(txt1, txt2 string) {

	// MD5
	hashA := md5.Sum([]byte(txt1))
	hashB := md5.Sum([]byte(txt2))
	compareHash("MD5", hashA[:], hashB[:])

	// SHA1
	hashC := sha1.Sum([]byte(txt1))
	hashD := sha1.Sum([]byte(txt2))
	compareHash("SHA1", hashC[:], hashD[:])

	// SHA256
	hashE := sha256.Sum256([]byte(txt1))
	hashF := sha256.Sum256([]byte(txt2))
	compareHash("SHA256", hashE[:], hashF[:])

	// SHA512
	hashG := sha512.Sum512([]byte(txt1))
	hashH := sha512.Sum512([]byte(txt2))
	compareHash("SHA512", hashG[:], hashH[:])

	// SHA3-256
	hashI := sha3.Sum256([]byte(txt1))
	hashJ := sha3.Sum256([]byte(txt2))
	compareHash("SHA3", hashI[:], hashJ[:])
}

// Function to compare two hash values
func compareHash(name string, outputA, outputB []byte) {

	fmt.Printf("\nHash (%s):\n", name)
	fmt.Println("Output A:", hex.EncodeToString(outputA))
	fmt.Println("Output B:", hex.EncodeToString(outputB))

	if hex.EncodeToString(outputA) == hex.EncodeToString(outputB) {
		fmt.Println("Result: Match!")
	} else {
		fmt.Println("Result: No Match!")
	}
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("======== Name + Hashing Program ========")

	fmt.Print("Please input value 1: ")
	input1, _ := reader.ReadString('\n')

	fmt.Print("Please input value 2: ")
	input2, _ := reader.ReadString('\n')

	// Remove newline characters
	input1 = input1[:len(input1)-1]
	input2 = input2[:len(input2)-1]

	proofMe(input1, input2)
}
