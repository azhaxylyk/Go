package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	charValue := func(char byte) int {
		if char >= '0' && char <= '9' {
			return int(char - '0')
		} else if char >= 'A' && char <= 'Z' {
			return int(char-'A') + 10
		}
		return -1
	}

	base10Num := 0
	for i := len(nbr) - 1; i >= 0; i-- {
		digitValue := charValue(nbr[i])
		if digitValue == -1 || digitValue >= len(baseFrom) {
			return "Invalid input"
		}
		base10Num += digitValue * pow(len(baseFrom), len(nbr)-1-i)
	}

	baseToNum := ""
	for base10Num > 0 {
		remainder := base10Num % len(baseTo)
		baseToNum = string(baseTo[remainder]) + baseToNum
		base10Num /= len(baseTo)
	}

	return baseToNum
}

func pow(base, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}
