package utils

func SkipWordsAndChars(line string, skipWordsCount, skipCharsCount int) string {
	return SkipChars(SkipWords(line, skipWordsCount), skipCharsCount)
}
