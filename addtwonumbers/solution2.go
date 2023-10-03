package addtwonumbers

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}

	current1Node := l1
	current2Node := l2
	current3Node := l3
	for current1Node != nil || current2Node != nil {
		var sum int
		var l1next *ListNode
		var l2next *ListNode

		if current1Node != nil {
			sum = sum + current1Node.Val
			l1next = current1Node.Next
		}

		if current2Node != nil {
			sum = sum + current2Node.Val
			l2next = current2Node.Next
		}

		quo := sum / 10
		mod := sum % 10

		current3Node.Val = mod
		if l1next != nil {
			l1next.Val = l1next.Val + quo
			current3Node.Next = &ListNode{}
		} else if l2next != nil {
			l2next.Val = l2next.Val + quo
			current3Node.Next = &ListNode{}
		} else if quo > 0 {
			current3Node.Next = &ListNode{Val: quo, Next: nil}
			break
		}

		current1Node = l1next
		current2Node = l2next
		current3Node = current3Node.Next
	}

	return l3
}
