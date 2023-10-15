//https://leetcode.com/problems/maximum-subsequence-score/description/
func maxScore(n1 []int, n2 []int, k int) int64 {
	n := len(n1)
	a := make([][2]int, n)

	for i := 0; i < n; i++ {
		a[i] = [2]int{n2[i], n1[i]}
	}

	// sort in descending
	sort.Slice(a, func(i, j int) bool {
		return a[j][0] < a[i][0]
	})

	minHeap := &MinHeap{}
	var sum, res int64

	for i := 0; i < n; i++ {
		sum += int64(a[i][1])
		heap.Push(minHeap, a[i][1])

		if minHeap.Len() > k {
			sum -= int64(heap.Pop(minHeap).(int))
		}

		if minHeap.Len() == k {
			res = max(res, sum*int64(a[i][0]))
		}
	}

	return res
}

type MinHeap []int

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h MinHeap) Top() interface{} {
	return h[0]
}

func max(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}