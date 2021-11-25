package leetcode

import "math"

func maxUncrossedLines(A []int, B []int) int {
	res, dp := 0, make([][]int, len(A)+1)
	for i := range dp {
		dp[i] = make([]int, len(B)+1)
	}

	for i := 1; i < len(A)+1; i++ {
		for j := 1; j < len(B)+1; j++ {
			if A[i-1] == B[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				res = int(math.Max(float64(res), float64(dp[i][j])))
			} else {
				dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i][j-1])))
			}
		}
	}

	return res
}
