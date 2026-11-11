
package crack

import (
	"crypto/sha512"
	"encoding/hex"
)

// SHA512Hash converts a password into a SHA-512 hash
func SHA512Hash(password string) string {

	// Convert password into bytes and generate SHA-512
	hash := sha512.Sum512([]byte(password))

	// Convert hash bytes into hexadecimal string
	return hex.EncodeToString(hash[:])
}
