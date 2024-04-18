package functions

import "strings"

func SingleLine(s string) string {
	s = strings.ReplaceAll(s, string(rune(10)), "\\n")
	result := ""
	escape := false
	for _, char := range s {
		if escape {
			switch char {
			case '\\':
				result += "\\"
			case 'n':
				result += "\n"
			default:
				result += "\\" + string(char)
			}
			escape = false
		} else if char == '\\' {
			escape = true
		} else {
			result += string(char)
		}
	}
	return result
}
