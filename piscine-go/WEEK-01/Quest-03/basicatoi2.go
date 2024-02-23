package piscine

func BasicAtoi2(s string) int {
	num := 0
	for _, digit := range s {
		if digit < '0' || digit > '9' {
			return 0 // If non-digit character found, return 0 immediately
		}
		num = num*10 + int(digit-'0')
	}
	return num
}
