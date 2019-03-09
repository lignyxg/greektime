package source

import (
	"math"
)

func MinimumTotal(triangle [][]int) int {
	m := len(triangle)
	if m == 0 {
		return 0
	}
	state := make([]int, m)
	state[0] = triangle[0][0]
	for i := 1; i < m; i++ {
		for j := i; j >= 0; j-- { // 从后向前更新，否则在向后更新的过程中会用到已经更新过的脏值
			min := math.MaxInt32
			if j-1 >= 0 {
				min = triangle[i][j] + state[j-1]
			}
			if j < i && min > triangle[i][j] + state[j] {
				min = triangle[i][j] + state[j]
			}
			state[j] = min
		}
		//fmt.Println(state)
	}
	min := state[0]
	for _, v := range state {
		if min > v {
			min = v
		}
	}
	return min
}
