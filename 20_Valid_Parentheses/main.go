package main

import "fmt"

var parentheses = map[string]string{"(": ")", "{": "}", "[": "]"}

func isValid(s string) bool {
	// if we don't have even just return
	if len(s)%2 != 0 {
		return false
	}

	// // create slice with already used indices
	var closingPs []string
	for i := 0; i < len(s); i++ {
		p := parentheses[string(s[i])]
		if p != "" {
			closingPs = append(closingPs, p)
		} else {
			if len(closingPs) == 0 || string(s[i]) != closingPs[len(closingPs)-1] {
				return false
			}
			closingPs = closingPs[:len(closingPs)-1]
		}
	}

	return len(closingPs) == 0
}

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("{[]}"))
	fmt.Println(isValid("([)]"))
	fmt.Println(isValid("(){}}{"))
	fmt.Println(isValid(")(){}"))
	fmt.Println(isValid("[[[]"))
	fmt.Println(isValid("[])"))
	fmt.Println(isValid("(([]){})"))
}
