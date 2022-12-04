package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	base := strconv.Itoa(x)
	res := ""
	for i := len(base) - 1; i >= 0; i-- {
		res = res + string(base[i])
	}
	return base == res
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(-121))
	fmt.Println(isPalindrome(10))
}
