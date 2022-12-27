package twosum

// time 76 ms
// memory 3,7 mb
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

// time 60 ms
// memory 3,5 mb
func twoSum1(nums []int, target int) []int {
	result := make([]int, 2)
out:
	for i := 0; i < len(nums); i++ {
		for j := len(nums) - 1; j > 0; j-- {
			if i != j && nums[i]+nums[j] == target {
				result[0] = i
				result[1] = j
				break out
			}
		}

	}
	return result
}

// time 8 ms
// memory 4,2 mb
func twoSum2(nums []int, target int) []int {
	result := make([]int, 2)
	numMap := make(map[int]int)

	for idx, num := range nums {
		calc := target - num
		valIdx, ok := numMap[calc]
		if !ok {
			numMap[num] = idx
		} else {
			result[0] = valIdx
			result[1] = idx
			break
		}
	}

	return result
}
