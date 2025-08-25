// Package utils contains password hashing utilities
package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword takes the password as a parameter and hashes it using the 'bcrypt' library, returns the hashed password and possibly an error.
func HashPassword(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(hash), err
}

// CheckPasswordHash takes the password and the hashed password as parameters and checks if the hash is valid for the plain text password given, returns a boolean.
func CheckPasswordHash(pass, hashedPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pass))
	return err == nil
}
