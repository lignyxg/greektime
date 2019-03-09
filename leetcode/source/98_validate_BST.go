package source

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	_, _, valid := maxMinValid(root)
	return valid
}

func maxMinValid(root *TreeNode) (max, min int, valid bool) {
	if root.Left == nil && root.Right == nil {
		return root.Val, root.Val, true
	}

	lMax, lMin, lValid := 0, 0, true
	rMax, rMin, rValid := 0, 0, true
	if root.Left != nil {
		lMax, lMin, lValid = maxMinValid(root.Left)
		if !lValid {
			return lMax, lMin, lValid
		}
		if root.Val <= lMax { //左子树最大值应该小于根节点
			return root.Val, lMin, false
		}
	} else { //左子树为空
		lMin = root.Val
	}
	if root.Right != nil {
		rMax, rMin, rValid = maxMinValid(root.Right)
		if !rValid {
			return rMax, rMin, rValid
		}
		if root.Val >= rMin { //右子树最小值应该大于根节点
			return root.Val,rMin,false
		}
	} else { //右子树为空
		rMax = root.Val
	}
	return rMax, lMin, true
}

func isValidBST2(root *TreeNode) bool {
	var flag bool
	helper(root, nil, nil, &flag)
	return flag
}

func helper(root, min, max *TreeNode, flag *bool) {
	if root == nil {
		return
	}
	if (min!=nil && root.Val<=min.Val) || (max!=nil && root.Val>=max.Val) {
		*flag = false
		return
	}
	helper(root.Left, min, root, flag)
	helper(root.Right, root, max, flag)
}

func isValidBST3(root *TreeNode) bool {
	if root == nil {
		return true
	}
	st := make([]*TreeNode, 1000) // 栈
	cur := 0 // 栈顶
	p := root
	firstTime, lastVal := 0, 0
	for p != nil || cur > 0 {
		for p != nil {
			st[cur] = p
			cur++
			p = p.Left
		}
		if cur > 0 {
			// 出栈
			cur--
			p = st[cur]
			if firstTime == 0 {
				firstTime = 1
			} else {
				if p.Val <= lastVal {
					return false
				}
			}
			lastVal = p.Val
			//右孩子准备入栈
			p = p.Right
		}
	}
	return true
}