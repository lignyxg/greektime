package main

import "fmt"

var results = make([][]int, 0)
/*
最长递增子序列
动态规划求解
 */
func maxAscendList(m []int) int {
	states := make([]int, len(m))
	states[0] = 1
	res := make([]int, len(m))//记录结果元素的索引
	for i := 1; i < len(m); i++ {
		max := 0
		trace := 0
		for j := 0; j < i; j++ {
			tmp := states[j]
			if m[j] < m[i] {
				tmp += 1
			}
			if tmp > max {
				max = tmp
				trace = j
			}
		}
		states[i] = max
		res[i] = trace
	}
	fmt.Println(states)
	fmt.Println(res)
	return states[len(m)-1]
}

func main() {
	a := []int{2,9,3,5,6,1,7}
	maxLength := maxAscendList(a)
	fmt.Println(maxLength)
}