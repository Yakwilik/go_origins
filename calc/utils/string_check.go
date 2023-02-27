package utils

import (
	"bufio"
	"os"
	"strings"
)

func IsPartOfNumber(value string) bool {
	digits := ".0123456789"
	return strings.Contains(digits, value)
}

func IsOpeningBracket(value string) bool {
	return value == "("
}

func IsClosingBracket(value string) bool {
	return value == ")"
}

func GetInput() (string, error) {
	if len(os.Args) > 1 {
		return os.Args[1], nil
	}

	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()

	return scanner.Text(), scanner.Err()
}
