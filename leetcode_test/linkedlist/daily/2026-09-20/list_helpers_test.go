package daily20260920

type ListNode struct {
	Val  int
	Next *ListNode
}

func listFromSlice(values []int) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for _, value := range values {
		tail.Next = &ListNode{Val: value}
		tail = tail.Next
	}
	return dummy.Next
}

func listToSlice(head *ListNode) []int {
	values := make([]int, 0)
	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}
	return values
}
