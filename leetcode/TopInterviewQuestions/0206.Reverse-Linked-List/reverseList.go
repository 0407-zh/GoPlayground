package topinterviewquestions

import (
	"github.com/0407-zh/GoPlayground/leetcode/structures"
)

type ListNode = structures.ListNode

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head.Next, prev, head = prev, head, head.Next
	}
	return prev
}
