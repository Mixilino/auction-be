package utils

import (
	"auction-be/internal/constants"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) ([]byte, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), constants.CostHashFactor)
	return hashedPassword, err
}

// CheckPasswordHash returns true if password matches hash
func CheckPasswordHash(password string, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	return err == nil
}
