package main

import "fmt"

var romansInt = map[string]int{"I": 1, "IV": 4, "V": 5, "IX": 9, "X": 10, "XL": 40, "L": 50, "XC": 90, "C": 100, "CD": 400, "D": 500, "CM": 900, "M": 1000}

func romanToInt(s string) int {
	var result int = 0
	for i := 0; i < len(s); i++ {
		char := string(s[i])

		if i < len(s)-1 {
			charAfter := string(s[i+1])

			if romansInt[char+charAfter] != 0 {
				result += romansInt[char+charAfter]
				i++
				continue
			}
		}

		if romansInt[char] != 0 {
			result += romansInt[char]
		}
	}
	return result
}

func main() {
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("LVIII"))
	fmt.Println(romanToInt("MCMXCIV"))
}
