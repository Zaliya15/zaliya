package main

import (
	"fmt"
	"strings"
)

func vigenereEncrypt(plaintext, key string) string {
	var encryptedText strings.Builder
	keyIndex := 0
	keyLength := len(key)

	for _, charPlaintext := range plaintext {
		if isAlpha(charPlaintext) {
			base := 'A'
			if charPlaintext >= 'a' && charPlaintext <= 'z' {
				base = 'a'
			}
			charKey := toupper(rune(key[keyIndex%keyLength]))
			encryptedChar := (charPlaintext-base+(charKey-'A'))%26 + base
			encryptedText.WriteRune(encryptedChar)
			keyIndex++
		} else {
			encryptedText.WriteRune(charPlaintext)
		}
	}

	return encryptedText.String()
}

func vigenereDecrypt(ciphertext, key string) string {
	var decryptedText strings.Builder
	keyIndex := 0
	keyLength := len(key)

	for _, charCiphertext := range ciphertext {
		if isAlpha(charCiphertext) {
			base := 'A'
			if charCiphertext >= 'a' && charCiphertext <= 'z' {
				base = 'a'
			}
			charKey := toupper(rune(key[keyIndex%keyLength]))
			decryptedChar := (charCiphertext-base-(charKey-'A')+26)%26 + base
			decryptedText.WriteRune(decryptedChar)
			keyIndex++
		} else {
			decryptedText.WriteRune(charCiphertext)
		}
	}

	return decryptedText.String()
}

func isAlpha(c rune) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')
}

func toupper(c rune) rune {
	if c >= 'a' && c <= 'z' {
		return c - ('a' - 'A')
	}
	return c
}

func main() {
	var plaintext, key string

	fmt.Print("Введите текст: ")
	fmt.Scanln(&plaintext)

	fmt.Print("Введите ключ: ")
	fmt.Scanln(&key)

	encrypted := vigenereEncrypt(plaintext, key)
	decrypted := vigenereDecrypt(encrypted, key)

	fmt.Println("Зашифрованный текст:", encrypted)
	fmt.Println("Расшифрованный текст:", decrypted)
}
