//https://leetcode.com/problems/valid-parentheses/description/
func isValid(s string) bool {
	var stack []rune

	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, c)
		} else if len(stack) == 0 {
			return false
		} else {
			prev := stack[len(stack)-1]
			if (c == ')' && prev == '(') || (c == ']' && prev == '[') ||
				(c == '}' && prev == '{') {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}