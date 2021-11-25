package leetcode

func increasingDigits(n int) int {
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, 10)
	}
	for i := 1; i < 10; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < n; i++ {
		for j := 1; j < 10; j++ {
			for k := 0; k < j; k++ {
				dp[i][j] += dp[i-1][k]
			}
		}
	}

	var result int
	for j := 1; j < 10; j++ {
		result += dp[n-1][j]
	}

	return result
}

// func main() {
// 	fmt.Println(increasingDigits(9))
// }
