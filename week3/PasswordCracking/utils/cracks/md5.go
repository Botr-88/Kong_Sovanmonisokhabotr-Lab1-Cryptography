
package crack

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

// CrackMD5 checks every password in the wordlist
func CrackMD5(targetHash string, wordlistPath string, verbosePath string) (string, bool, error) {

	// Open the wordlist file
	file, err := os.Open(wordlistPath)
	if err != nil {
		return "", false, err
	}
	defer file.Close()

	// Create a file to save verbose output
	logFile, err := os.Create(verbosePath)
	if err != nil {
		return "", false, err
	}
	defer logFile.Close()

	// Read the wordlist line by line
	scanner := bufio.NewScanner(file)

	// Allow longer lines in the wordlist
	scanner.Buffer(make([]byte, 64*1024), 1024*1024)

	attempt := 0

	fmt.Fprintln(logFile, "======== MD5 Password Cracking ========")
	fmt.Fprintln(logFile, "Target Hash:", targetHash)

	// Loop through every password
	for scanner.Scan() {

		password := strings.TrimSuffix(scanner.Text(), "\r")

		if password == "" {
			continue
		}

		attempt++

		// Calculate MD5 hash
		hash := md5.Sum([]byte(password))

		// Convert hash into hexadecimal
		hashText := hex.EncodeToString(hash[:])

		// Compare with target hash
		if hashText == strings.ToLower(targetHash) {

			fmt.Printf("\n[+] PASSWORD FOUND: %s\n", password)

			fmt.Fprintf(logFile,
				"[%d] %s -> %s : MATCH!\n",
				attempt, password, hashText)

			fmt.Fprintln(logFile, "\nPASSWORD FOUND:", password)
			fmt.Fprintln(logFile, "Total attempts:", attempt)

			return password, true, nil
		}

		// Save unsuccessful attempts
		fmt.Fprintf(logFile,
			"[%d] %s -> %s : NO MATCH\n",
			attempt, password, hashText)
	}

	if err := scanner.Err(); err != nil {
		return "", false, err
	}

	fmt.Println("\n[-] Password not found!")

	fmt.Fprintln(logFile, "\nPassword not found!")
	fmt.Fprintln(logFile, "Total attempts:", attempt)

	return "", false, nil
}
