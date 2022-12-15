//https://leetcode.com/problems/intersection-of-two-arrays/description/
func intersection(nums1 []int, nums2 []int) []int {
	var setMap = map[int]bool{}
	var res = []int{}

	for _, v := range nums1 {
		setMap[v] = true
	}

	for _, v := range nums2 {
		if setMap[v] == true {
			res = append(res, v)
			setMap[v] = false
		}
	}

	return res
}
