package main

import (
	"fmt"
	"strings"
)

type Enigma struct {
	rotors       [][]int
	reflector    []int
	ringSettings []int
	positions    []int
}

// Функция для поворота роторов
func (e *Enigma) rotateRotors() {
	e.positions[0] = (e.positions[0] + 1) % 26
	if e.positions[0] == 0 { // Двойной шаг
		e.positions[1] = (e.positions[1] + 1) % 26
		if e.positions[1] == 0 {
			e.positions[2] = (e.positions[2] + 1) % 26
		}
	}
}

// Шифрование одного символа
func (e *Enigma) encodeChar(ch rune) rune {
	if ch < 'A' || ch > 'Z' {
		return ch // Пропускаем неалфавитные символы
	}

	index := int(ch - 'A')
	e.rotateRotors()

	// Проход через роторы в прямом направлении
	for i := 0; i < len(e.rotors); i++ {
		index = (e.rotors[i][(index+e.positions[i]-e.ringSettings[i]+26)%26] - e.positions[i] + e.ringSettings[i] + 26) % 26
	}

	// Проход через отражатель
	index = e.reflector[index]

	// Проход через роторы в обратном направлении
	for i := len(e.rotors) - 1; i >= 0; i-- {
		for j := 0; j < 26; j++ {
			if e.rotors[i][j] == (index+e.positions[i])%26 {
				index = (j - e.positions[i] + e.ringSettings[i] + 26) % 26
				break
			}
		}
	}

	return rune(index + 'A')
}

// Шифрование текста
func (e *Enigma) encrypt(text string) string {
	var result strings.Builder
	text = strings.ToUpper(text)
	for _, ch := range text {
		result.WriteRune(e.encodeChar(ch))
	}
	return result.String()
}

func main() {
	// Пример настроек
	rotors := [][]int{
		{4, 10, 12, 5, 11, 6, 3, 16, 21, 25, 13, 19, 14, 22, 24, 7, 23, 20, 18, 15, 0, 8, 1, 17, 2, 9},
		{0, 9, 3, 10, 18, 8, 17, 20, 23, 1, 11, 7, 22, 19, 12, 2, 16, 6, 25, 13, 15, 24, 5, 21, 14, 4},
		{1, 3, 5, 7, 9, 11, 2, 15, 17, 19, 23, 21, 25, 13, 24, 4, 8, 22, 6, 0, 10, 12, 20, 18, 16, 14},
	}
	reflector := []int{24, 17, 20, 7, 16, 18, 11, 3, 15, 23, 13, 6, 14, 10, 12, 8, 4, 1, 5, 25, 2, 22, 21, 9, 0, 19}
	ringSettings := []int{0, 0, 0}
	initialPositions := []int{0, 0, 0}

	// Инициализация машины
	enigma := Enigma{
		rotors:       rotors,
		reflector:    reflector,
		ringSettings: ringSettings,
		positions:    initialPositions,
	}

	// Шифрование текста
	plaintext := "HELLO WORLD"
	ciphertext := enigma.encrypt(plaintext)

	fmt.Println("Зашифрованный текст:", ciphertext)
}
