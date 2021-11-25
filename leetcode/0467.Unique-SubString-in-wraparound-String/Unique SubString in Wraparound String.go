package leetcode

import "math"

func findSubstringInWraproundString(p string) int {
	if len(p) == 0 || len(p) == 1 {
		return len(p)
	}
	m := make(map[byte]int)
	w := 1
	for i := 0; i < len(p); i++ {
		if p[i]-p[i-1] == 1 || p[i-1]-p[i] == 25 {
			w += 1
		} else {
			w = 1
		}
		m[p[i]] = int(math.Max(float64(m[p[i]]), float64(w)))
	}

	sum := 0
	for _, v := range m {
		sum += v
	}

	return sum
}
