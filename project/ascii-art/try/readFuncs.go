package main

import (
	"log"
	"os"
	"regexp"
	"strings"
)

func ProcessText(s string) (string, int) {
	for _, ch := range s {
		if (ch < 32 && ch != 10) || ch > 126 {
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
				res += "\n"
				count++
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

func CreateMap(filename string) map[int][]string {
	contents, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}
	text := string(contents)

	symbolRegex := regexp.MustCompile(`([^\r\n]+(\r)*\n){8}`)
	asciiArtMap := make(map[int][]string)
	asciiArtArr := symbolRegex.FindAllString(text, 95)
	for i, symbol := range asciiArtArr {
		symbol = strings.ReplaceAll(symbol, "\r\n", "\n")
		asciiArtMap[i+32] = strings.Split(symbol, "\n")
	}
	return asciiArtMap
}
