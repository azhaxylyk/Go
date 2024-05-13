package function

import (
	"adel/ascii-art/container"
	"adel/ascii-art/implementation"
)

func generateAsciiArt(text, banner string) (string, error) {
	container.Global_sentence = text
	result, err := implementation.Solve(banner)
	if err != nil {
		return "", err
	}
	return result, nil
}
