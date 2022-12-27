package validanagram

import (
	"fmt"
	"sort"
	"strings"
)

// one way is to sort and then compare the strings
// very slow 29 ms
// lots of mem 6.4 mb
func isAnagramSort(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	splittedS := strings.Split(s, "")
	splittedT := strings.Split(t, "")
	sort.Strings(splittedS)
	sort.Strings(splittedT)
	// tempS := strings.Join(splittedS, "")
	// tempT := strings.Join(splittedT, "")

	return fmt.Sprint(splittedS) == fmt.Sprint(splittedT)
}

// another way is to count each letter and compare the values of each letter in two different maps
// faster than sort 16 ms
// better usage 3.5 mb
func isAnagramMap(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mapperS := map[string]int{}
	mapperT := map[string]int{}

	for i := 0; i < len(s); i++ {
		letterS := string(s[i])
		letterT := string(t[i])
		addLetter(letterS, mapperS)
		addLetter(letterT, mapperT)
	}

	return fmt.Sprint(mapperS) == fmt.Sprint(mapperT)
}

// faster 4 ms
// better usage 2.8 mb
func isAnagramMap2(s string, t string) bool {
	m1 := makeMap(s)
	m2 := makeMap(t)

	return fmt.Sprint(m1) == fmt.Sprint(m2)
}

func makeMap(s string) map[rune]int {
	ret := map[rune]int{}
	for _, str := range s {
		val, ok := ret[str]
		if ok {
			ret[str] = val + 1
		} else {
			ret[str] = 0
		}
	}
	return ret
}

func addLetter(l string, t map[string]int) {
	val, ok := t[l]
	if !ok {
		t[l] = 1
	} else {
		t[l] = val + 1
	}
}
