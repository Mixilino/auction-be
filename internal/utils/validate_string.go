package utils

import (
	"auction-be/internal/errors/resterrors"
	"unicode"
)

func ValidatePassword(password string) error {
	if len(password) < 8 || len(password) > 255 {
		return resterrors.ConstPasswordNotStrongEnough
	}
	hasLower := false
	hasUpper := false
	hasDigit := false
	for _, char := range password {
		if unicode.IsLower(char) {
			hasLower = true
		} else if unicode.IsUpper(char) {
			hasUpper = true
		} else if unicode.IsDigit(char) {
			hasDigit = true
		}

		if hasLower && hasUpper && hasDigit {
			return nil
		}
	}
	return resterrors.ConstPasswordNotStrongEnough
}
