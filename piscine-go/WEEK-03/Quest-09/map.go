package piscine

func Map(f func(int) bool, a []int) []bool {
	result := make([]bool, len(a))
	for i, val := range a {
		result[i] = f(val)
	}
	return result
}
