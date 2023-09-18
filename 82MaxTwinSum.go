//https://leetcode.com/problems/maximum-twin-sum-of-a-linked-list/description
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func pairSum(head *ListNode) int {
	if head == nil {
		return 0
	}

	//go till half
	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	//reverse the second half
	revHead := reverse(slow.Next)

	//break first
	slow.Next = nil

	//find max sum
	max := 0

	for head != nil && revHead != nil {
		sum := head.Val + revHead.Val
		if sum > max {
			max = sum
		}
		head = head.Next
		revHead = revHead.Next
	}

	return max
}

func reverse(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}