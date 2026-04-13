package main

func IsValidParenthesis(s string) bool {
	map1 := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}

	stack := []byte{}
	for i := 0; i < len(s); i++ {
		char := s[i]
		if expected, ok := map1[char]; ok {
			top := stack[len(stack)-1]
			if top != expected {
				return false
			}
			stack = stack[:len(stack)-1] // pop
		} else {
			stack = append(stack, char)
		}
	}
	return len(stack) == 0
}
