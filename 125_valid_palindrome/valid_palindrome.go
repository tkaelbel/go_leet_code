package validpalindrome

import (
	"regexp"
	"strings"
	"unicode"
)

// first variant
func isPalindrome(s string) bool {
	alphaNumericOnly := regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllLiteralString(s, "")
	alphaNumericOnly = strings.ReplaceAll(alphaNumericOnly, " ", "")
	alphaNumericOnly = strings.ToLower(alphaNumericOnly)
	ret := ""
	for i := len(alphaNumericOnly) - 1; i >= 0; i-- {
		ret = ret + string(alphaNumericOnly[i])
	}
	return ret == alphaNumericOnly
}

func isPalindrome1(s string) bool {
	leftIdx := 0
	rightIdx := len(s) - 1

	runes := []rune(s)
	for leftIdx < rightIdx {
		left := unicode.ToLower(runes[leftIdx])
		right := unicode.ToLower(runes[rightIdx])

		if !(unicode.IsLetter(left) || unicode.IsDigit(left)) {
			leftIdx++
			continue
		}

		if !(unicode.IsLetter(right) || unicode.IsDigit(right)) {
			rightIdx--
			continue
		}

		if left != right {
			return false
		}

		rightIdx--
		leftIdx++
	}

	return true
}
