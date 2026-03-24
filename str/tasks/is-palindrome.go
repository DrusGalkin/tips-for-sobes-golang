package tasks

import (
	"strings"
	"unicode"
)

// Напишите функцию, которая проверяет, можно ли прочитать строку одинаково
// в обоих направлениях, игнорируя пробелы, знаки припинания и регистр букв
// Пример:
//		1) топот -> true
//		2) топор -> false

func IsPalindrome(str string) bool {
	// Приводим к нижнему регистру и убираем лишние символы
	var cleaned []rune
	for _, r := range strings.ToLower(str) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			cleaned = append(cleaned, r)
		}
	}

	// Проверка на палиндром
	left, right := 0, len(cleaned)-1
	for left < right {
		if cleaned[left] != cleaned[right] {
			return false
		}
		left++
		right--
	}

	return true
}
