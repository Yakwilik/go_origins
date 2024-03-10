package utils

import (
	"fmt"
	"regexp"
	"strings"
)

func getWordsCount(line string) int {
	return strings.Count(strings.Trim(line, " "), " ") + 1
}

func SkipWords(line string, skipWordsCount int) (resultLine string) {
	wordsCount := getWordsCount(line)
	if wordsCount <= skipWordsCount {
		return resultLine
	}
	re := regexp.MustCompile(fmt.Sprintf("^(.*? ){%d}", skipWordsCount))
	return re.ReplaceAllString(strings.TrimLeft(line, " "), "")
}

func SkipChars(line string, skipCharsCount int) (resultLine string) {
	if len(line) < skipCharsCount {
		return resultLine
	}
	return line[skipCharsCount:]
}

func SkipWordsAndChars(line string, skipWordsCount, skipCharsCount int) string {
	return SkipChars(SkipWords(line, skipWordsCount), skipCharsCount)
}

func GetStringCompareFunc(ignoreRegister bool) func(lhs string, rhs string) bool {
	if ignoreRegister {
		return strings.EqualFold
	}
	return func(lhs string, rhs string) bool {
		return lhs == rhs
	}
}
