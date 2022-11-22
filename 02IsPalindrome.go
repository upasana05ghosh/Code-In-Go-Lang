// Que link - https://leetcode.com/problems/palindrome-number/description/
func isPalindrome(x int) bool {
	original := x
	rev := 0

	for x > 0 {
		rev = x%10 + rev*10
		x = x / 10
	}

	return original == rev
}