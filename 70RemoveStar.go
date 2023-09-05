//https://leetcode.com/problems/removing-stars-from-a-string/description/
func removeStars(s string) string {
	var stack []rune

	for _, c := range s {
		if c != '*' {
			stack = append(stack, c)
		} else if len(stack) > 0 {
			stack = stack[:len(stack)-1]
		}
	}

	return string(stack)

}