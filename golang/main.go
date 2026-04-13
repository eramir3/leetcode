package main

import "fmt"

func main() {
	// result := TwoSum([]int{2, 7, 11, 15}, 9)
	// result := IsPalindromeString("A man, a plan, a canal - Panama!")
	// result := RomanToInteger("XCVIII")
	// result := IsValidParenthesis("({[)]})")
	// result := RemoveDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	// result := RemoveElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2)
	// result := SearchInsert([]int{1, 3, 5, 6}, 2) // Change second para to 5,7 and the answer should be 1,4 respectively
	result := LengthOfLastWord("   fly me   to   the moon  ")
	fmt.Println(result)
}
