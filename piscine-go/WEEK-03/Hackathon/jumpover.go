package piscine

func JumpOver(str string) string {
	if len(str) == 0 {
		return "\n"
	}

	var output string
	for i := 2; i < len(str); i += 3 {
		output += string(str[i])
	}

	return output + "\n"
}
