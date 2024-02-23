package piscine

func SortIntegerTable(table []int) {
	length := len(table)
	for i := 0; i < length-1; i++ {
		for j := 0; j < length-i-1; j++ {
			if table[j] > table[j+1] {
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}
}
