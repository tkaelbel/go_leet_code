package topkfrequentelements

func topKFrequent(nums []int, k int) []int {
	freq := make([][]int, len(nums)+1)
	nMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		currentNum := nums[i]
		_, ok := nMap[currentNum]
		if ok {
			count := nMap[currentNum] + 1
			nMap[currentNum] = count
		} else {
			nMap[currentNum] = 1
		}
	}

	for key, val := range nMap {
		freq[val] = append(freq[val], key)
	}

	res := make([]int, 0)
	for i := len(freq) - 1; i > 0; i-- {
		for _, j := range freq[i] {
			res = append(res, j)
			if len(res) == k {
				return res
			}
		}
	}

	return res
}
