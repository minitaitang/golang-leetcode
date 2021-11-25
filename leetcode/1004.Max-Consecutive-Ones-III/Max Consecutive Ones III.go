package leetcode

import "math"

func longestOnes(A []int, K int) int {
	left, maxWindow, zeroCount := 0, 0, 0
	for right := 0; right < len(A); right++ {
		if A[right] == 0 {
			zeroCount++
		}

		for zeroCount > K {
			if A[left] == 0 {
				zeroCount--
			}
			left++
		}
		maxWindow = int(math.Max(float64(maxWindow), float64(right-left+1)))
	}
	return maxWindow
}
