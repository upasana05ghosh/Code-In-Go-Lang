//https://leetcode.com/problems/unique-number-of-occurrences/description/
func uniqueOccurrences(arr []int) bool {
	occurance_map := make(map[int]int)
	for _, val := range arr { //index, val
		occurance_map[val] += 1
	}

	set := make(map[int]bool)
	for _, count := range occurance_map { //key, val
		if v, ok := set[count]; ok { //val -> true/false
			fmt.Print(v)
			return false
		}
		set[count] = true
	}
	return true
}