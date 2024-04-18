package functions

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func EmpetyTextsCheker(textArr []string) (bool, error) {
	var EmpetyTexts bool = false

	for i := 0; i < len(textArr); i++ {
		if len(textArr[i]) == 0 {
			EmpetyTexts = true
		} else {
			EmpetyTexts = false
			return EmpetyTexts, nil
		}
	}
	return EmpetyTexts, nil
}

func AsciiCheeker(s string) bool {
	for i := 0; i < len(s); i++ {
		if !((s[i] >= ' ' && s[i] <= '~') || s[i] == 10 || s[i] == 13) {
			return false
		}
	}
	return true
}

func GetTerminalSize() (width, height int, err error) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		return 0, 0, err
	}
	size := strings.Split(strings.TrimSpace(string(out)), " ")
	if len(size) != 2 {
		return 0, 0, fmt.Errorf("Invalid stty size output format!")
	}
	fmt.Sscanf(size[0], "%d", &height)
	fmt.Sscanf(size[1], "%d", &width)
	return width, height, nil
}

const (
	standard   = "e194f1033442617ab8a78e1ca63a2061f5cc07a3f05ac226ed32eb9dfd22a6bf"
	shadow     = "26b94d0b134b77e9fd23e0360bfd81740f80fb7f6541d1d8c5d85e73ee550f73"
	thinkertoy = "64285e4960d199f4819323c4dc6319ba34f1f0dd9da14d07111345f5d76c3fa3"
)

func CalculateFileHash(style string) error {
	filePath := style + ".txt"
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		return err
	}

	hashInBytes := hash.Sum(nil)
	hashString := hex.EncodeToString(hashInBytes)

	switch style {
	case "standard":
		if hashString == standard {
			return nil
		}
	case "shadow":
		if hashString == shadow {
			return nil
		}
	case "thinkertoy":
		if hashString == thinkertoy {
			return nil
		}
	}
	return errors.New("Bad Hash")
}
