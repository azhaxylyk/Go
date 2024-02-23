package piscine

func BasicAtoi(s string) int {
	num := 0
	for _, digit := range s {
		if digit < '0' || digit > '9' {
			continue
		}
		num = num*10 + int(digit-'0')
	}
	return num
}
