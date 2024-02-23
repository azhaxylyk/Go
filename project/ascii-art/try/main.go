package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 || len(os.Args[1]) == 0 {
		fmt.Printf("Usage: go run . [STRING] [BANNER]\n\nEX: go run . something standard\n")
		return
	}

	banner := "standard"
	if len(os.Args) == 3 {
		banner = os.Args[2]
	}

	txt, numSlashN := ProcessText(os.Args[1])
	text := strings.Split(txt, "\n")
	asciiArtMap := CreateMap("fonts/" + banner + ".txt")
	res := Join(asciiArtMap, text, numSlashN)
	if CheckTerminal(res) {
		log.Fatalln("Terminal size is too small")
	}
	fmt.Print(res)
}
