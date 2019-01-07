package source

import (
	"fmt"
	"math"
)

func NumSquares(n int) int {
	c := make([]int, n+1) // 定义状态数组
	c[0] = 0
	c[1] = 1
	/*
	状态转移方程:
	c[i] = max(c[i-j*j] + 1), j = 1 ... Sqrt(i)
	 */
	for i := 2; i <= n; i++ {
		min := math.MaxInt16
		upl := upLimit(i)
		for j := 1; j <= upl; j++ {
			if x := i - j*j; x >= 0 {
				if min > c[x] + 1 {
					min = c[x] + 1
				}
			}
		}
		c[i] = min
	}
	fmt.Println(c)
	return c[n]
}

func upLimit(n int) int {
	upLimit := math.Sqrt(float64(n))
	upLimit = math.Ceil(upLimit)

	return int(upLimit)
}