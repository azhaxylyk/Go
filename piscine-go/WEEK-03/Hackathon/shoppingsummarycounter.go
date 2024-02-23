package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	item := ""

	for i := 0; i <= len(str); i++ {
		if i == len(str) || str[i] == ' ' || str[i] == '\t' || str[i] == '\n' || str[i] == '\r' {
			summary[item]++
			item = ""
			continue
		}

		item += string(str[i])
	}

	return summary
}
