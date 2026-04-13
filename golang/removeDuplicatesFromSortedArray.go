package main

func RemoveDuplicates(nums []int) int {
	k := 0
	pointer := 0
	for _, num := range nums {
		if num != pointer {
			nums[k] = num
			pointer = num
			k++
		}
	}
	return k
}
