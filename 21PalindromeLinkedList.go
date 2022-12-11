//https://leetcode.com/problems/palindrome-linked-list/description/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}

	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	slow = reverse(slow)
	return checkIfSame(head, slow)
}

func reverse(head *ListNode) *ListNode {
	curr := head

	var prev, temp *ListNode

	for curr != nil {
		temp = curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}

	return prev
}

func checkIfSame(head *ListNode, slow *ListNode) bool {
	for head != nil && slow != nil {
		if head.Val != slow.Val {
			return false
		}
		head = head.Next
		slow = slow.Next
	}
	return true
}

