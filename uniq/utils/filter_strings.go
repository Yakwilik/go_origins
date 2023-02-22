package utils

func doInvert(value bool, flag bool) bool {
	switch flag {
	case true:
		{
			return !value
		}
	case false:
		{
			return value
		}
	}
	return value
}
func FilterStringsByMetCount(linesWithMetInfo []StringWithMetCount, uniq bool) (filteredLines []string) {
	for _, lineWithMetInfo := range linesWithMetInfo {
		if doInvert(lineWithMetInfo.MetCount == 1, !uniq) {
			filteredLines = append(filteredLines, lineWithMetInfo.Str)
		}
	}
	return filteredLines
}
