package common

import (
	"crypto/sha256"
	"encoding/hex"
)

// CheckPassword compares the given password with the hashed password and returns true if they match.
func HashPasswordSHA256(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}
