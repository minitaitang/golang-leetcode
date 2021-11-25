package leetcode

func numSubarrayBoundedMax(A []int, L int, R int) int {
	return notGreater(A, R) - notGreater(A, L-1)
}

func notGreater(A []int, R int) int {
	ans, cnt := 0, 0
	for _, a := range A {
		if a <= R {
			cnt += 1
		} else {
			cnt = 0
		}
		ans += cnt
	}

	return ans
}
