package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 2 {
		if len(os.Args[1]) == 0 {
			return
		}

		text, num := Processing(os.Args[1])
		text = strings.Replace(text, "\\n", "\n", -1)
		lines := strings.Split(text, "\n")
		asciiArtMap := MakeMap("standard.txt")
		result := TextToAscii(asciiArtMap, lines, num)

		fmt.Print(result)
		return
	}
}
