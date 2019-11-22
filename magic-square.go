package main

import "fmt"

func formingMagicSquare(s [][]int32) int32 {
	basicPatterns := createBasicPatterns()

	minCost := comparePatternsToS(basicPatterns, s)

	return minCost
}

func createBasicPatterns() [][][]int32 {
	basicPatterns := [][][]int32{
		[][]int32{
			[]int32{8, 1, 6}, []int32{3, 5, 7}, []int32{4, 9, 2},
		},
		{
			[]int32{6, 1, 8}, []int32{7, 5, 3}, []int32{2, 9, 4},
		},
		{
			[]int32{4, 9, 2}, []int32{3, 5, 7}, []int32{8, 1, 6},
		},
		{
			[]int32{2, 9, 4}, []int32{7, 5, 3}, []int32{6, 1, 8},
		},
		{
			[]int32{8, 3, 4}, []int32{1, 5, 9}, []int32{6, 7, 2},
		},
		{
			[]int32{4, 3, 8}, []int32{9, 5, 1}, []int32{2, 7, 6},
		},
		{
			[]int32{6, 7, 2}, []int32{1, 5, 9}, []int32{8, 3, 4},
		},
		{
			[]int32{2, 7, 6}, []int32{9, 5, 1}, []int32{4, 3, 8},
		},
	}

	return basicPatterns
}

func comparePatternsToS(basicPatterns [][][]int32, s [][]int32) int32 {
	totals := []int32{1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000}
	var min int32 = 1000
	for i := 0; i < 8; i++ {
		totals[i] = calculateCost(basicPatterns[i], s)
		if min > totals[i] {
			min = totals[i]
		}
	}

	return min
}

func calculateCost(basicPattern [][]int32, s [][]int32) int32 {
	var total int32
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			total += absValue(basicPattern[i][j] - s[i][j])
		}
	}

	return total
}

func absValue(value int32) int32 {
	if value < 0 {
		return -value
	}

	return value
}

func main() {
	basicPatterns := createBasicPatterns()
	fmt.Println(basicPatterns[0][0])
}

func testMagicSquare(s [][]int32) bool {
	totals := []int32{}
	totals[0] = s[0][0] + s[0][1] + s[0][2]
	totals[1] = s[1][0] + s[1][1] + s[1][2]
	totals[2] = s[2][0] + s[2][1] + s[2][2]

	totals[3] = s[0][0] + s[1][0] + s[2][0]
	totals[4] = s[0][1] + s[1][1] + s[2][1]
	totals[5] = s[0][2] + s[1][2] + s[2][2]

	totals[6] = s[0][0] + s[1][1] + s[2][2]
	totals[7] = s[2][0] + s[1][1] + s[0][2]

	for i := 0; i < 7; i++ {
		if totals[i] != 15 {
			return false
		}
	}

	return true
}

// center is always going to be 5!
// centerDiff := 5 - s[1][1]
// s[1][1] = centerDiff + 5
