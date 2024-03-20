package main

import (
	"log"
	"os"
	"regexp"
	"strings"
)

func Processing(s string) (string, int) {
	for _, i := range s {
		if (i < 32 && i != 10) || i > 126 {
			log.Fatalln("Bad Request")
		}
	}
	res := ""
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '\\' {
			i++
			if i == len(s) {
				res += "\\"
				return res, count
			} else if s[i] == '\\' {
				res += "\\"
			} else if s[i] == 'n' {
				if i > 0 && s[i-1] == '\\' {
					res += "\\n"
				} else {
					res += "\n"
					count++
				}
			} else {
				res += "\\"
				res += string(s[i])
			}
		} else if s[i] == '\n' {
			count++
			res += string(s[i])
		} else {
			res += string(s[i])
		}
	}
	return res, count
}

func MakeMap(filename string) map[int][]string {
	contents, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalln("Could not read the file")
	}
	text := string(contents)

	symbolRegex := regexp.MustCompile(`([^\r\n]+(\r)*\n){8}`)
	asciiMap := make(map[int][]string)
	asciiArray := symbolRegex.FindAllString(text, 95)
	for i, symbol := range asciiArray {
		symbol = strings.ReplaceAll(symbol, "\r\n", "\n")
		asciiMap[i+32] = strings.Split(symbol, "\n")
	}
	return asciiMap
}

func TextToAscii(asciiArtMap map[int][]string, text []string, count int) string {
	isEmpty := false
	for _, ch := range text {
		if ch != "" {
			isEmpty = true
		}
	}
	if !isEmpty {
		text = text[1:]
	}

	var result string
	for _, word := range text {
		wordLines := [][]string{}
		if len(word) != 0 {
			for _, r := range word {
				wordLines = append(wordLines, asciiArtMap[int(r)])
			}
			for i := 0; i < 8; i++ {
				for j := 0; j < len(word); j++ {
					result += wordLines[j][i]
				}
				result += "\n"
			}
		} else if word == "" && count >= 0 {
			result += "\n"
		}
	}
	return result
}