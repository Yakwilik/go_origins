package utils

func FilterStringsByMetCount(linesWithMetInfo []StringWithMetCount, duplicates bool) (filteredLines []string) {
	for _, lineWithMetInfo := range linesWithMetInfo {
		if !duplicates {
			filteredLines = append(filteredLines, lineWithMetInfo.Str)
		} else if lineWithMetInfo.MetCount > 1 {
			filteredLines = append(filteredLines, lineWithMetInfo.Str)
		}

	}
	return filteredLines
}

func FilterUniqueStrings(linesWithMetInfo []StringWithMetCount) (filteredLines []string) {
	stringMetCount := make(map[string]int, 0)
	for _, stringWithMetCount := range linesWithMetInfo {
		if stringWithMetCount.MetCount == 1 {
			stringMetCount[stringWithMetCount.Str]++
		}
	}
	for _, stringWithMetCount := range linesWithMetInfo {
		if stringMetCount[stringWithMetCount.Str] == 1 {
			filteredLines = append(filteredLines, stringWithMetCount.Str)
		}
	}
	return filteredLines
}
