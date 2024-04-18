package ascii

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
	"regexp"
)

var hashMap = map[string]string{
	"ascii/banners/standard.txt":   "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf",
	"ascii/banners/shadow.txt":     "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73",
	"ascii/banners/thinkertoy.txt": "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3",
}

var (
	newlineRegex = regexp.MustCompile(`\\n`)
	quoteRegex   = regexp.MustCompile(`\\"`)
)

func HashChecker(font string) error {
	fontPath := "ascii/banners/" + font
	data, err := os.ReadFile(fontPath)
	if err != nil {
		return err
	}

	generatedHash := fmt.Sprintf("%x", sha256.Sum256(data))
	expectedHash, ok := hashMap[fontPath]
	if !ok {
		return fmt.Errorf("hash not found for font: %s", font)
	}

	if expectedHash != generatedHash {
		return fmt.Errorf("invalid font hash for %s", font)
	}

	return nil
}

func Convert(input, fontName string) (string, error) {
	err := HashChecker(fontName)
	if err != nil {
		return "", err
	}

	text := NewLineBreaker(input)

	font, err := GetFont("ascii/banners/" + fontName)
	if err != nil {
		return "", err
	}

	result := AsciiArt(text, font)
	return result, nil
}

func GetFont(fontName string) (map[rune][]string, error) {
	file, err := os.Open(fontName)
	if err != nil {
		return map[rune][]string{}, err
	}

	count := 0
	r := ' '
	font := make(map[rune][]string, 8)

	scanner := bufio.NewScanner(file)
	for ; scanner.Scan(); count++ {
		text := scanner.Text()
		if count != 0 && count != 9 {
			font[r] = append(font[r], text)
		} else if count == 9 {
			count = 0
			r++
		}
	}
	return font, nil
}

func AsciiArt(text []string, font map[rune][]string) string {
	result := ""
	output := make([][8]string, len(text))
	for k, val := range text {
		if val == "" {
			output[k] = [8]string{"\n"}
			continue
		}
		for _, r := range val {
			for i, g := range font[r] {
				output[k][i] += g
			}
		}
	}

	for _, h := range output {
		for _, k := range h {
			if k == "\n" {
				result += "\n"
				break
			}
			result += k + "\n"
		}
	}

	return result
}

func NewLineBreaker(input string) []string {
	if len(input) == 0 {
		return nil
	}

	input = newlineRegex.ReplaceAllString(input, "\n")
	input = quoteRegex.ReplaceAllString(input, "\"")

	return splitInput(input)
}

func splitInput(input string) []string {
	var lines []string
	var line string

	for _, r := range input {
		if r == '\n' {
			lines = append(lines, line)
			line = ""
		} else {
			line += string(r)
		}
	}
	if line != "" {
		lines = append(lines, line)
	}

	return lines
}
