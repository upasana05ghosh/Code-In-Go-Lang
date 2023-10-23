//https://leetcode.com/problems/online-stock-span/description/
type StockSpanner struct {
	stack [][2]int
}

func Constructor() StockSpanner {
	return StockSpanner{stack: make([][2]int, 0)}
}

func (this *StockSpanner) Next(price int) int {
	res := 1

	for len(this.stack) != 0 && this.stack[len(this.stack)-1][0] <= price {
		res += this.stack[len(this.stack)-1][1]
		this.stack = this.stack[:len(this.stack)-1]
	}
	this.stack = append(this.stack, [2]int{price, res})
	return res
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */