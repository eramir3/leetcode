package main

import (
	"regexp"
	"strings"
)

func IsPalindromeString(text string) bool {
	re := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	cleaned := strings.ToLower(re.ReplaceAllString(text, ""))
	runes := []rune(cleaned)

	for i, j := 0, len(cleaned)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return cleaned == string(runes)
}
