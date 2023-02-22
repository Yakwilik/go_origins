package utils

func SkipChars(line string, skipCharsCount int) (resultLine string) {
	if len(line) < skipCharsCount {
		return resultLine
	}
	return line[skipCharsCount:]
}
