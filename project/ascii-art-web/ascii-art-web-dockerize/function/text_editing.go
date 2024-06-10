package function

import (
	"Go/project/ascii-art-web/ascii-art-web-dockerize/ascii-art/container"
	"Go/project/ascii-art-web/ascii-art-web-dockerize/ascii-art/implementation"
)

func generateAsciiArt(text, banner string) (string, error) {
	container.Global_sentence = text
	result, err := implementation.Solve(banner)
	if err != nil {
		return "", err
	}
	return result, nil
}
