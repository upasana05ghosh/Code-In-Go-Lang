//https://leetcode.com/problems/majority-element/description/
func majorityElement(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	return nums[n/2]
}