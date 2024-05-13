package convert

import (
	"adel/ascii-art/check"
	"adel/ascii-art/container"
	"fmt"
	"os"
)

func Convert(content string, number int) {
	var array_sentence [][]string = make([][]string, 126-31)
	var object string
	var flag bool = false
	var count_row int = 0
	for i := 0; i < len(content); i++ {
		if content[i] >= 32 && content[i] <= 126 {
			flag = true
			object += string(content[i])
		} else if flag == true {
			flag = false
			array_sentence[count_row] = append(array_sentence[count_row], object)
			object = ""
			if i+1 == len(content) || content[i+1] == '\n' {
				count_row++
				i++
			}
		}
	}

	if number == 0 {
		container.Format_standard = array_sentence
	} else if number == 1 {
		container.Format_shadow = array_sentence
	} else if number == 2 {
		container.Format_thinkertoy = array_sentence
	}
}

func Extract_From_Banner(file_name string) (string, error) { // Getting a line which contains the content of a particular banner text

	file_name = "../ascii-art/banner/" + file_name
	file, err := os.Open(file_name)
	if err != nil {
		fmt.Print("Function name: Extract_From_Banner ----> ")
		fmt.Println("Error opening file:", err)
		return "", err
	}
	defer file.Close()

	fileInfo, err := os.Stat(file_name)
	if err != nil {
		fmt.Print("Function name: Extract_From_Banner ---->")
		fmt.Println("Error getting file info:", err)
		return "", err
	}

	data := make([]byte, fileInfo.Size())
	n, err := file.Read(data)
	if err != nil {
		fmt.Print("Function name: Extract_From_Banner ---->")
		fmt.Println("Error reading file:", err)
		return "", err
	}

	content := string(data[:n])

	if file_name == "../ascii-art/banner/thinkertoy.txt" {
		content = Change(content)
	}

	return content, nil
}

func Initialize_Format() error { // Extracts the elements from banner text, and it saves into array of string.

	if err := check.Comparing_hash("standard.txt"); err != nil {
		fmt.Println("Function name: Initialize_Format")
		return err
	}

	content_standard, err := Extract_From_Banner("standard.txt")
	if err != nil {
		fmt.Print("Function name: Convert")
		return err
	}
	Convert(content_standard, 0)

	if err := check.Comparing_hash("shadow.txt"); err != nil {
		fmt.Println("Function name: Initialize_Format")
		return err
	}
	content_shadow, err := Extract_From_Banner("shadow.txt")
	if err != nil {
		fmt.Print("Function name: Convert")
		return err
	}
	Convert(content_shadow, 1)

	if err := check.Comparing_hash("thinkertoy.txt"); err != nil {
		fmt.Println("Function name: Initialize_Format")
		return err
	}
	content_thinkertoy, err := Extract_From_Banner("thinkertoy.txt")
	if err != nil {
		fmt.Print("Function name: Convert")
		return err
	}
	Convert(content_thinkertoy, 2)

	return nil
}
