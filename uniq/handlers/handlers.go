package handlers

import (
	"GolangCourse/uniq/options"
)

func HandleLines(lines []string, opts options.Options) (resultLines []string) {
	switch {
	case opts.EShowStrMeetCount:
		resultLines = showStrMeetCount(lines, opts)
	case opts.EShowNotUniqueStr:
		resultLines = ShowNotUniqueStr(lines, opts)
	case opts.EShowUniqueStr:
		resultLines = ShowUniqueStrStrict(lines, opts)
	default:
		resultLines = ShowUniqueStr(lines, opts)
	}
	return resultLines
}
