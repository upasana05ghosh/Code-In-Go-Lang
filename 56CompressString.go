//https://leetcode.com/problems/string-compression/description/
func compress(chars []byte) int {
	prev, i, updateIndex, count := 0, 0, 0, 0

	for i < len(chars) {
		//Reset values
		prev = i
		count = 0

		for i < len(chars) && chars[i] == chars[prev] {
			i++
			count++
		}

		chars[updateIndex] = chars[prev]
		updateIndex++
		if count > 1 {
			//If count is 12, we need to set it as '1' '2'
			for _, c := range strconv.Itoa(count) {
				chars[updateIndex] = byte(c)
				updateIndex++
			}
		}

	}
	return updateIndex
}