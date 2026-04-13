package main

func SearchInsert(nums []int, target int) int {
	position := 0
	for idx, num := range nums {
		if num == target {
			return idx
		} else if num < target {
			position = position + 1
		}
	}
	return position
}
