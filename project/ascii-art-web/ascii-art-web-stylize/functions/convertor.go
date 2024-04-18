package functions

import (
	"fmt"
	"os"
	"strings"
)

func Manage(input string, banner string) (string, error) {
	err := CalculateFileHash(banner)
	if err != nil {
		return "", err
	}

	res, err := Art(input, banner)
	if err != nil {
		return "", err
	}

	return res, nil
}

func Art(text, style string) (string, error) {
	var textArr []string
	if strings.Contains(text, "\\n") {
		textArr = strings.Split(text, "\\n")
	} else {
		textArr = strings.Split(text, "\n")
	}

	fmt.Println(textArr)

	res, err := AsciiReturner(textArr, style)
	if err != nil {
		return "", err
	}

	return res, nil
}

func AsciiReturner(textarr []string, style string) (string, error) {
	if len(textarr) == 0 {
		return "", nil
	}

	linesIsEmpety, err := EmpetyTextsCheker(textarr)
	if err != nil {
		return "", err
	}

	if linesIsEmpety {
		var res string
		for i := 0; i < len(textarr)-1; i++ {
			res = res + "\n"
		}
		return res, nil
	}

	var res string

	lines, err := PrintAscii(style)
	if err != nil {
		return "", err
	}

	for i := 0; i < len(textarr); i++ {
		text := textarr[i]
		if text == "" {
			res = res + "\n"
			continue
		}

		var slice [8]string

		for j := 0; j < len(text); j++ {
			simsNum := int(text[j] - 32)
			if simsNum >= 0 && (simsNum+1)+(simsNum*8) < len(lines) {
				for k := 0; k < len(slice); k++ {
					slice[k] = slice[k] + lines[(simsNum+1)+(simsNum*8)+k]
				}
			} else {
				for k := 0; k < len(slice); k++ {
					slice[k] = slice[k] + " "
				}
			}
		}

		var tempResString string
		for j := 0; j < len(slice); j++ {
			tempResString = tempResString + slice[j] + "\n"
		}
		res = res + tempResString
	}
	return res, nil
}

func PrintAscii(s string) ([]string, error) {
	s = s + ".txt"
	file, err := os.ReadFile(s)
	if err != nil {
		return nil, err
	}
	text := string(file)
	if s == "thinkertoy.txt" {
		text = strings.ReplaceAll(text, "\r", "")
	}
	mass := strings.Split(text, "\n")
	return mass, nil
}
