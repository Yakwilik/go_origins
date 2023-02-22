package handlers

import (
	"GolangCourse/uniq/options"
	"fmt"
)

func showStrMeetCount(lines []string, opts options.Options) (resultLines []string) {
	linesWithMetInfo := []stringWithMetCount{{lines[0], 1}}
	compareRule := getStringCompareRule(opts.IgnoreRegister)
	uniqStrCount := 1
	wordsToSkip := opts.SkippedStringsCount
	charsToSkip := opts.SkippedCharsCount
	for i := 1; i < len(lines); i++ {
		prevLine := lines[i-1]
		currLine := lines[i]
		if !compareRule(processLine(prevLine, wordsToSkip, charsToSkip), processLine(currLine, wordsToSkip, charsToSkip)) {
			linesWithMetInfo = append(linesWithMetInfo, stringWithMetCount{currLine, 1})
			uniqStrCount++
		} else {
			linesWithMetInfo[uniqStrCount-1].metCount++
		}
	}
	return formatResult(linesWithMetInfo)
}

func formatResult(linesWithMetInfo []stringWithMetCount) (result []string) {
	for _, strWithMetInfo := range linesWithMetInfo {
		result = append(result, fmt.Sprintf("%d %s", strWithMetInfo.metCount, strWithMetInfo.str))
	}
	return result
}
