package main

import (
	"bytes"
	"errors"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func ProccessText(text string) string {
	// ProccessText takes a string of text and applies several fixing functions to it, line by line.
	var result []string
	lines := strings.Split(text, "\n")

	for _, line := range lines {
		modifiedLine, err := ProcessTextModifications(line)
		if err != nil {
			log.Printf("Error applying functions to line: %v", err)
		}
		modifiedLine = CorrectQuotes(modifiedLine)
		modifiedLine = CorrectPunctuation(modifiedLine)
		modifiedLine = CorrectArticles(modifiedLine)
		modifiedLine = CorrectSpacing(modifiedLine)
		result = append(result, modifiedLine)
	}

	return strings.Join(result, "\n")
}

func ProcessTextModifications(text string) (string, error) {
	// ProcessTextModifications dynamically searches for and applies modifications based on special patterns found in the text.
	for {
		re := regexp.MustCompile(`\(\s*([Uu][pP]|[Ll][Oo][Ww]|[Cc][Aa][Pp]|[Hh][Ee][Xx]|[Bb][Ii][Nn])\s*(?:,)?\s*((-?\d+)\s*)?\)`)
		subMatches := re.FindAllStringSubmatch(text, -1)
		if len(subMatches) == 0 {
			break
		}

		for _, match := range subMatches {
			modType := strings.ToLower(match[1])
			numStr := strings.TrimRight(match[2], " ")

			if numStr == "" {
				numStr = "1"
			}
			num, err := strconv.Atoi(numStr)
			if err != nil {

				return text, err
			}
			if num < 0 {
				log.Printf("Error: Negative number (%d) found for modification type '%s'", num, modType)
				return text, errors.New("Negative number not allowed for modification")
			}

			switch modType {
			case "hex":
				if _, err := strconv.ParseInt(numStr, 16, 64); err != nil {
					log.Printf("Error: Invalid hexadecimal number (%s) found for (hex) function", numStr)
					return text, errors.New("Invalid hexadecimal number")
				}
			case "bin":
				if _, err := strconv.ParseInt(numStr, 2, 64); err != nil {
					log.Printf("Error: Invalid binary number (%s) found for (bin) function", numStr)
					return text, errors.New("Invalid binary number")
				}
			}

			splitted := re.Split(text, 2)
			textToBeChanged := splitted[0]
			runesToBeChanged := []rune(textToBeChanged)
			textToBeChanged = ToBeModified(runesToBeChanged, num, modType)
			splitted[0] = textToBeChanged
			if len(splitted) > 1 && len(strings.TrimSpace(splitted[1])) > 0 {
				splitted[0] += " "
			}
			text = strings.Join(splitted, "")
		}
	}

	return text, nil
}

func CorrectArticles(text string) string {
	// CorrectArticles corrects the usage of articles ('a' and 'an') throughout the text.
	var result []string
	lines := strings.Split(text, "\n")

	for _, line := range lines {
		line = CorrectArticlesPerLine(line)
		result = append(result, line)
	}

	return strings.Join(result, "\n")
}

func CorrectArticlesPerLine(line string) string {
	// CorrectArticlesPerLine applies corrections for the use of 'a' and 'an' within a single line.
	line = HandleArticleSequences(line)

	reLowerA := regexp.MustCompile(`\b(a)\s+([aeiouhAEIOUH])`)
	reUpperA := regexp.MustCompile(`\b(A|AN|aN)\s+([aeiouhAEIOUH])`)
	reLowerAn := regexp.MustCompile(`\b(an)\s+([bcdfgjklmnpqrstvwxyzBCDFGJKLMNPQRSTVWXYZ])`)
	reUpperAn := regexp.MustCompile(`\b(An|AN|aN)\s+([bcdfgjklmnpqrstvwxyzBCDFGJKLMNPQRSTVWXYZ])`)

	line = reLowerA.ReplaceAllString(line, "an $2")
	line = reUpperA.ReplaceAllString(line, "An $2")
	line = reLowerAn.ReplaceAllString(line, "a $2")
	line = reUpperAn.ReplaceAllString(line, "A $2")

	return line
}

func CorrectPunctuation(text string) string {
	// CorrectPunctuation corrects punctuation spacing and handling throughout the text.
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
	// CorrectQuotes corrects the usage and spacing of quotes within the text.
	apostrophePattern := regexp.MustCompile(`(\w)\s*` + "`" + `\s*(\w)`)
	text = apostrophePattern.ReplaceAllString(text, "$1`$2")

	text = regexp.MustCompile(`(')(\S)`).ReplaceAllString(text, "$1 $2")
	text = ProcessPairedPunctuation(text, `'`)
	text = ProcessPairedPunctuation(text, `"`)

	return text
}

func ProcessPairedPunctuation(text string, punctuation string) string {
	// ProcessPairedPunctuation corrects spacing for paired punctuation marks (quotes) within the text.
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
	// CorrectSpacing corrects excessive or insufficient spacing within the text.
	var result []string
	lines := strings.Split(text, "\n")

	for _, line := range lines {
		trimmedLine := strings.TrimSpace(line)
		if trimmedLine != "" {
			trimmedLine = regexp.MustCompile(`\s+`).ReplaceAllString(trimmedLine, " ")
			result = append(result, trimmedLine) // Adds the corrected line to the result.
		}
	}

	return strings.Join(result, "\n")
}

func HandleArticleSequences(line string) string {
	// HandleArticleSequences corrects improper sequences of articles ('a' and 'an') within a line.
	words := strings.Fields(line)
	for i := 0; i < len(words); i++ {
		if strings.ToLower(words[i]) == "a" {
			if i < len(words)-1 && strings.ToLower(words[i+1]) == "a" {
				words[i] = "an"
			}
		}
	}
	return strings.Join(words, " ")
}

func ToBeModified(text []rune, count int, modType string) string {
	// ToBeModified applies specific modifications to the text based on the command and number provided.
	words := strings.Fields(string(text))
	modifiedWords := make([]string, len(words))

	for i := len(words) - 1; i >= 0; i-- {
		if count > 0 {
			switch modType {
			case "up":
				words[i] = strings.ToUpper(words[i])
			case "low":
				words[i] = strings.ToLower(words[i])
			case "cap":
				words[i] = strings.Title(strings.ToLower(words[i]))
			case "hex":
				if num, err := strconv.ParseInt(words[i], 16, 64); err == nil {
					words[i] = strconv.Itoa(int(num))
				}
			case "bin":
				if num, err := strconv.ParseInt(words[i], 2, 64); err == nil {
					words[i] = strconv.Itoa(int(num))
				}
			}
			count--
		}
		modifiedWords[i] = words[i]
	}

	return strings.Join(modifiedWords, " ")
}
