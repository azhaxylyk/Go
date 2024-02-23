package piscine

func Sqrt(nb int) int {
	if nb == 1 {
		return 1
	}

	for i := 1; i <= nb/2; i++ {
		if i*i == nb {
			return i
		}
	}

	return 0
}
