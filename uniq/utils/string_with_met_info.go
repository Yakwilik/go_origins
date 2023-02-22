package utils

type StringWithMetCount struct {
	Str      string
	MetCount int
}

func GetStringsWithMetCount(lines []string, ignoreRegister bool, wordsToSkip, charsToSkip int) (linesWithMetInfo []StringWithMetCount) {
	if len(lines) < 1 {
		return linesWithMetInfo
	}
	linesWithMetInfo = append(linesWithMetInfo, StringWithMetCount{
		Str:      lines[0],
		MetCount: 1,
	})

	compareRule := GetStringCompareFunc(ignoreRegister)
	uniqStrCount := 1
	for i := 1; i < len(lines); i++ {
		prevLine := lines[i-1]
		currLine := lines[i]
		if !compareRule(SkipWordsAndChars(prevLine, wordsToSkip, charsToSkip), SkipWordsAndChars(currLine, wordsToSkip, charsToSkip)) {
			linesWithMetInfo = append(linesWithMetInfo, StringWithMetCount{currLine, 1})
			uniqStrCount++
		} else {
			linesWithMetInfo[uniqStrCount-1].MetCount++
		}
	}
	return linesWithMetInfo
}
