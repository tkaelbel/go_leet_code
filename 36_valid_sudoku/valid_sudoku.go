package validsudoku

func isValidSudoku(board [][]byte) bool {

	// keep index in mind
	rowMap := make(map[byte]int)
	colMap := make(map[int]map[byte]int)

	// initialize colMap
	for i := 0; i < 9; i++ {
		colMap[i] = make(map[byte]int)
	}

	subBox1 := make(map[byte]int)
	subBox2 := make(map[byte]int)
	subBox3 := make(map[byte]int)

	// row loop
	for rowIdx, loopRow := range board {
		for innerIdx, colVal := range loopRow {
			if colVal == '.' {
				continue
			}

			_, ok := rowMap[colVal]
			if !ok {
				rowMap[colVal] = 1
			} else {
				// we found a dupliate
				return false
			}

			_, ok = colMap[innerIdx][colVal]
			if !ok {
				colMap[innerIdx][colVal] = 1
			} else {
				return false
			}

			// fill subBox1
			if innerIdx < 3 {
				_, ok := subBox1[colVal]
				if !ok {
					subBox1[colVal] = 1
				} else {
					// we found a dupliate
					return false
				}
			}

			if innerIdx > 2 && innerIdx < 6 {
				_, ok := subBox2[colVal]
				if !ok {
					subBox2[colVal] = 1
				} else {
					// we found a dupliate
					return false
				}
			}

			if innerIdx > 5 {
				_, ok := subBox3[colVal]
				if !ok {
					subBox3[colVal] = 1
				} else {
					// we found a dupliate
					return false
				}
			}

		}
		// reset rowmap
		rowMap = make(map[byte]int)

		// reset boxes
		if rowIdx == 2 || rowIdx == 5 {
			subBox1 = make(map[byte]int)
			subBox2 = make(map[byte]int)
			subBox3 = make(map[byte]int)
		}
	}

	return true
}
