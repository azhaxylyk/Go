package piscine

func Abort(a, b, c, d, e int) int {
	values := []int{a, b, c, d, e}

	for i := 0; i < len(values)-1; i++ {
		swapped := false
		for j := 0; j < len(values)-i-1; j++ {
			if values[j] > values[j+1] {
				values[j], values[j+1] = values[j+1], values[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}

	return values[2]
}
