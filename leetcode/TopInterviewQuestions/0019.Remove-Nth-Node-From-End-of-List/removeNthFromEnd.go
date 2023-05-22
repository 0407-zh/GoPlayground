package topinterviewquestions

import "github.com/0407-zh/GoPlayground/leetcode/structures"

type ListNode = structures.ListNode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyHead := &ListNode{Next: head}
	slow, fast := dummyHead, dummyHead

	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	for fast != nil {
		slow = slow.Next
		fast = fast.Next
	}
	slow.Next = slow.Next.Next

	return dummyHead.Next
}
