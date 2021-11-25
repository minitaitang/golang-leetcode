package main

import (
	"fmt"
)

func solve(nums []int) int {
	ans := 1
	n := len(nums)
	dp := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}

	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			dp[0][i] = dp[0][i-1] + 1
			dp[1][i] = dp[1][i-1] + 1
		} else {
			dp[0][i] = 1
			dp[1][i] = 1
		}
		if i > 1 && nums[i] > nums[i-2] {
			if dp[1][i] < dp[0][i-2]+1 {
				dp[1][i] = dp[0][i-2] + 1
			}
		}

		if ans < dp[0][i] {
			ans = dp[0][i]
		}
		if ans < dp[1][i] {
			ans = dp[1][i]
		}
	}

	return ans
}

func main() {
	fmt.Println(solve([]int{1, 2, 3, 0, 4}))
}
