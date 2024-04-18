package functions

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func Manage() error {
	if len(os.Args) == 1 {
		return errors.New("Plese write minimum 1 Argument")
	}

	if strings.HasPrefix(os.Args[1], "--reverse=") {
		fileName := strings.TrimPrefix(os.Args[1], "--reverse=")

		err := ReverseAsciiArt(fileName)
		if err != nil {
			return err
		}
		return nil
	}

	var (
		text  string
		style string = "standard"
	)

	if len(os.Args) == 2 {
		text = os.Args[1]
	} else if len(os.Args) == 3 {
		text = os.Args[1]
		style = os.Args[2]
	}
	err := CalculateFileHash(style)
	if err != nil {
		return err
	}

	res, err := Art(text, style)
	if err != nil {
		return err
	}
	fmt.Print(res)
	return nil
}

var (
	originalStandard   []string
	originalShadow     []string
	originalThinkertoy []string
)

func ReverseAsciiArt(fileName string) error {
	lines, err := FileReader(fileName)
	if err != nil {
		return err
	}
	originalStandard, err = FileReader("banner/standard.txt")
	if err != nil {
		return err
	}
	originalShadow, err = FileReader("banner/shadow.txt")
	if err != nil {
		return err
	}
	originalThinkertoy, err = FileReader("banner/thinkertoy.txt")
	if err != nil {
		return err
	}
	temp := make([]string, 8)
	res := ""

	res, lines, temp = Reverse(res, lines, temp)
	res = strings.TrimRight(res, "\n")
	fmt.Println(res)

	return nil
}

func Reverse(res string, lines, temp []string) (string, []string, []string) {
	for i := 0; len(lines[0]) > 0; i++ {

		err := ErrCasesCheker(lines)
		if err != nil {
			log.Fatal(err)
		}

		res, lines = Newlinedetector(lines, res)

		temp, lines = FromGraphicaltoTemp(temp, lines)

		a, finded := Sravnitel(temp)
		if finded == true {
			res = res + string(a)
		}

	}

	res, lines = Newlinedetector(lines, res)

	if len(lines) != 0 {
		return Reverse(res, lines, temp)
	}

	return res, lines, temp
}

func Sravnitel(temp []string) (rune, bool) {
	for i := 32; i < 127; i++ {

		orginalLeteer := getOriginalLetter(i, originalStandard)

		finded := Equalizer(temp, orginalLeteer)

		if finded {
			for i := 0; i < len(temp); i++ {
				temp[i] = ""
			}
			return rune(i), true
		}

	}

	return 0, false
}

func Equalizer(temp, orginalLeteer []string) bool {
	var finded bool = true

	for i := 0; i < len(temp); i++ {
		if temp[i] != orginalLeteer[i] {

			finded = false

			break

		}
	}
	return finded
}

func getOriginalLetter(i int, style []string) []string {
	var res []string

	simsnum := i - 32

	for j := 0; j < 8; j++ {
		res = append(res, style[(simsnum+1)+(simsnum*8)+j])
	}

	return res
}

func FromGraphicaltoTemp(temp, lines []string) ([]string, []string) {
	for i := 0; i < len(temp); i++ {

		temp[i] = temp[i] + string(lines[i][0])

		lines[i] = lines[i][1:]

	}

	isEmpety := true

	for i := 0; i < 8; i++ {
		if len(lines[i]) != 0 {
			isEmpety = false
			break
		}
	}

	if isEmpety {
		lines = lines[7:]
	}

	return temp, lines
}

func FileReader(fileName string) ([]string, error) {
	err := CalculateFileHash("standard")
	if err != nil {
		return make([]string, 0), err
	}
	err = CalculateFileHash("shadow")
	if err != nil {
		return make([]string, 0), err
	}
	err = CalculateFileHash("thinkertoy")
	if err != nil {
		return make([]string, 0), err
	}

	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		return make([]string, 0), err
	}

	fileString := string(fileContent)

	fileString = strings.ReplaceAll(fileString, "\r", "")

	spliter := strings.Split(fileString, "\n")

	return spliter, nil
}

func WriteToFile(fileName, text, style string) error {
	asciiArt, err := Art(text, style)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(fileName, []byte(asciiArt), 0644)
	if err != nil {
		return err
	}

	return nil
}

func PrintAsciiArt(text, style string) error {
	asciiArt, err := Art(text, style)
	if err != nil {
		return err
	}
	err = CalculateFileHash(style)

	if err != nil {
		return err
	}
	fmt.Print(asciiArt)
	return nil
}

func Art(text, style string) (string, error) {
	err := AsciiCheeker(text)
	if err != nil {
		return "", err
	}

	text = SingleLine(text)
	textArr := strings.Split(text, "\n")

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

	for i := 0; i < len(textarr); i++ {
		lines, err := PrintAscii(style)
		if err != nil {
			return "", err
		}

		text := textarr[i]
		if text == "" {
			res = res + "\n"
			continue
		}

		var slice [8]string

		for i := 0; i < len(text); i++ {
			simsNum := int(text[i] - 32)

			for j := 0; j < len(slice); j++ {
				slice[j] = slice[j] + lines[(simsNum+1)+(simsNum*8)+j]
			}
		}

		var tempResString string
		for i := 0; i < len(slice); i++ {
			tempResString = tempResString + slice[i] + "\n"
		}
		res = res + tempResString //

	}
	return res, nil
}

func PrintAscii(s string) ([]string, error) {
	s = "banner/" + s + ".txt"
	file, err := os.ReadFile(s)
	if err != nil {
		return nil, err
	}
	text := string(file)
	if s == "banner/thinkertoy.txt" {
		text = strings.ReplaceAll(text, "\r", "")
	}
	mass := strings.Split(text, "\n")
	return mass, nil
}
