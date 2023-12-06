//https://leetcode.com/problems/minimum-remove-to-make-valid-parentheses/description/
func minRemoveToMakeValid(s string) string {
	var stack []int
	var sb strings.Builder

	ch := []rune(s)

	for i, c := range s {
		if c == '(' {
			stack = append(stack, i) //noted the location of '(' in case we plann to remove it
		} else if c == ')' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1] //if find a valid combination
			} else {
				ch[i] = '#' //invalid
			}
		}
	}

	//mark invalid ')' in the stack
	for _, v := range stack {
		ch[v] = '#'
	}

	for _, v := range ch {
		if v != '#' {
			sb.WriteRune(v)
		}
	}

	return sb.String()

}