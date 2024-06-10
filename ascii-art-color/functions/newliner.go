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

func Newlinedetector(lines []string, res string) (string, []string) {
	if len(lines) == 0 {
		return res, lines
	} else {
		for len(lines) > 0 && len(lines[0]) == 0 {
			lines = lines[1:]

			res = res + "\n"
		}
		return res, lines
	}
}
