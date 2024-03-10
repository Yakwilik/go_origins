package handlers

import (
	"fmt"

	"GolangCourse/uniq/options"
	"GolangCourse/uniq/utils"
)

func HandleLines(lines []string, opts options.Options) (resultLines []string) {
	switch {
	case opts.EShowStrMeetCount:
		resultLines = showStrMeetCount(lines, opts)
	case opts.EShowNotUniqueStr:
		resultLines = showNotUniqueStr(lines, opts)
	case opts.EShowUniqueStr:
		resultLines = showUniqueStrStrict(lines, opts)
	default:
		resultLines = showUniqueStr(lines, opts)
	}
	return resultLines
}

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

func showNotUniqueStr(lines []string, opts options.Options) (resultLines []string) {
	linesWithMetInfo := utils.GetStringsWithMetCount(lines, opts.IgnoreRegister, opts.SkippedStringsCount, opts.SkippedCharsCount)
	return utils.FilterStringsByMetCount(linesWithMetInfo, true)
}

func showUniqueStr(lines []string, opts options.Options) (resultLines []string) {
	linesWithMetInfo := utils.GetStringsWithMetCount(lines, opts.IgnoreRegister, opts.SkippedStringsCount, opts.SkippedCharsCount)
	return utils.FilterStringsByMetCount(linesWithMetInfo, false)
}

func showUniqueStrStrict(lines []string, opts options.Options) (resultLines []string) {
	linesWithMetInfo := utils.GetStringsWithMetCount(lines, opts.IgnoreRegister, opts.SkippedStringsCount, opts.SkippedCharsCount)
	return utils.FilterUniqueStrings(linesWithMetInfo)
}
