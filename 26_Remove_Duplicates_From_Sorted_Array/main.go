package main

import "fmt"

// in progress..
func removeDuplicates(nums []int) int {
	k := 0

	initial := []int{}
	initial = append(initial, nums...)

	// filter duplicates
	for i := 0; i < len(nums); i++ {
		for j := len(nums) - 1; j >= 0; j-- {
			if nums[i] == nums[j] && i != j {
				nums[i] = 0
				continue
			}
		}
	}

	// shift nulls to back (if no null is in there)
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[count] = nums[i]
			count++
		}
	}

	for count < len(nums) {
		nums[count] = 0
		count++
	}

	// add leading 0
	if initial[0] == 0 && nums[len(nums)-1] == 0 {
		insert(nums, 0, 0)
	}

	// fill k
	for i := 0; i < len(nums); i++ {
		if i == 0 && nums[i] == 0 {
			k++
		}
		if nums[i] != 0 {
			k++
		}
	}

	return k
}

func insert(a []int, index int, value int) []int {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

func main() {
	firstIn := []int{1, 1, 2}
	fmt.Println(fmt.Sprintf("k: %v ", removeDuplicates(firstIn)) + fmt.Sprintf("array: %v", firstIn))
	secondIn := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(fmt.Sprintf("k: %v ", removeDuplicates(secondIn)) + fmt.Sprintf("array: %v", secondIn))
}
