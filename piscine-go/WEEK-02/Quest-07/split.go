package piscine

func Split(s, sep string) []string {
	var result []string
	start := 0

	for i := 0; i <= len(s)-len(sep); i++ {
		if s[i:i+len(sep)] == sep {
			result = append(result, s[start:i])
			i += len(sep) - 1
			start = i + 1
		}
	}

	if start < len(s) {
		result = append(result, s[start:])
	}

	return result
}
