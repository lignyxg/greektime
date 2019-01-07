package main

import (
	"fmt"
	"math"
)

/*
从矩阵的左上角到右下角的最短路径
回溯
 */
func FindshortestPath(m [][]int, x, y, n int) int {
	fmt.Printf("x=%d,y=%d\n",x,y)
	if x == n-1 && y == n-1 {
		return m[x][y]
	}
	if x >= n || y >= n {
		return math.MaxInt16
	}
	p1, p2 := 0, 0
	if x < n {
		p1 = FindshortestPath(m, x+1, y, n)
	}
	if y < n {
		p2 = FindshortestPath(m, x, y+1, n)
	}
	min := 0
	if p1 < p2 {
		min = p1
	} else {
		min = p2
	}
	return m[x][y]+min
}

var minDist = math.MaxInt16
func FindshortestPath2(m [][]int, dist, x, y, n int) {
	//fmt.Printf("dist=%d,x=%d,y=%d\n",dist,x,y)
	//if (x == n && y == n-1) || (x == n-1 && y == n) {
	//	fmt.Println("length: ", dist)
	//	if minDist > dist {
	//		minDist = dist
	//	}
	//}
	if x == n-1 && y == n-1 {
		if minDist > dist+m[x][y] {
			minDist = dist+m[x][y]
		}
		return
	}
	if x >= n || y >= n {
		return
	}
	if x < n {
		FindshortestPath2(m, dist+m[x][y], x+1, y, n)
	}
	if y < n {
		FindshortestPath2(m, dist+m[x][y], x, y+1, n)
	}
}
/*

 */
func FindshortestPath3(m [][]int, n int) {
	states := make([][]int, n)
	for i := 0; i < n; i++ {
		e := make([]int, n)
		states[i] = e
	}
	//初始化
	sum := 0
	for i := 0; i < n; i++ {
		sum += m[0][i]
		states[0][i] = sum
	}
	sum = 0
	for i := 0; i < n; i++ {
		sum += m[i][0]
		states[i][0] = sum
	}
	for i := 1; i < n; i++ {
		for j := 1; j < n; j++ {
			if states[i-1][j] < states[i][j-1] {
				states[i][j] = states[i-1][j]
			} else {
				states[i][j] = states[i][j-1]
			}
			states[i][j] += m[i][j]
		}
	}
	fmt.Printf("shortest path: %d\n", states[n-1][n-1])
}

func main() {
	aMap := [][]int{
		{1,3,5,9},
		{2,1,3,4},
		{5,2,6,7},
		{6,8,4,3},
	}
	//aMap2 := [][]int{
	//	{2,1,4,6},
	//	{4,7,9,1},
	//	{2,7,9,10},
	//	{5,3,8,7},
	//}
	min := FindshortestPath(aMap,0,0,4)
	fmt.Printf("first method: %d\n", min)
	FindshortestPath2(aMap,0,0,0,4)
	fmt.Printf("second method: %d\n", minDist)
	FindshortestPath3(aMap, 4)
}