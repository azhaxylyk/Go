package piscine

func StringToIntSlice(str string) []int {
	if str == "" {
		return nil
	}

	result := []int{}

	for _, char := range str {
		result = append(result, int(char))
	}

	return result
}
