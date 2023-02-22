package handlers

import (
	"GolangCourse/uniq/options"
)

func HandleLines(lines []string, opts options.Options) (resultLines []string) {
	switch {
	case opts.EShowStrMeetCount:
		resultLines = showStrMeetCount(lines, opts)
		break
	case opts.EShowNotUniqueStr:
		break
		//return getUniqLines(lines, opts)
	case opts.EShowUniqueStr:
		break
		//return getDuplicateLines(lines, opts)
	default:
		//return defaultHandlerLines(lines, opts)
	}
	return resultLines
}
