package main

func Merge(nums1 []int, m int, nums2 []int, n int) []int {
	i := n - 1
	j := n - 1
	k := m + n - 1
	for j >= 0 {
		num1 := nums1[i]
		num2 := nums2[j]
		if i >= 0 && num1 < num2 {
			nums1[k] = num2
			j--
		} else {
			nums1[k] = num1
			i--
		}
		k--
	}
	return nums1
}
