package leetcode

func removeDuplicateLetters(s string) string {
	if len(s) == 0 {
		return ""
	}
	// 1 哈希表， key: 字母， value: 出现次数
	var hm [26]int
	for _, c := range s {
		hm[c-'a']++
	}
	// 2 该字符是否已经放入stack中
	var visited [26]bool
	// 栈
	stack := make([]rune, 0)

	// 3 遍历字符串
	for _, c := range s {
		hm[c-'a']--
		// 3.1 如果stack中已经存在，则跳过
		if visited[c-'a'] {
			continue
		}
		// 如果栈中有数据，且栈顶字符大于当前字符，且当前字符数量
		// for 要保证最小的字符尽量放在栈底
		for len(stack) > 0 && c < stack[len(stack)-1] && hm[stack[len(stack)-1]-'a'] > 0 {
			// 设置状态为false，弹出栈顶
			visited[stack[len(stack)-1]-'a'] = false
			stack = stack[:len(stack)-1]
		}
		// 将当前字符设置为栈顶
		stack = append(stack, c)
		visited[c-'a'] = true
	}

	return string(stack)
}
