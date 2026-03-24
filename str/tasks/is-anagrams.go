package tasks

import (
	"strings"
	"unicode"
)

//	Напишите функцию AreAnagrams(s1, s2 string) bool, которая проверяет,
//	являются ли две строки анаграммами друг друга (игнорируя регистр и пробелы).
//
//	Пример:
//		"Listen" и "Silent" -> true
//		"hello" и "billion" -> false

func IsAnagrams(str1, str2 string) bool {
	wres1 := runeToMap(str1)
	wres2 := runeToMap(str2)

	if len(wres1) != len(wres2) {
		return false
	}

	for _, char := range str1 {
		if wres1[char] != wres2[char] {
			return false
		}
	}

	return true
}

func runeToMap(str string) (wres map[rune]int) {
	wres = make(map[rune]int)
	for _, r := range strings.Trim(strings.ToLower(str), " ") {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			wres[r]++
		}
	}
	return
}
