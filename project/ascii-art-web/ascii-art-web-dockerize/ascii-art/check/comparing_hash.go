package check

import (
	"bufio"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"os"
)

func CalculateSHA256Hash(filePath string) ([]byte, error) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Print("Function name is: CalculateSHA256Hash-----> ")
		fmt.Println("You have a problem with style file. And it's error message is:", err)
		return []byte{}, err
	}
	defer file.Close()

	hash := sha256.New()
	_, err = io.Copy(hash, file)

	if err != nil {
		fmt.Print("Function name is: CalculateSHA256Hash-----> ")
		fmt.Println("There is problem with hashing operation, and it's error message is:", err)
		return []byte{}, err
	}
	return hash.Sum(nil), nil
}

func Comparing_hash(filePath string) error {
	current_file_name := "../ascii-art/banner/" + filePath
	first, err := CalculateSHA256Hash(current_file_name)
	if err != nil {
		fmt.Println("Function name: Comparing_hash")
		return err
	}

	original_file_name := "../ascii-art/check/file_key/" + filePath

	file, err := os.Open(original_file_name)
	if err != nil {
		fmt.Print("Function name is: Comparing_hash-----> ")
		fmt.Println("You have a problem with original file name, and it's error message is:", err)
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var text string
	for scanner.Scan() {
		text = scanner.Text()
	}
	second := []byte(text)

	for i := 0; i < len(first); i++ {
		if first[i] != second[i] {
			fmt.Print("Function name is: Comparing_hash-----> ")
			fmt.Println("There is mismatch between original file and current file. The banner is:", filePath)
			return errors.New("There is mismatch between original file and current file.")
		}
	}

	return nil
}
