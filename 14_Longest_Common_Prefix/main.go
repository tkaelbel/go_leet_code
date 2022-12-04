package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	firstWord := strs[0]
	otherWords := strs[1:]
	result := ""
	for i := 0; i < len(firstWord); i++ {
		firstWordChar := firstWord[i]

		matchesChar := false
		for _, otherWord := range otherWords {
			if len(otherWord) > i {
				otherWordChar := otherWord[i]
				matchesChar = firstWordChar == otherWordChar

				if !matchesChar {
					break
				}

				if matchesChar {
					continue
				}
			} else {
				matchesChar = false
				break
			}
		}

		if matchesChar {
			result = result + string(firstWordChar)
			matchesChar = false
		} else {
			break
		}
	}

	return result
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Println(longestCommonPrefix([]string{"a"}))
	fmt.Println(longestCommonPrefix([]string{"abab", "aba", ""}))
	fmt.Println(longestCommonPrefix([]string{"c", "acc", "ccc"}))
}
