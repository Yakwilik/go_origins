package handlers

import (
	"fmt"
	"regexp"
	"strings"
)

func getWordsCount(line string) int {
	return strings.Count(strings.Trim(line, " "), " ") + 1
}

func skipWords(line string, skipWordsCount int) (resultLine string) {
	wordsCount := getWordsCount(line)
	if wordsCount <= skipWordsCount {
		return resultLine
	}
	re := regexp.MustCompile(fmt.Sprintf("^(.*? ){%d}", skipWordsCount))
	return re.ReplaceAllString(line, "")
}
