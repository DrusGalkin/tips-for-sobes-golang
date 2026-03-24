package str

import (
	"fmt"
	"unicode/utf8"
)

// Длинна строк
func StrCounter(str string) {
	fmt.Println(len(str), "- неправильно, тк через len")
	fmt.Println(len([]rune(str)), "- вроде правильно, но долго")
	fmt.Println(utf8.RuneCountInString(str), "- самый лучший вариант")
}
