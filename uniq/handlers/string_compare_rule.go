package handlers

import (
	"strings"
)

func getStringCompareRule(ignoreRegister bool) func(lhs string, rhs string) bool {
	if ignoreRegister {
		return strings.EqualFold
	}
	return func(lhs string, rhs string) bool {
		return lhs == rhs
	}
}
