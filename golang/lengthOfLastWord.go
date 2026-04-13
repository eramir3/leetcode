package main

import "strings"

func LengthOfLastWord(s string) int {
	text := strings.Fields(s)
	return len(text[len(text)-1])
}
