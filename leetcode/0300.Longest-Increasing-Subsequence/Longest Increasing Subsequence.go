package leetcode

import (
	"math"
	"sort"
)

// 解法1 O(n^2) DP
func lengthOfLIS(nums []int) int {
	dp, res := make([]int, len(nums)+1), 0
	dp[0] = 0
	for i := 1; i <= len(nums); i++ {
		for j := 1; j < i; j++ {
			if nums[j-1] < nums[i-1] {
				dp[i] = int(math.Max(float64(dp[i]), float64(dp[j])))
			}
		}
		dp[i] = dp[i] + 1
		res = int(math.Max(float64(res), float64(dp[i])))
	}

	return res
}

// 解法二 O(n log n) DP
func lengthOfLIS1(nums []int) int {
	dp := []int{}
	for _, num := range nums {
		i := sort.SearchInts(dp, num)
		if i == len(dp) {
			dp = append(dp, num)
		} else {
			dp[i] = num
		}
	}
	return len(dp)
}
