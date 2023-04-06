package hash

import (
	"golang.org/x/crypto/bcrypt"
)

//TODO add the salt string for passwords

// HashPassword hashes the password of user
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(bytes), err
}

// CheckPasswordHash method checks if the
// human-readable password corresponds to a hashed password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
