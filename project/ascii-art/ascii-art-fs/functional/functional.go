package functional

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type (
	BannerMap   map[rune][]string
	BannersHash map[string]string
)

var banners BannersHash = BannersHash{
	"standard.txt":   "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf",
	"shadow.txt":     "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73",
	"thinkertoy.txt": "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3",
}

func IsAscii32_126(s string) bool {
	for _, char := range s {
		if char < 32 && char != '\n' || char > 126 {
			return false
		}
	}
	return true
}

func AsciiArt(input, filename string) (string, error) {
	err := hashChecker(filename)
	if err != nil {
		return "", err
	}
	result, err := stylization(input, filename)
	if err != nil {
		return "", err
	}
	return result, nil
}

func ReadBannerFile(fileName string) (BannerMap, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, errors.New("Error opening file: " + err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	banner := BannerMap{}
	for i, j := 0, 31; scanner.Scan(); i++ {
		if i%9 == 0 {
			j++
		}
		key := rune(j)
		banner[key] = append(banner[key], scanner.Text())
	}
	if scanner.Err() != nil {
		return nil, errors.New("Error reading file:" + err.Error())
	}
	return banner, nil
}

func stylization(inputText, fileName string) (string, error) {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println("Error: Program can't show current directory")
		return "", err
	}
	bannerPath := pathCorrector(path) + "banners/" + fileName
	bannerMap, err := ReadBannerFile(bannerPath)
	if err != nil {
		fmt.Println("Error opening file")
		return "", err
	}
	result := ""
	text := strings.Replace(inputText, "\\n", "\n", -1)
	splittedText := strings.Split(text, "\n")
	if !isOnlyNewLine(splittedText) {
		splittedText = splittedText[1:]
	}

	for _, line := range splittedText {
		if line == "" {
			result += "\n"
			continue
		}
		for i := 1; i <= 8; i++ {
			for _, c := range line {
				result += bannerMap[c][i]
			}
			result += "\n"
		}
	}
	return result, nil
}

func isOnlyNewLine(splitted []string) bool {
	for _, part := range splitted {
		if part != "" {
			return true
		}
	}
	return false
}

func hashChecker(fileName string) error {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println("Error: Program can't show current directory")
		return err
	}
	bannerPath := pathCorrector(path) + "banners/" + fileName
	data, err := ioutil.ReadFile(bannerPath)
	if err != nil {
		return fmt.Errorf("There is no:" + fileName + "banner")
	}
	hasher := sha256.New()
	hasher.Write(data)
	hash := hasher.Sum(nil)
	hashString := hex.EncodeToString(hash)
	if hashString != banners[fileName] {
		return fmt.Errorf("Error:" + fileName + " banner was changed")
	}
	return nil
}

func pathCorrector(path string) string {
	return strings.Split(path, "fs")[0] + "fs/"
}
