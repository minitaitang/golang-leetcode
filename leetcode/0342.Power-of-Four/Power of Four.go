package leetcode

// 1.是二的幂次方
// 2.减去 1 是三的倍数

func isPowerOfFour(num int) bool {
	return num > 0 && (num&(num-1)) == 0 && (num-1)%3 == 0
}
