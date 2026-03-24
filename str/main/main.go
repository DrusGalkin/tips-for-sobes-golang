package main

import (
	"awesomeProject2/str"
	"fmt"
)

func main() {
	str.StrCounter("Привет")

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

	str3 := reverseString("Привет")
	fmt.Println(str3)
}

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
