package piscine

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	}

	maxValue := a[0]
	for _, num := range a {
		if num > maxValue {
			maxValue = num
		}
	}

	return maxValue
}
