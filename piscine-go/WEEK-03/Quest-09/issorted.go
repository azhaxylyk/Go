package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	isAscending := true
	isDescending := true

	for i := 1; i < len(a); i++ {
		cmp := f(a[i-1], a[i])
		if cmp > 0 {
			isAscending = false
		} else if cmp < 0 {
			isDescending = false
		}
	}

	return isAscending || isDescending
}

func nCom(a, b int) int {
	if a > b {
		return -1
	} else if a < b {
		return 1
	}
	return 0
}
