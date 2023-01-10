//https://leetcode.com/problems/longest-substring-without-repeating-characters/description/
func lengthOfLongestSubstring(s string) int {
	indexMap := make(map[rune]int)
	start, end, maxLen := 0, 0, 0

	for end < len(s) {
		c := rune(s[end])
		val, ok := indexMap[c]
		if ok && val >= start {
			maxLen = max(maxLen, end-start)
			start = val + 1
		}
		indexMap[c] = end
		end++
	}
	return max(maxLen, end-start)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
