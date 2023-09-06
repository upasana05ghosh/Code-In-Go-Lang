// /https://leetcode.com/problems/asteroid-collision/description/
func asteroidCollision(a []int) []int {
	var stack []int

	for _, v := range a {
		addFlag := true
		for len(stack) > 0 && (stack[len(stack)-1] > 0 && v < 0) {
			if math.Abs(float64(stack[len(stack)-1])) == math.Abs(float64(v)) {
				stack = stack[:len(stack)-1]
				addFlag = false
				break
			} else if math.Abs(float64(stack[len(stack)-1])) < math.Abs(float64(v)) {
				//pop prev and add new
				stack = stack[:len(stack)-1]
			} else {
				addFlag = false
				break
			}
		}
		if addFlag {
			stack = append(stack, v)
		}
	}

	return stack
}