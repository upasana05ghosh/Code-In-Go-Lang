//https://leetcode.com/problems/letter-combinations-of-a-phone-number/description/
func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	m := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var res []string
	letterComb(digits, 0, "", m, &res)
	return res
}

func letterComb(digits string, index int, curr string, m []string, res *[]string) {
	if index == len(digits) {
		*res = append(*res, curr)
		return
	}

	currNum := m[digits[index]-'0']
	for _, c := range currNum {
		letterComb(digits, index+1, curr+string(c), m, res)
	}
} 