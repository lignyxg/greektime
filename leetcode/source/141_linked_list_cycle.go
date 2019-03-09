package source

type ListNode struct {
	Val int
	Next *ListNode
}

func HasCycle(head *ListNode) bool {
	q, s := head, head
	for s != nil && q != nil {
		s = s.Next
		if q.Next != nil {
			q = q.Next.Next
		} else {
			return false
		}
		if s == q {
			return true
		}
	}
	return false
}
