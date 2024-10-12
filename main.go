package main

import (
	"fmt"
)

func atbashCipher(input string) string {
	result := ""

	for _, char := range input {
		if char >= 'A' && char <= 'Z' {
			result += string('Z' - (char - 'A'))
		} else if char >= 'a' && char <= 'z' {
			result += string('z' - (char - 'a'))
		} else {
			result += string(char) // сохраняем пробелы и другие символы
		}
	}

	return result
}

func main() {
	var input string
	fmt.Println("Введите текст:")
	fmt.Scanln(&input)

	encryptedText := atbashCipher(input)
	fmt.Println("Зашифрованный текст:", encryptedText)
}
