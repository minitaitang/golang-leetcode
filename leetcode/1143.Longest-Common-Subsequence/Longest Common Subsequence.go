package leetcode

import "math"

func longestCommonSubsequence(text1 string, text2 string) int {
	res, dp := 0, make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}

	for i := 1; i < len(text1)+1; i++ {
		for j := 1; j < len(text2)+1; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				res = int(math.Max(float64(res), float64(dp[i][j])))
			} else {
				dp[i][j] = int(math.Max(float64(dp[i-1][j]), float64(dp[i][j-1])))
			}
		}
	}

	return res
}
