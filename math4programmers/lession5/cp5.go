package main

import (
	"fmt"
)

var results1 = make([][]int, 0)
// 整数分解
func FindAllMultiplicators(num int, res []int) {
	if num == 1 {
		if !contains(res, 1) {
			res = append(res, 1)
		}
		results1 = append(results1, res)
		return
	} else if num < 1 {
		return
	}
	for i := 1; i <= num; i++ {
		newRes := make([]int, len(res))
		copy(newRes, res)
		if contains(newRes, 1) && i == 1 { // 保证结果中1只出现1次
			continue
		}
		if num/i * i < num { // 说明num/i不是整数
			continue
		}
		newRes = append(newRes, i)
		FindAllMultiplicators(num/i, newRes)
	}
}

func contains(s []int, e int) bool {
	for _, n := range s {
		if e == n {
			return true
		}
	}
	return false
}

/*
赏金问题: 给定总额，面值1，2，5，10的纸币，总额分别由以上面值中的一种或多种的纸币组成，
总共可以有多少种组合方式
 */
var amount = []int{1,2,5,10}
var results2 = make([][]int, 0)

func Reward(total int, res []int) {
	if total == 0 {
		results2 = append(results2, res)
		return
	} else if total < 0 {
		return
	}
	for i := 0; i < len(amount); i++ {
		newRes := make([]int, len(res))
		copy(newRes, res)
		newRes = append(newRes, amount[i])
		Reward(total-amount[i], newRes)
	}
}


func main() {
	res1 := make([]int, 0)
	FindAllMultiplicators(8, res1)
	for _, s := range results1 {
		fmt.Println(s)
	}

	fmt.Println("=======================")
	res2 := make([]int, 0)
	Reward(10, res2)
	for _, s := range results2 {
		fmt.Println(s)
	}
}