//https://leetcode.com/problems/product-of-array-except-self/description/
func productExceptSelf(a []int) []int {
	//Ref - https://www.youtube.com/watch?v=gREVHiZjXeQ&ab_channel=Techdose
	n := len(a)
	left := make([]int, n)

	left[0] = a[0]
	for i := 1; i < n; i++ {
		left[i] = a[i] * left[i-1]
	}

	temp := a[n-1]
	left[n-1] = left[n-2]

	for i := n - 2; i > 0; i-- {
		left[i] = left[i-1] * temp
		temp = temp * a[i]
	}

	left[0] = temp
	return left
}