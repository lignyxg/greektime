package source

import (
	"fmt"
)

func MergeKLists(lists []*ListNode) *ListNode {
	copyList := make([]*ListNode, 0)
	for _, v := range lists {
		if v != nil {
			copyList = append(copyList, v)
		}
	}
	if len(copyList) == 0 {
		return nil
	}
	if len(copyList) == 1 {
		return copyList[0]
	}
	createHeap(copyList)
	head := new(ListNode)
	p := head
	k := len(copyList)-1 //记录合并列表大小
	for k > 0 {
		p.Next = copyList[0]
		p = p.Next
		if copyList[0].Next == nil {//堆顶链表只有一个元素了
		//类似堆排序，将末尾的链表移动至堆顶，堆减小1
			copyList[0] = copyList[k]
			k--
		} else {
			copyList[0] = copyList[0].Next //加载堆顶链表的下一个元素
		}

		adjustSmall(copyList[:k+1], 0) //自顶向下进行堆调整
	}
	if copyList[0] != nil {
		p.Next = copyList[0]
	}
	return head.Next
}

func createHeap(nodes []*ListNode) {
	for i := len(nodes)/2; i >= 0; i-- {
		adjustSmall(nodes, i)
	}
}
//自顶向下递归调整
func adjustSmall(nodes []*ListNode, index int) {
	//if index*2+2 >= len(nodes) {
	//	return
	//}
	//l, r := index*2+1, index*2+2
	//minIndex := l
	//if nodes[l].Val > nodes[r].Val {
	//	minIndex = r
	//}
	//if nodes[index].Val > nodes[minIndex].Val {
	//	tmp := nodes[index]
	//	nodes[index] = nodes[minIndex]
	//	nodes[minIndex] = tmp
	//	adjustSmall(nodes, minIndex)
	//}
	iMin, left, right := index, index*2+1, index*2+2
	if left < len(nodes) && nodes[iMin].Val > nodes[left].Val {
		iMin = left
	}
	if right < len(nodes) && nodes[iMin].Val > nodes[right].Val {
		iMin = right
	}
	if iMin != index {
		tmp := nodes[iMin]
		nodes[iMin] = nodes[index]
		nodes[index] = tmp
		adjustSmall(nodes, iMin)
	}
}

//func popHeap()

func GenerateLinkedLists(values [][]int) []*ListNode {
	lists := make([]*ListNode, len(values))
	for k, row := range values {
		p := new(ListNode)
		p.Val = row[0]
		lists[k] = p
		for i, v := range row {
			if i == 0 {
				continue
			}
			n := new(ListNode)
			n.Val = v
			p.Next = n
			p = n
		}
	}
	return lists
}

func PrintList(n *ListNode) {
	if n != nil {
		fmt.Printf("%d->",n.Val)
		PrintList(n.Next)
	} else {
		fmt.Printf("nil")
	}
}

