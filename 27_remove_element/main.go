package main

import (
	"sort"
)

func removeElement(nums []int, val int) int {

	k := 0
	nmap := make(map[int]int)
	sort.Ints(nums)
	for _, item := range nums {
		if item != val {
			nmap[k] = item
			k++
		}
	}

	for idx, item := range nmap {
		nums[idx] = item
	}

	return k
}
