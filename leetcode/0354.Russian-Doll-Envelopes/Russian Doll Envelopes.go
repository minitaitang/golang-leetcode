package leetcode

import (
	"math"
	"sort"
)

func maxEnvelopes(envelopes [][]int) int {
	if len(envelopes) == 0 {
		return 0
	}

	sort.Sort(Envelopes(envelopes))
	dp, res := make([]int, len(envelopes)), 0
	for i := range dp {
		dp[i] = 1
	}

	for i := 0; i < len(envelopes); i++ {
		for j := i + 1; j < len(envelopes); j++ {
			if envelopes[i][0] < envelopes[j][0] && envelopes[i][1] < envelopes[j][1] {
				dp[j] = int(math.Max(float64(dp[j]), float64(dp[i]+1)))
			}
		}
	}

	for _, v := range dp {
		res = int(math.Max(float64(res), float64(v)))
	}

	return res
}

// Intervals define
type Envelopes [][]int

func (a Envelopes) Len() int {
	return len(a)
}
func (a Envelopes) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a Envelopes) Less(i, j int) bool {
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
