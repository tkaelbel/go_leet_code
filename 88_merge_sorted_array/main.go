package main

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums1) != (m + n) {
		return
	}

	if len(nums2) != n {
		return
	}

	if n == 0 {
		return
	}

	if m == 0 {
		for i := 0; i < len(nums1); i++ {
			nums1[i] = nums2[i]
		}
		return
	}

	// fill missing nulls in nums1
	for i := 0; i < len(nums2); i++ {
		nums1[m+i] = nums2[i]
	}

	sort.Ints(nums1)
}
