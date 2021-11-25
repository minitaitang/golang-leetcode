package leetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, _ := range nums {
		diff := target - nums[i]
		if j, ok := m[diff]; ok {
			return []int{i, j}
		} else {
			m[nums[i]] = i
		}
	}

	return []int{}
}
