package main

import (
	"fmt"
)

//0-1背包问题
var weight = []int{2,2,4,6,3}
var bagMax = 9

func bag(elems []int, maxWeight int) {
	num := len(elems)
	//初始化状态数组
	states := make([][]int, num)
	for i := range states {
		states[i] = make([]int, maxWeight+1)
	}

	states[0][0] = 1
	states[0][elems[0]] = 1
	for i := 1; i < num; i++ {
		//第i个物品不放入背包
		for j := 0; j <= maxWeight; j++ {
			if states[i-1][j] == 1 {
				states[i][j] = 1
			}
		}
		//第i个物品放入背包
		for j := 0; j <= maxWeight-elems[i]; j++ {
			if states[i-1][j] == 1 {
				states[i][j+elems[i]] = 1
			}
		}
	}
	//调试用，打印状态二维数组
	for _, e := range states {
		for _, n := range e {
			fmt.Printf("%d ", n)
		}
		fmt.Println("")
		/*elems == []int{2,2,4,6,3}, maxWeight == 9
			1 0 1 0 0 0 0 0 0 0
			1 0 1 0 1 0 0 0 0 0
			1 0 1 0 1 0 1 0 1 0
			1 0 1 0 1 0 1 0 1 0
			1 0 1 1 1 1 1 1 1 1
		 */
	}

	//输出结果
	for i := maxWeight; i >= 0; i-- {
		if states[num-1][i] == 1 {
			fmt.Printf("max Weight: %d\n", i)
			return
		}
	}
}

//0-1背包问题：优化存储空间
func bag2(elems []int, maxWeight int) {
	num := len(elems)
	//初始化状态空间
	states := make([]int, maxWeight+1)
	//第一个物品放入或者不放入背包的两种情况，单独处理
	states[0] = 1
	states[elems[0]] = 1
	for i := 1; i < num; i++ {//对第i个物品进行决策
		for j := maxWeight-elems[i]; j >= 0; j-- {
			/*遍历当前状态空间，放入第i个物品，若不放入物品，自然不用更新
			j从大到小遍历，因为j在遍历过程中会更新下标为"j+elems[i]"的元素，
			j+elems[i] > j，若从小到大遍历，这次的更新过程可能会干扰本轮遍历的后续更新
			 */
			if states[j] == 1 {
				states[j+elems[i]] = 1
			}
		}
	}
	//调试，打印当前状态数组
	for _, s := range states {
		fmt.Printf("%d ", s)
	}
	fmt.Println("")

	//输出结果
	for i := maxWeight; i >= 0; i-- {
		if states[i] == 1 {
			fmt.Printf("max weight: %d\n", i)
			return
		}
	}
}

//0-1背包问题：回溯方案
/*
maxWeight: 背包允许的最大重量
totalElems:物品总数
w:当前背包重量
n:要放入的物品编号
elems:物品列表
result:装入背包的物品列表
 */
func bag3(maxWeight, totalElems, w, n int, elems, result []int) {
	if w == maxWeight || n == totalElems {
		fmt.Println("max Weight: ", w)
		fmt.Println("elems: ", result)
		return
	}

	//不放入物品
	r1 := make([]int, len(result))
	copy(r1, result)
	bag3(maxWeight, totalElems, w, n+1, elems, r1)

	//放入当前物品
	r2 := make([]int, len(result))
	copy(r2, result)
	if w + elems[n] <= maxWeight {
		r2 = append(r2, elems[n])
		bag3(maxWeight, totalElems, w+elems[n], n+1, elems, r2)
	}
}

//0-1背包问题：回溯+memo方案
/*
maxWeight: 背包允许的最大重量
totalElems:物品总数
w:当前背包重量
n:要放入的物品编号
elems:物品列表
result:装入背包的物品列表
 */
 var memo [5][10]int
func bag4(maxWeight, totalElems, w, n int, elems, result []int) {
	if w == maxWeight || n == totalElems {
		fmt.Println("n=", n)
		fmt.Println("max Weight: ", w)
		fmt.Println("elems: ", result)
		return
	}
	//查看memo
	if memo[n][w] == 1 { // 状态重复
		fmt.Printf("memo[%d][%d] == 1\n", n, w)
		return
	}
	memo[n][w] = 1

	//不放入物品
	r1 := make([]int, len(result))
	copy(r1, result)
	bag4(maxWeight, totalElems, w, n+1, elems, r1)

	//放入当前物品
	r2 := make([]int, len(result))
	copy(r2, result)
	if w + elems[n] <= maxWeight {
		r2 = append(r2, elems[n])
		bag4(maxWeight, totalElems, w+elems[n], n+1, elems, r2)
	}
}

 /*
 升级版0-1背包问题：动态规划求解
 物品列表weight，对应价值value，背包总承重量maxWeight
 求在不超过背包总承重量的前提下，最大能装入背包的价值
  */
func bag5(maxWeight int, weight, value []int) {
	if maxWeight == 0 || weight == nil || value == nil {
		return
	}
	//定义状态数组, states[i][j]表示: 考察（装入/不装入）第i个物品且当前总重为j时，背包中容纳的总价值
	states := make([][]int, len(weight)+1)
	for i := 0; i < len(states); i++ {
		s := make([]int, maxWeight+1)
		states[i] = s
	}
	for i := 1; i <= len(weight); i++ {
		for j := 0; j <= maxWeight; j++ {
			if states[i-1][j] > 0 {
				states[i][j] = states[i-1][j]
			}
		}
		for j := 0; j <= maxWeight-weight[i]; j++ {
			v := states[i-1][j] + value[i]
			if states[i-1][j+weight[i]] < v {
				states[i][j+weight[i]] = v
			}
		}
	}
	maxValue := -1
	for k := 0; k <= maxWeight; k++ {
		if v := states[len(weight)][k]; v > maxValue {
			maxValue = v
		}
	}
	fmt.Printf("max Value: %d\n", maxValue)
}

func main() {
	//bag(weight, bagMax)
	//bag2(weight, bagMax)
	//bag3(bagMax, len(weight), 0, 0, weight, make([]int, 0))
	bag4(bagMax, len(weight), 0, 0, weight, make([]int, 0))
}