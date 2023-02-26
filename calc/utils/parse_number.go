package utils

import "fmt"

func ParseNumber(str string, currentIndex int) (res string, parsedRunes int, err error) {
	metDot := false
	for currentIndex < len(str) && IsNumber(string(str[currentIndex])) {
		token := string(str[currentIndex])
		if token == "." {
			if metDot {
				return res, parsedRunes, fmt.Errorf("встреча второй точки "+
					"во время парсинга числа по индексу %d, исходная строка – %s", currentIndex+parsedRunes, str)
			}
			metDot = true
			if res == "" {
				res += "0"
			}
		}
		res += token
		parsedRunes++
		currentIndex++
	}
	return res, parsedRunes - 1, nil
}
