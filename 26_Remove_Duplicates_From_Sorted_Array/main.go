package main

import "fmt"

func removeDuplicates(nums []int) int {

	k := 0
	nmap := make(map[int]int)
	for idx := range nums {
		_, ok := nmap[nums[idx]]
		if ok {
			continue
		}
		nmap[nums[idx]] = nums[idx]
		nums[k] = nums[idx]
		k++
	}

	return k
}

func main() {
	firstIn := []int{1, 1, 2}
	fmt.Println(fmt.Sprintf("k: %v ", removeDuplicates(firstIn)) + fmt.Sprintf("array: %v", firstIn))
	secondIn := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(fmt.Sprintf("k: %v ", removeDuplicates(secondIn)) + fmt.Sprintf("array: %v", secondIn))
	thirdIn := []int{-1, 0, 0, 0, 0, 3, 3}
	fmt.Println(fmt.Sprintf("k: %v ", removeDuplicates(thirdIn)) + fmt.Sprintf("array: %v", thirdIn))
	fourthIn := []int{0, 0, 0, 0, 0}
	fmt.Println(fmt.Sprintf("k: %v ", removeDuplicates(fourthIn)) + fmt.Sprintf("array: %v", fourthIn))
	fithIn := []int{-3, -1, 0, 0}
	fmt.Println(fmt.Sprintf("k: %v ", removeDuplicates(fithIn)) + fmt.Sprintf("array: %v", fithIn))
}
