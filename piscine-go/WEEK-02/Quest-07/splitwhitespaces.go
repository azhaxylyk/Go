package piscine

func isWhiteSpace(char byte) bool {
	return char == ' ' || char == '\t' || char == '\n'
}

func SplitWhiteSpaces(s string) []string {
	var result []string
	var currentWord string

	for i := 0; i < len(s); i++ {
		if isWhiteSpace(s[i]) {
			if currentWord != "" {
				result = append(result, currentWord)
				currentWord = ""
			}
		} else {
			currentWord += string(s[i])
		}
	}

	if currentWord != "" {
		result = append(result, currentWord)
	}

	return result
}
