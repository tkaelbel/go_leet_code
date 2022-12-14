package main

import "strings"

func lengthOfLastWord(s string) int {
	if s == "" {
		return 0
	}

	s = strings.TrimSpace(s)
	words := strings.Split(s, " ")

	if len(words) == 1 {
		return len(words[0])
	}

	for i := len(words) - 1; i >= 0; i-- {
		if words[i] != "" {
			return len(words[i])
		}
	}
	return 0
}
