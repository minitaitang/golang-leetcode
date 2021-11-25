package leetcode

func distributeCandies(candies []int) int {
	m := make(map[int]struct{}, len(candies))
	for _, candy := range candies {
		m[candy] = struct{}{}
	}
	res := len(m)
	if len(candies)/2 < res {
		return len(candies) / 2
	}
	return res
}
