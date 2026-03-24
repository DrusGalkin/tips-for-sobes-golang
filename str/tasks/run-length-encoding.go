package tasks

import (
	"strconv"
	"strings"
	"unicode/utf8"
)

func RunLengthEncoding(str string) string {
	if utf8.RuneCountInString(str) == 0 || str == "" {
		return str
	}

	var builder strings.Builder
	var runes = []rune(str)
	var currentRune = runes[0]
	var count = 1

	for i := 0; i < len(runes); i++ {
		if runes[i] == currentRune {
			count++
		} else {
			builder.WriteRune(currentRune)
			builder.WriteString(strconv.Itoa(count))

			currentRune = runes[i]
			count = 1
		}
	}

	builder.WriteRune(currentRune)
	builder.WriteString(strconv.Itoa(count))

	compressed := builder.String()

	if utf8.RuneCountInString(compressed) >= utf8.RuneCountInString(str) {
		return str
	}

	return compressed
}
