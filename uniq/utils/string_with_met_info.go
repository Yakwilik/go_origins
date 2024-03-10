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

func FilterStringsByMetCount(linesWithMetInfo []StringWithMetCount, duplicates bool) (filteredLines []string) {
	for _, lineWithMetInfo := range linesWithMetInfo {
		if !duplicates {
			filteredLines = append(filteredLines, lineWithMetInfo.Str)
		} else if lineWithMetInfo.MetCount > 1 {
			filteredLines = append(filteredLines, lineWithMetInfo.Str)
		}

	}
	return filteredLines
}

func FilterUniqueStrings(linesWithMetInfo []StringWithMetCount) (filteredLines []string) {
	stringMetCount := make(map[string]int, 0)
	for _, stringWithMetCount := range linesWithMetInfo {
		if stringWithMetCount.MetCount == 1 {
			stringMetCount[stringWithMetCount.Str]++
		}
	}
	for _, stringWithMetCount := range linesWithMetInfo {
		if stringMetCount[stringWithMetCount.Str] == 1 {
			filteredLines = append(filteredLines, stringWithMetCount.Str)
		}
	}
	return filteredLines
}
