package main

import (
	"ascii-art-fs/functions"
	"flag"
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	// Определение флага для вывода в файл
	var outputFileName string
	flag.StringVar(&outputFileName, "output", "", "Имя файла для записи")

	// Определение флага для разворачивания ASCII-арта
	var reverseFileName string

	flag.StringVar(&reverseFileName, "reverse", "", "Имя файла для разворачивания ASCII-арта")

	flag.Parse()

	// Проверяем, был ли указан флаг --reverse
	if reverseFileName != "" {
		if ".txt" != filepath.Ext(reverseFileName) {
			fmt.Println("Wrong extension, please use txt file")
			return
		}
		err := functions.ReverseAsciiArt(reverseFileName)
		if err != nil {
			fmt.Println("Ошибка при разворачивании ASCII-арта:", err)
			return
		}

		return
	}

	// Получение аргументов строки
	args := flag.Args()

	// Если не указан ни один аргумент, выводим сообщение о использовании
	if len(args) < 1 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println("EX: go run . --output=<fileName.txt> something standard")
		return
	}

	// Извлечение строки и стиля из аргументов
	text := args[0]
	style := "standard"
	if len(args) > 1 {
		style = args[1]
		if style != "standard" && style != "thinkertoy" && style != "shadow" {
			fmt.Println("Wrong style! available styles: standard,shadow,thinkertoy")
			return
		}
	}

	// Проверка наличия флага --output
	if outputFileName != "" {
		// Запись в файл
		if len(outputFileName) <= 4 {
			fmt.Println("ERROR: Too short name for file")
			return
		}
		if !strings.HasSuffix(outputFileName, ".txt") {
			fmt.Println("Output file must be in txt format")
			return
		}

		err := functions.WriteToFile(outputFileName, text, style)
		if err != nil {
			fmt.Println("Ошибка при записи в файл:", err)
			return
		}

		fmt.Println("Информация успешно записана в файл:", outputFileName)
	} else {
		// Вывод на экран
		err := functions.PrintAsciiArt(text, style)
		if err != nil {
			fmt.Println("Ошибка при выводе ASCII-арт:", err)
			return
		}
	}
}
