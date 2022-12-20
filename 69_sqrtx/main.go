package main

func mySqrt(x int) int {
	i := 1
	ret := 0
	lastI := 0

	for {
		val := i * i
		if val == x {
			ret = i
			break
		}

		if val > x {
			ret = lastI
			break
		}
		lastI = i
		i++
	}

	return ret
}
