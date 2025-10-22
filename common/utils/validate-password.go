package utils

import (
	"strings"
	"unicode"
)

func ValidatePassword(s string) bool {
	s = strings.TrimSpace(s)
	if len(s) < 8 {
		return false
	}

	hasUpper := false
	hasLower := false
	hasDigit := false
	hasSpecial := false

	for _, r := range s {
		switch {
		case unicode.IsUpper(r):
			hasUpper = true
		case unicode.IsLower(r):
			hasLower = true
		case unicode.IsDigit(r):
			hasDigit = true
		case strings.ContainsRune("!@#$%^&*()-_=+[]{}|;:'\",.<>?/~`", r):
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasDigit && hasSpecial
}