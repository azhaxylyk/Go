package piscine

func ShoppingListSort(slice []string) []string {
	var minLength int
	var minIndex int

	for i := range slice {
		if len(slice[i]) < minLength {
			minLength = len(slice[i])
			minIndex = i
		}
	}

	for i := len(slice) - 1; i >= 0; i-- {
		if len(slice[i]) == minLength {
			slice[i], slice[minIndex] = slice[minIndex], slice[i]
			minIndex = i
		}
	}

	return slice
}
