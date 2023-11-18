//https://leetcode.com/problems/group-anagrams/description/
func groupAnagrams(s []string) [][]string {
	var res [][]string

	if len(s) == 0 {
		return res
	}

	mp := make(map[[26]int][]string)

	for _, v := range s {
		var a [26]int

		for i := 0; i < len(v); i++ {
			a[v[i]-'a']++
		}

		mp[a] = append(mp[a], v)
	}

	for _, v := range mp {
		res = append(res, v)
	}

	return res
}