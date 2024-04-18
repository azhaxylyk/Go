package functional

import (
	"log"
	"os"
	"strings"
	"testing"
)

func TestAsciiArt(t *testing.T) {
	testCases := []struct {
		input    string
		banner   string
		expected string
	}{
		{
			input:    "hello",
			banner:   "standard",
			expected: "test01.txt",
		},
		{
			input:    "hello world",
			banner:   "shadow",
			expected: "test02.txt",
		},
		{
			input:    "nice 2 meet you",
			banner:   "thinkertoy",
			expected: "test03.txt",
		},
		{
			input:    "you & me",
			banner:   "standard",
			expected: "test04.txt",
		},
		{
			input:    "123",
			banner:   "shadow",
			expected: "test05.txt",
		},
		{
			input:    "/(\")",
			banner:   "thinkertoy",
			expected: "test06.txt",
		},
		{
			input:    "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			banner:   "shadow",
			expected: "test07.txt",
		},
		{
			input:    "\"#$%&/()*+,-./",
			banner:   "thinkertoy",
			expected: "test08.txt",
		},
		{
			input:    "It's Working",
			banner:   "thinkertoy",
			expected: "test09.txt",
		},
	}

	for _, testCase := range testCases {
		got, err := AsciiArt(testCase.input, testCase.banner+".txt")
		if err != nil {
			t.Fatalf("Error: Program can't stylize current text")
		}
		path, _ := os.Getwd()
		expected, err := os.ReadFile(pathCorrector(path) + "/testcases/" + testCase.expected)
		if err != nil {
			log.Fatal(err)
		}

		gotNormalized := strings.Replace(strings.TrimSpace(got), "\r\n", "\n", -1)
		expectedNormalized := strings.Replace(strings.TrimSpace(string(expected)), "\r\n", "\n", -1)

		if gotNormalized != expectedNormalized {
			t.Fatalf("got:\n%s\nwant\n%s\n", gotNormalized, string(expectedNormalized))
		}
	}
}
