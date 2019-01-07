package leetcode

import "math"

func numSquares(n int) int {
	c := make([]int, n+1)
	c[1] = 1

	for i := 2; i <= n; i++ {
		min := math.MaxInt16
		for j := 1; j < upLimit(i); j++ {
			if x := i - j*j; x > 0 {
				if min > c[x] + 1 {
					min = c[x] + 1
				}
			}
		}
		c[i] = min
	}
	return c[n]
}

func upLimit(n int) int {
	upLimit := math.Sqrt(float64(n))
	upLimit = math.Ceil(upLimit)

	return int(upLimit)
}