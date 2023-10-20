//https://leetcode.com/problems/number-of-recent-calls/description/
type RecentCounter struct {
	req []int
}

func Constructor() RecentCounter {
	return RecentCounter{req: make([]int, 0)}
}

func (this *RecentCounter) Ping(t int) int {
	this.req = append(this.req, t)
	min := t - 3000
	i := bs(this.req, min)
	return len(this.req) - i
}

func bs(a []int, k int) int {
	l, r := 0, len(a)-1

	for l <= r {
		mid := (l + r) / 2
		if a[mid] == k {
			return mid
		}
		if a[mid] < k {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */