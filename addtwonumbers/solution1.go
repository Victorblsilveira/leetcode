package addtwonumbers

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	var sum int
	var l1next *ListNode
	var l2next *ListNode

	if l1 != nil {
		sum = sum + l1.Val
		l1next = l1.Next
	}

	if l2 != nil {
		sum = sum + l2.Val
		l2next = l2.Next
	}

	quo := sum / 10
	mod := sum % 10

	l3 := &ListNode{}
	l3.Val = mod

	if l1next != nil {
		l1next.Val = l1next.Val + quo
		l3.Next = addTwoNumbers(l1next, l2next)
	} else if l2next != nil {
		l2next.Val = l2next.Val + quo
		l3.Next = addTwoNumbers(l1next, l2next)
	} else if quo > 0 {
		l3.Next = &ListNode{Val: quo, Next: nil}
	}

	return l3
}
