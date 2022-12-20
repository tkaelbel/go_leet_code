package main

func climbStairs(n int) int {
	if n == 1 || n == 2 {
		return n
	}

	currOneStep := 1
	currCombinedStep := 2

	for i := 3; i <= n; i++ {
		currOneStep, currCombinedStep = currCombinedStep, currOneStep+currCombinedStep
	}

	return currCombinedStep
}
