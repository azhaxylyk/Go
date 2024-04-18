package main

import (
	"ascii-art/functional"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) > 2 {
		fmt.Println(
			`Usage: go run . [STRING] [BANNER]

EX: go run . something standard`)
		return
	}
	StdInString := args[0]
	if len(StdInString) == 0 {
		fmt.Println("Error: empty input")
		return
	}
	if !functional.IsAscii32_126(StdInString) {
		fmt.Println("Error: Program works with only Ascii symbols with indexes 32 - 126")
		return
	}

	banneName := "standard"
	if len(args) == 2 {
		banneName = args[1]
	}

	if !isBanner(banneName) {
		fmt.Println("Wrong bannerName")
		return
	}
	result, err := functional.AsciiArt(StdInString, banneName+".txt")
	if err != nil {
		fmt.Println("Error: Program can't stylize current text")
	}
	fmt.Print(result)
}

func isBanner(bannerName string) bool {
	bannerSlice := []string{"standard", "shadow", "thinkertoy"}
	for _, r := range bannerSlice {
		if r == bannerName {
			return true
		}
	}
	return false
}
