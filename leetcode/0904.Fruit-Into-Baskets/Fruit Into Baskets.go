package leetcode

import "math"

func totalFruit(nums []int) int {
	i, ans := 0, 0
	k := 2
	win := make(map[int]int)
	for j := 0; j < len(nums); j++ {
		if win[nums[j]] == 0 {
			k -= 1
		}
		win[nums[j]] += 1
		for k < 0 {
			win[nums[i]] -= 1
			if win[nums[i]] == 0 {
				k += 1
			}
			i += 1
		}
		ans = int(math.Max(float64(ans), float64(j-i+1)))
	}

	return ans
}
