package handlers

func processLine(line string, skipWordsCount, skipCharsCount int) string {
	return skipChars(skipWords(line, skipCharsCount), skipWordsCount)
}
