package main

import (
	"fmt"
	"math"
)

/*
寻找w1和w2的最小编辑距离
 */
func minEditDis(w1, w2 []byte) int {
	if w1 == nil || w2 == nil {
		return 0
	}
	m, n := len(w1), len(w2)
	d := make([][]int, m+1) // 子问题状态二维数组
	for i := 0; i <= m; i++ {
		s := make([]int, n+1)
		s[0] = i // 初始化w2为空串的情况
		d[i] = s
	}
	//初始化w1为空串的情况
	for j := 1; j <= n; j++ {
		d[0][j] = j
	}
	// 迭代求解子问题
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			first := d[i-1][j] + 1 // w1增加一个字符
			second := d[i][j-1] + 1 // w2增加一个字符
			third := 0
			if w1[i-1] != w2[j-1] { // 这里i和j是字符串长度，因此i-1和j-1才是下标
				third = 1
			}
			third += d[i-1][j-1] // w1，w2同时增加一个字符
			//d[i][j]取三者中的最小值
			// 1)
			if first < second {
				d[i][j] = first
			} else {
				d[i][j] = second
			}
			// 2)
			if third < d[i][j] {
				d[i][j] = third
			}
		}
	}
	return d[m][n]
}


var v = []int{1,2,5,10} // 面值
var res = make([]int,0) // 最终结果数组
var flag = false // res是否已经赋值过，false代表从未赋值
/*
给定总额和一定面值的纸币，求达到所需的总额最少需要几张纸币
递归求解
 */
func FindMinCombine1(totalAmount, target int, result []int) {
	if totalAmount == target { // 找到结果
		if !flag || len(res) > len(result) {
			res = result
			flag = true
		}
		return
	}
	if totalAmount > target {
		return
	}
	for i := 0; i < len(v); i++ {
		nextLevel := make([]int, len(result))
		copy(nextLevel, result)
		nextLevel = append(nextLevel, v[i])
		FindMinCombine1(totalAmount+v[i], target, nextLevel)
	}
}
//动态规划求解
// target: 目标总金额
// value: 面值数组
func FindMinCombine2(target int, value []int) int {
	if target <= 0 || value == nil {
		return -1
	}
	c := make([]int, target+1) //c[i]表示当总金额为i时，最少的钱币张数
	for i := 1; i <= target; i++ {
		min := math.MaxInt16
		for j := 0; j < len(value); j++ {
			if i - value[j] >= 0 {
				if min > c[i-value[j]] {
					min = c[i-value[j]]
				}
			}
		}
		c[i] = min + 1
	}
	return c[target]
}

func main() {

	//var b1, b2 bytes.Buffer
	//b1.WriteString("mou")
	//b2.WriteString("m")
	//dis := minEditDis(b1.Bytes(), b2.Bytes())
	//fmt.Printf("min edit distense between %s and %s is %d", b1.String(), b2.String(), dis)

	// 最小钱币数量
	FindMinCombine1(0,15,make([]int, 0))
	fmt.Println(res)
	n := FindMinCombine2(15,v)
	fmt.Printf("min combine: %d", n)
}