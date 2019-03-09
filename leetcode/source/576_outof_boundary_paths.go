package source

import "fmt"

var count = 0

type cords struct {
	i int
	j int
}
var steps = make(map[int][]cords)

func FindPaths(m, n, N, i, j int) int {
	findPath(m, n, N, i, j)
	for i, v := range steps {
		fmt.Println(i,":",v)
	}
	return count
}

func findPath(m, n, N, i, j int) {
	if N == 0 {
		if i < 0 || i >= m || j < 0 || j >= n {
			count++
			fmt.Println(i,j)
			steps[3] = append(steps[3], cords{i,j})
		}
		return
	}
	if i < 0 || i >= m || j < 0 || j >= n {
		count++
		fmt.Println(i,j)
		steps[3-N] = append(steps[3-N], cords{i,j})
		return
	}

	findPath(m,n,N-1,i+1,j)
	findPath(m,n,N-1,i-1,j)
	findPath(m,n,N-1,i,j+1)
	findPath(m,n,N-1,i,j-1)
}