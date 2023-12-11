//https://leetcode.com/problems/min-stack/description/
type MinStack struct {
	minVal int
	stack  []int
}

func Constructor() MinStack {
	return MinStack{
		stack:  []int{},
		minVal: math.MaxInt32,
	}
}

func (this *MinStack) Push(val int) {
	if val <= this.minVal {
		this.stack = append(this.stack, this.minVal)
		this.minVal = val
	}
	this.stack = append(this.stack, val)
}

func (this *MinStack) Pop() {
	if this.stack[len(this.stack)-1] == this.minVal {
		//remove top and next and assign next to min
		this.stack = this.stack[:len(this.stack)-1]

		//remove next and assign it to min
		this.minVal = this.stack[len(this.stack)-1]
		this.stack = this.stack[:len(this.stack)-1]
	} else {
		this.stack = this.stack[:len(this.stack)-1]
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minVal
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
