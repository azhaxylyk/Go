package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	} else if nb <= 1 {
		return 1
	} else {
		result := 1
		for i := 2; i <= nb; i++ {
			result *= i
			if result <= 0 {
				return 0
			}
		}
		return result
	}
}
