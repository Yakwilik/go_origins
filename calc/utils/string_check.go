package utils

import "strings"

func IsNumber(value string) bool {
	digits := ".0123456789"
	return strings.Contains(digits, value)
}

func IsOpeningBracket(value string) bool {
	return value == "("
}

func IsClosingBracket(value string) bool {
	return value == ")"
}
