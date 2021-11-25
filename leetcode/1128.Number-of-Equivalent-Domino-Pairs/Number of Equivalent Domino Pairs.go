package leetcode

import (
	"strconv"
)

func numEquivDominoPairs(dominoes [][]int) int {
	if len(dominoes) == 0 {
		return 0
	}
	result := 0
	cntM := make(map[string]int)
	var key string

	for _, dominoe := range dominoes {
		if dominoe[0] > dominoe[1] {
			key = strconv.Itoa(dominoe[0]) + strconv.Itoa(dominoe[1])
		} else {
			key = strconv.Itoa(dominoe[1]) + strconv.Itoa(dominoe[0])
		}
		cntM[key]++
	}

	for _, v := range cntM {
		if v > 1 {
			result += (v * (v - 1)) / 2
		}
	}

	return result
}
