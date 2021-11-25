package main

import (
	"fmt"
)

func dp(n, k, i, cnt int) int {
	if n == i {
		return 1
	}

	ans := dp(n, k, i+1, 0)
	if cnt < k {
		ans += dp(n, k, i+1, cnt+1)
	}
	return ans
}

func solve(n, k int) int {
	return dp(n, k, 0, 0) % (10 ^ 9 + 7)
}

func main() {
	fmt.Println(solve(3, 2))
}
