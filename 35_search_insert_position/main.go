package main

func searchInsert(nums []int, target int) int {
	return binarySearch(nums, 0, len(nums)-1, target)
}

func binarySearch(nums []int, low int, high int, target int) int {
	mid := low + (high-low)/2
	if high >= low {
		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			return binarySearch(nums, low, mid-1, target)
		}

		return binarySearch(nums, mid+1, high, target)
	}
	return mid
}
