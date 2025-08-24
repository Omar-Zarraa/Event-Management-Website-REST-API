package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	return string(hash), err
}

func CheckPasswordHash(pass, hashedPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(pass))
	return err == nil
}
