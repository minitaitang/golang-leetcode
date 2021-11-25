package leetcode

import (
	"math"
)

func subSort(array []int) []int {
	n := len(array)
	l, r := -1, -1
	min := math.MinInt32
	max := math.MaxInt32

	for i := n - 1; i >= 0; i-- {
		if array[i] > min {
			l = i
		} else {
			min = array[i]
		}
	}

	for i, num := range array {
		if num < max {
			r = i
		} else {
			max = num
		}
	}

	return []int{l, r}
}
