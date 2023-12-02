//https://leetcode.com/problems/generate-parentheses/description/
func generateParenthesis(n int) []string {
	var res []string
	var temp string
	generate(&res, temp, n, 0)
	return res
}

func generate(res *[]string, temp string, left, right int) {
	if left == 0 && right == 0 {
		//may need deep copy
		*res = append(*res, temp)
		return
	}

	if left > 0 {
		generate(res, temp+"(", left-1, right+1)
	}

	if right > 0 {
		generate(res, temp+")", left, right-1)
	}
}

