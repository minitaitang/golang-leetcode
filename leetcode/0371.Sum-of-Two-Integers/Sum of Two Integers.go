package leetcode

import "fmt"

func getSum(a int, b int) int {
	fmt.Println(a, b)
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}

	return getSum((a&b)<<1, a^b)
}

// func main() {
// 	fmt.Println(getSum(10, 7))
// }
