package handlers

import (
	"GolangCourse/uniq/options"
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
