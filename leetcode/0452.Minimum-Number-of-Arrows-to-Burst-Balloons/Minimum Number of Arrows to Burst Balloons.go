package leetcode

import (
	"math"
	"sort"
)

// 解法一 代码也和435.无重叠区间一样
func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	sort.Sort(Points(points))
	dp, res := make([]int, len(points)), 0
	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < len(points); i++ {
		for j := 0; j < i; j++ {
			if points[i][0] > points[j][1] {
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
type Points [][]int

func (a Points) Len() int {
	return len(a)
}
func (a Points) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a Points) Less(i, j int) bool {
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
