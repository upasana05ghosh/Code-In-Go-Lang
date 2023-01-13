//https://leetcode.com/problems/merge-two-sorted-lists/description/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := ListNode{0, nil}
	curr := &head

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}

	for list1 != nil {
		curr.Next = list1
		list1 = list1.Next
		curr = curr.Next
	}

	for list2 != nil {
		curr.Next = list2
		list2 = list2.Next
		curr = curr.Next
	}

	return head.Next
}