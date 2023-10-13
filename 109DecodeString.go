//https://leetcode.com/problems/decode-string/description/
func decodeString(s string) string {
	if len(s) == 0 {
		return s
	}

	// track backward
	var stack []string
	for i := len(s) - 1; i >= 0; i-- {
		c := s[i]

		if isDigit(c) {
			//build the number
			n, r := 0, 1
			for i >= 0 && isDigit(s[i]) {
				n += int(s[i]-'0') * r
				r *= 10
				i--
			}
			fmt.Println(n)
			i++                          // reset
			stack = stack[:len(stack)-1] //pop [

			//pop string till ]
			var sb strings.Builder
			for stack[len(stack)-1] != "]" {
				sb.WriteString(stack[len(stack)-1])
				stack = stack[:len(stack)-1] //pop it
			}

			stack = stack[:len(stack)-1] //pop ]

			//repeat sb till n times
			var final strings.Builder
			for n > 0 {
				final.WriteString(sb.String())
				n--
			}

			// put it back to stack
			stack = append(stack, final.String())

		} else {
			//add it to stack
			stack = append(stack, string(c))
		}
	}

	var sb strings.Builder
	for len(stack) > 0 {
		sb.WriteString(stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}

	return sb.String()

}

func isDigit(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	return false
}