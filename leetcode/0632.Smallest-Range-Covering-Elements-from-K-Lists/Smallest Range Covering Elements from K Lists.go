package leetcode

import "sort"

type pair struct {
	val   int
	group int
}

func smallestRange(nums [][]int) []int {
	k := len(nums)
	pairs := []*pair{}

	for i := 0; i < k; i++ {
		for _, num := range nums[i] {
			pairs = append(pairs, &pair{
				val:   num,
				group: i,
			})
		}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].val < pairs[j].val
	})

	res := []int{}
	groups, start := 0, 0
	groupCount := map[int]int{}
	for i := 0; i < len(pairs); i++ {
		if _, ok := groupCount[pairs[i].group]; !ok {
			groups++
		}
		groupCount[pairs[i].group]++

		if k == groups {
			for groupCount[pairs[start].group] > 1 {
				groupCount[pairs[start].group]--
				start++
			}
			if len(res) == 0 || res[1]-res[0] > pairs[i].val-pairs[start].val {
				res = []int{pairs[start].val, pairs[i].val}
			}
		}
	}

	return res
}
