package handlers

import (
	"GolangCourse/uniq/options"
	"GolangCourse/uniq/utils"
)

func showUniqueStrStrict(lines []string, opts options.Options) (resultLines []string) {
	linesWithMetInfo := utils.GetStringsWithMetCount(lines, opts.IgnoreRegister, opts.SkippedStringsCount, opts.SkippedCharsCount)
	return utils.FilterUniqueStrings(linesWithMetInfo)
}
