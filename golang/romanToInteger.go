package main

func RomanToInteger(s string) int {
	map1 := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	sum := 0
	last := 0
	for i := len(s) - 1; i >= 0; i-- {
		current := map1[s[i]]
		if current >= last {
			sum += current
		} else {
			sum -= current
		}
		last = current
	}
	return sum
}
