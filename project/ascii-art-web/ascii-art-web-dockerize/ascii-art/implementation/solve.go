package implementation

import (
	"Go/project/ascii-art-web/ascii-art-web-dockerize/ascii-art/container"
	"errors"
)

func Operate(start_index, end_index int, Format [][]string) string {
	var object string
	for i := 0; i < 8; i++ {
		for j := start_index; j < end_index; j++ {
			object += string(Format[container.Global_sentence[j]-32][i])
		}
		object += "\n"
	}
	return object
}

func Solve(banner string) (string, error) {
	var Format [][]string
	if banner == "standard" {
		Format = container.Format_standard
	} else if banner == "shadow" {
		Format = container.Format_shadow
	} else if banner == "thinkertoy" {
		Format = container.Format_thinkertoy
	} else {
		return "", errors.New("Not Found")
	}

	var outputting_string string = ""

	for i := 0; i < len(container.Global_sentence); i++ {
		if !((container.Global_sentence[i] >= 32 && container.Global_sentence[i] <= 126) || container.Global_sentence[i] == 10 || container.Global_sentence[i] == 13) {
			return "", errors.New("Bad Request")
		}
	}

	var start, end int = 0, 0
	for i := 0; i < len(container.Global_sentence); i++ {
		if container.Global_sentence[i] == '\n' {
			end = i - 1
			outputting_string += Operate(start, end, Format)
			start = i + 1
		}
	}
	outputting_string += Operate(start, len(container.Global_sentence), Format)

	return outputting_string, nil
}
