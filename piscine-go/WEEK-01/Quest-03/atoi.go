package piscine

func Atoi(s string) int {
	num := 0
	sign := 1
	firstDigit := true
	for _, c := range s {
		if c == '+' || c == '-' {
			if !firstDigit {
				return 0
			}
			if c == '-' {
				sign = -1
			}
			firstDigit = false
			continue
		}
		if c < '0' || c > '9' {
			return 0
		}
		num = num*10 + int(c-'0')
		firstDigit = false
	}
	return num * sign
}
