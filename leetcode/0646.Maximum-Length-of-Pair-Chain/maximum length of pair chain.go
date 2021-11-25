package leetcode

import (
	"math"
	"sort"
)

// 解法一 代码和435.无重叠区间一样
func findLongestChain(pairs [][]int) int {
	if len(pairs) == 0 {
		return 0
	}

	sort.Sort(Pairs(pairs))
	dp, res := make([]int, len(pairs)), 0
	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < len(pairs); i++ {
		for j := 0; j < i; j++ {
			if pairs[i][0] > pairs[j][1] {
				dp[i] = int(math.Max(float64(dp[i]), float64(dp[j]+1)))
			}
		}
	}

	for _, v := range dp {
		res = int(math.Max(float64(res), float64(v)))
	}

	return res
}

// Pairs define
type Pairs [][]int

func (a Pairs) Len() int {
	return len(a)
}
func (a Pairs) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a Pairs) Less(i, j int) bool {
	for k := 0; k < len(a[i]); k++ {
		if a[i][k] < a[j][k] {
			return true
		} else if a[i][k] == a[j][k] {
			continue
		} else {
			return false
		}
	}
	return true
}
