package main

func TwoSum(nums []int, target int) []int {

	map1 := make(map[int]int)

	for i, num := range nums {
		complement := target - num
		if idx, found := map1[complement]; found {
			return []int{idx, i}
		}
		map1[num] = i
	}
	return nil
}
