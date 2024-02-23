package main

import (
	"bytes"
	"regexp"
	"strconv"
	"strings"
)

func ProcessText(text string) string {
	text = handleHexAndBin(text)
	text = handleComplexCaseConversion(text)
	text = handleAAnConversion(text)
	text = CorrectPunctuation(text)
	text = CorrectQuotes(text)
	text = CorrectSpacing(text)
	return text
}

func handleAAnConversion(text string) string {
	aAnRegex := regexp.MustCompile(`\ba ([aeiouhAEIOUH])`)
	return aAnRegex.ReplaceAllString(text, "an $1")
}

func CorrectPunctuation(text string) string {
	text = regexp.MustCompile(`\s*([.,;:!?])`).ReplaceAllString(text, "$1")
	text = regexp.MustCompile(`([.,;:!?])(\s+)([.,;:!?])`).ReplaceAllString(text, "$1$3")
	text = regexp.MustCompile(`([.,;:!?])\s*([^.,;:!?])`).ReplaceAllString(text, "$1 $2")

	var result []string
	lines := strings.Split(text, "\n")

	for _, line := range lines {
		var buffer bytes.Buffer
		singleQuotesCount, doubleQuotesCount := 0, 0

		for _, r := range line {
			switch r {
			case '\'':
				singleQuotesCount++
				if singleQuotesCount%2 == 0 && buffer.Len() > 0 && buffer.Bytes()[buffer.Len()-1] == ' ' {
					buffer.Truncate(buffer.Len() - 1)
				}
				buffer.WriteRune(r)
			case '"':
				doubleQuotesCount++
				if doubleQuotesCount%2 == 0 && buffer.Len() > 0 && buffer.Bytes()[buffer.Len()-1] == ' ' {
					buffer.Truncate(buffer.Len() - 1)
				}
				buffer.WriteRune(r)
			default:
				buffer.WriteRune(r)
			}
		}

		result = append(result, buffer.String())
	}

	return strings.Join(result, "\n")
}

func CorrectQuotes(text string) string {
	apostrophePattern := regexp.MustCompile(`(\w)\s*` + "`" + `\s*(\w)`)
	text = apostrophePattern.ReplaceAllString(text, "$1`$2")

	text = regexp.MustCompile(`(')(\S)`).ReplaceAllString(text, "$1 $2")
	text = ProcessPairedPunctuation(text, `'`)
	text = ProcessPairedPunctuation(text, `"`)

	return text
}

func ProcessPairedPunctuation(text string, punctuation string) string {
	pattern := regexp.QuoteMeta(punctuation) + `([\s\S]*?)` + regexp.QuoteMeta(punctuation)
	re := regexp.MustCompile(pattern)

	text = re.ReplaceAllStringFunc(text, func(match string) string {
		content := strings.TrimSpace(re.FindStringSubmatch(match)[1])
		if content == "" {
			return punctuation + punctuation
		}
		content = regexp.MustCompile(`\s+`).ReplaceAllString(content, " ")
		return punctuation + content + punctuation
	})

	text = regexp.MustCompile(`(`+regexp.QuoteMeta(punctuation)+`)\s*(`+regexp.QuoteMeta(punctuation)+`)`).ReplaceAllString(text, "$1 $2")

	return text
}

func CorrectSpacing(text string) string {
	var result []string
	lines := strings.Split(text, "\n")

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine != "" {
			trimmedLine = regexp.MustCompile(`\s+`).ReplaceAllString(trimmedLine, " ")
			result = append(result, trimmedLine)
		}
	}

	return strings.Join(result, "\n")
}

func handleHexAndBin(text string) string {
	hexRegex := regexp.MustCompile(`\b([0-9a-fA-F]+) \(hex\)`)
	text = hexRegex.ReplaceAllStringFunc(text, func(match string) string {
		hexMatch := hexRegex.FindStringSubmatch(match)[1]
		decimalNum, err := strconv.ParseInt(hexMatch, 16, 64)
		if err != nil {
			return match
		}
		return strconv.FormatInt(decimalNum, 10)
	})
	binRegex := regexp.MustCompile(`\b([01]+) \(bin\)`)
	text = binRegex.ReplaceAllStringFunc(text, func(match string) string {
		binMatch := binRegex.FindStringSubmatch(match)[1]
		decimalNum, err := strconv.ParseInt(binMatch, 2, 64)
		if err != nil {
			return match
		}
		return strconv.FormatInt(decimalNum, 10)
	})
	return text
}

func handleComplexCaseConversion(text string) string {
	regex := regexp.MustCompile(`\((cap|up|low)(, (\d+))?\)`)
	matches := regex.FindAllStringSubmatchIndex(text, -1)
	for i := len(matches) - 1; i >= 0; i-- {
		match := matches[i]
		operation := text[match[2]:match[3]]
		n := 1
		if match[4] != -1 {
			num, err := strconv.Atoi(text[match[6]:match[7]])
			if err == nil {
				n = num
			}
		}
		words := strings.Fields(text[:match[0]])
		if len(words) < n {
			continue
		}
		startIndex := max(0, len(words)-n)
		for j := startIndex; j < len(words); j++ {
			word := words[j]
			switch operation {
			case "up":
				words[j] = strings.ToUpper(word)
			case "low":
				words[j] = strings.ToLower(word)
			case "cap":
				words[j] = strings.Title(strings.ToLower(word))
			}
		}
		beforeMatch := strings.Join(words[:startIndex], " ")
		afterMatch := text[match[1]:]
		if startIndex > 0 {
			beforeMatch += " "
		}
		text = beforeMatch + strings.Join(words[startIndex:], " ") + afterMatch
	}
	return strings.TrimSpace(text)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
