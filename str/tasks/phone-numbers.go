package tasks

import "strings"

func PhoneNumber(str string) string {
	var builder strings.Builder
	runes := []rune(str)

	builder.Grow(len(str) + len("+ () -"))

	for index, rune := range runes {
		if index == 0 {
			runeStr := string(rune)
			if runeStr == "9" || runeStr == "8" {
				runeStr = "7"
			}
			builder.WriteString("+" + runeStr + " ")
		} else if index == 1 {
			builder.WriteString("(" + string(rune))
		} else if index == 3 {
			builder.WriteString(string(rune) + ") ")
		} else if index == 6 {
			builder.WriteString(string(rune) + " ")
		} else if index == 8 {
			builder.WriteString(string(rune) + "-")
		} else {
			builder.WriteRune(rune)
		}

	}

	return builder.String()
}
