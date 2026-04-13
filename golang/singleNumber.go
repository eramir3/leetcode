package main

func SingleNumber(nums []int) int {
	set := make(map[int]struct{})
	for _, num := range nums {
		if _, ok := set[num]; ok {
			delete(set, num)
		} else {
			set[num] = struct{}{}
		}
	}
	for num := range set {
		return num
	}
	return 0
}
