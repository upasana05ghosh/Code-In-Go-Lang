//https://leetcode.com/problems/evaluate-reverse-polish-notation/description/
func evalRPN(tokens []string) int {
	var stack []int
	var n, x, y int

	for _, val := range tokens {
		v, err := strconv.Atoi(val)

		if err == nil {
			stack = append(stack, v) //it's a number, push it to stack
		} else {
			n = len(stack)
			if n > 1 {
				x = stack[n-1]
				y = stack[n-2]
				stack = stack[:n-2] //pop both element
				stack = append(stack, getVal(val, x, y))
			}
		}
	}
	return stack[0]
}

func getVal(val string, x int, y int) int {
	switch val {
	case "+":
		return x + y
	case "-":
		return y - x
	case "*":
		return y * x
	case "/":
		return y / x
	}
	return -1
}