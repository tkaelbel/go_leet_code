package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	result := []int{}
out:
	for i := 0; i < len(nums); i++ {
		for j := len(nums) - 1; j > 0; j-- {
			if i != j && nums[i]+nums[j] == target {
				result = append(result, i)
				result = append(result, j)
				break out
			}
		}

	}
	return result
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{3, 3}, 6))
	fmt.Println(twoSum([]int{3, 2, 3}, 6))
	fmt.Println(twoSum([]int{1, 3, 4, 2}, 6))
}
