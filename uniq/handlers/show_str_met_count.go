package handlers

import (
	"GolangCourse/uniq/options"
	"GolangCourse/uniq/utils"
	"fmt"
)

func formatMeetCountResult(linesWithMetInfo []utils.StringWithMetCount) (result []string) {
	for _, strWithMetInfo := range linesWithMetInfo {
		result = append(result, fmt.Sprintf("%d %s", strWithMetInfo.MetCount, strWithMetInfo.Str))
	}
	return result
}

func showStrMeetCount(lines []string, opts options.Options) (resultLines []string) {
	linesWithMetInfo := utils.GetStringsWithMetCount(lines, opts.IgnoreRegister, opts.SkippedStringsCount, opts.SkippedCharsCount)
	return formatMeetCountResult(linesWithMetInfo)
}
