package main

func Join(asciiArtMap map[int][]string, text []string, count int) string {
	res := ""
	for _, word := range text {
		arr := [][]string{}
		if len(word) != 0 {
			for _, ch := range word {
				arr = append(arr, asciiArtMap[int(ch)])
			}
			for r := 0; r < 8; r++ {
				for i := 0; i < len(word); i++ {
					res += arr[i][r]
				}
				res += "\n"
			}
		} else if word == "" && count > 0 {
			res += "\n"
			count--
		}
	}
	return res
}
