package main

import (
	"awesomeProject2/str"
	"awesomeProject2/str/tasks"
	"fmt"
	"strings"
	"unicode"
)

func main() {
	str.StrCounter("Привет")

	pricol := "Hello 😂"

	for _, r := range pricol {
		fmt.Print(r) // смайли(⌐■_■)ки тоже есть в таблице аски (■_■⌐)
	}
	fmt.Println()

	// хорошо
	str1 := str.StrBuilder(
		"c",
		"a",
		"с",
		"а",
		"л",
	)

	// плохо
	str2 := str.ConcatPlus(
		"c",
		"a",
		"с",
		"а",
		"л",
		"228",
	)

	fmt.Println(str1, str2)

	// Задачки
	fmt.Println("************************************")

	// Переворачиваем строку
	//str3 := tasks.ReverseString("Привет")
	str3 := reverseString("Привет")
	fmt.Println(str3)

	// Проверка на полиндром
	fmt.Println(tasks.IsPalindrome("212"))
	fmt.Println(isPalindrome("топот"))

	// Уникальный символ в строке

	str4 := tasks.UniqueSymbol("Привет")
	fmt.Println(str4)

	// это анаграмма
	fmt.Println(tasks.IsAnagrams("123", "321 "))

	// сжатие строки
	fmt.Println(tasks.RunLengthEncoding("aaaa                                                     ssss"))

	fmt.Printf("tasks.PhoneNumber(\"79093331904\"): %v\n", tasks.PhoneNumber("79093331904"))

}

// тут я проверяю свои знания епт
func reverseString(str string) string {
	runes := []rune(str)
	left, right := 0, len(runes)-1

	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}

	return string(runes)
}

func isPalindrome(str string) bool {
	var cleaned []rune
	for _, r := range strings.ToLower(str) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			cleaned = append(cleaned, r)
		}
	}

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
