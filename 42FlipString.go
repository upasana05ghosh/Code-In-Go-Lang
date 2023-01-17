//https://leetcode.com/problems/flip-string-to-monotone-increasing/description/
func minFlipsMonoIncr(s string) int {
	//dp
	res, cntOne := 0, 0

	for _, val := range s {
		r := rune(val)
		if r == '1' {
			cntOne++
		} else {
			//when it's 0 -> either convert 0 to 1 val (res +1) or convert all past 1 to 0
			res = min(res+1, cntOne)
		}
	}

	return res

}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}