package containsduplicate

import "sort"

// time = O(nlogn)
// mem = O(1)
func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		if i+1 < len(nums) && nums[i] == nums[i+1] {
			return true
		}
	}

	return false
}

// time = O(n)
// mem = O(n)
func containsDuplicateMap(nums []int) bool {
	vals := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		_, ok := vals[nums[i]]
		if !ok {
			vals[nums[i]] = false
			continue
		}
		return true
	}
	return false
}
