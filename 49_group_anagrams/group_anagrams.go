package groupanagrams

import (
	"fmt"
	"sort"
	"strings"
)

func groupAnagrams(strs []string) [][]string {
	ret := make([][]string, 0)

	groups := make(map[string][]string)

	for i := 0; i < len(strs); i++ {
		// create map for each letter
		m := makeMap(strs[i])
		id := fmt.Sprint(m)
		// check if there is no entry in the group map
		_, ok := groups[id]
		if !ok {
			// create a new entry for the current anagram group
			arr := make([]string, 1)
			arr[0] = strs[i]
			groups[id] = arr
		} else {
			arr := groups[id]
			groups[id] = append(arr, strs[i])
		}
	}

	for k := range groups {
		ret = append(ret, groups[k])
	}

	return ret
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

func groupAnagramsSort(strs []string) [][]string {
	ret := make([][]string, 0)
	groups := make(map[string][]string, 0)

	for i := 0; i < len(strs); i++ {

		letters := strings.Split(strs[i], "")
		sort.Strings(letters)
		sorted := strings.Join(letters, "")

		// check if there is no entry in the group map
		_, ok := groups[sorted]
		if !ok {
			// create a new entry for the current anagram group
			arr := make([]string, 1)
			arr[0] = strs[i]
			groups[sorted] = arr
		} else {
			arr := groups[sorted]
			groups[sorted] = append(arr, strs[i])
		}
	}

	for k := range groups {
		ret = append(ret, groups[k])
	}

	return ret
}
