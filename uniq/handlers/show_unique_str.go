package handlers

import (
	"GolangCourse/uniq/options"
	"GolangCourse/uniq/utils"
)

func ShowUniqueStr(lines []string, opts options.Options) (resultLines []string) {
	linesWithMetInfo := utils.GetStringsWithMetCount(lines, opts.IgnoreRegister, opts.SkippedStringsCount, opts.SkippedCharsCount)
	return utils.FilterStringsByMetCount(linesWithMetInfo, false)
}
