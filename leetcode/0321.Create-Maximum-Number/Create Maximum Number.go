package leetcode

import "math"

func maxNumber(nums1 []int, nums2 []int, k int) []int {
	var result []int
	l1, l2 := len(nums1), len(nums2)
	for i := 0; i <= k; i++ {
		if i > l1 || k-i > l2 {
			continue
		}

		t := maxMergeNum(biggestSubArr(nums1, i), biggestSubArr(nums2, k-i))
		if isBigger(t, result) {
			result = t
		}
	}

	return result
}

// 合并两个数字，使合并后的数字最大，双指针
func maxMergeNum(A, B []int) (arr []int) {
	for len(A) != 0 && len(B) != 0 {
		if isBigger(A, B) {
			arr = append(arr, A[0])
			A = A[1:]
		} else {
			arr = append(arr, B[0])
			B = B[1:]
		}
	}

	if len(A) != 0 {
		arr = append(arr, A...)
	}

	if len(B) != 0 {
		arr = append(arr, B...)
	}

	return arr
}

// 从数组中选出能组成最大数的k个数字
func biggestSubArr(nums []int, k int) (result []int) {
	drop := len(nums) - k // 可以舍弃的数字数目
	for _, v := range nums {
		for len(result) != 0 && result[len(result)-1] < v && drop != 0 {
			result = result[:len(result)-1]
			drop--
		}
		result = append(result, v)
	}
	return result[:k]
}

// 返回较大的数组
func isBigger(nums1, nums2 []int) bool {
	l1, l2 := len(nums1), len(nums2)
	n := int(math.Min(float64(l1), float64(l2)))
	for i := 0; i < n; i++ {
		if nums1[i] == nums2[i] {
			continue
		}
		if nums1[i] > nums2[i] {
			return true
		}

		return false
	}
	return l1 > l2
}
