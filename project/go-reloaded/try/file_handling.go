package main

import (
	"io/ioutil"
)

func ReadFile(fileName string) (string, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func WriteFile(fileName, text string) error {
	return ioutil.WriteFile(fileName, []byte(text), 0744)
}
