package binarysearch

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	if len(nums) == 1 && nums[0] == target {
		return 0
	}

	mid := len(nums) / 2
	if nums[mid] == target {
		return mid
	}

	if nums[mid] < target {
		for i := mid + 1; i < len(nums); i++ {
			if nums[i] == target {
				return i
			}
		}
		return -1
	}

	if nums[mid] > target {
		for i := mid - 1; i >= 0; i-- {
			if nums[i] == target {
				return i
			}
		}
		return -1
	}

	return -1

}
