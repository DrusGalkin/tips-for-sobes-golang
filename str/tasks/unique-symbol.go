package tasks

import "strings"

func UniqueSymbol(str string) string {
	unique := make(map[rune]int)

	runes := []rune(str)
	for _, r := range runes {
		unique[r]++
	}

	var builder strings.Builder
	for _, r := range runes {
		if unique[r] == 1 {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}
