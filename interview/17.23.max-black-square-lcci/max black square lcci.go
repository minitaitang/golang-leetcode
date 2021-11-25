package leetcode

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 从左到右，从上到下
func findSquare(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return []int{}
	}
	n := len(matrix[0])
	if n == 0 {
		return []int{}
	}

	// 预先计算，对于每一行，当以某个位置i结尾时，最大的连续黑色的长度
	dp1 := make([][]int, m)
	for i := 0; i < m; i++ {
		dp1[i] = make([]int, n)

		prev := 0
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				prev++
				dp1[i][j] = prev
			} else {
				prev = 0
			}
		}
	}

	// 预先计算，对于每一列，当以某个位置j结尾时，最大的连续黑色的长度
	dp2 := make([][]int, n)
	for j := 0; j < n; j++ {
		dp2[j] = make([]int, m)

		prev := 0
		for i := 0; i < m; i++ {
			if matrix[i][j] == 0 {
				prev++
				dp2[j][i] = prev
			} else {
				prev = 0
			}
		}
	}

	// 找最大矩形
	var maxSize int = 0
	var x, y int // 右下角的点
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// i和j是右下角，尝试每一个可能的长度，最大的可能长度是两个边长的最小值
			size := min(dp1[i][j], dp2[j][i])
			for size > 0 {
				// 查看 矩形top那一行和left那一列是不是也都是连续黑色
				if dp1[i-size+1][j] >= size && dp2[j-size+1][i] >= size {
					break
				}
				size--
			}

			if size > maxSize { // 而非>=, 保证相同大小时取左上角
				x, y = i, j
				maxSize = size
			}
		}
	}

	if maxSize == 0 { // 全白就返回空
		return []int{}
	}

	return []int{x - maxSize + 1, y - maxSize + 1, maxSize}
}
