package main

import (
	"fmt"
	"greektime/leetcode/source"
)

func main() {
	//fmt.Println(source.NumSquares(12))
	//a := []int{-7,-3,2,3,11}
	//b := source.SortedSquares(a)
	//fmt.Println(b)

	//nums := []int{-2,0,0,2,2}
	//res := source.ThreeSum1(nums)
	//for _, item := range res {
	//	fmt.Println(item)
	//}

	//a := []int{3,4,-1,1}
	//fmt.Println(source.FirstMissingPositive(a))

	//listVals := [][]int{
	//	{1,4,5},
	//	{1,3,4},
	//	{2,6},
	//}
	//lists := source.GenerateLinkedLists(listVals)
	//fmt.Println("lists:")
	//lists = []*source.ListNode{
	//	nil,nil,
	//}
	//for _,h := range lists {
	//	source.PrintList(h)
	//	fmt.Println()
	//}
	//
	//h := source.MergeKLists(lists)
	//
	//source.PrintList(h)

	//Test Stack
	//st := source.NewStack(2)
	//fmt.Println(st.IsEmpty(), st.IsFull())
	//st.Push(1)
	//st.Push(2)
	//fmt.Println(st.IsFull())
	//fmt.Println(st.Peek())
	//fmt.Println(st.Pop(),st.Pop())
	//fmt.Println(st.IsEmpty())
	//fmt.Println(st.Peek())

	//最长匹配括号
	//s := "()()((()()())"
	////s = ")()())"
	////s = "()(())"
	//s = ")(((((()())()()))()(()))("
	//s = "()()"
	//fmt.Println(source.LongestValidParentheses(s))
	//fmt.Println(source.LongestValidParentheses2(s))

	//逆波兰表达式求值
	//tokens := []string{"2","1","+","3","*"}
	//fmt.Println(source.EvalRPN(tokens))

	//球出界路径
	//fmt.Println(source.FindPaths(2,2,3,0,0))

	//爬楼梯
	//fmt.Println(source.ClimbStairs(4))

	//x的平方根
	//fmt.Println(source.MySqrt(10))

	// 反转字符串里的单词
	//fmt.Println(source.ReverseWords("   a   b  c d   e  "))
	//字符串转数字
	//fmt.Printf("%d", source.MyAtoi("  0000000000012345678"))

	//正则表达式匹配
	//s, p := "b","b*"
	//s, p = "",".*"
	//fmt.Printf("%s, %s: %v", s, p, source.IsMatch2(s,p))

	//零钱兑换
	//coins := []int{2,5,1}
	//amount := 11
	//fmt.Printf("coins change(amount:%d):%d",amount, source.CoinChange(coins, amount))

	//乘积最大子序列
	//nums := []int{-2,0,-1}
	////nums = []int{-4,-3}
	////nums = []int{0,2}
	////nums = []int{-2,3,-4}
	//nums = []int{3,-1,4}
	//nums = []int{2,-5,-2,-4,3}
	//fmt.Printf("%v max product: %d", nums, source.MaxProduct2(nums))

	//三角形最小路径
	tri := [][]int {
		{-1},{3,2},{-3,1,-1},
	}
	fmt.Printf("triangle: %d\n", source.MinimumTotal(tri))
}

