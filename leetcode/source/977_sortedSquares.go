package source

func SortedSquares(A []int) []int {
	l := len(A)
	res := make([]int, l)
	sqrt := func(x int) int {
		return x * x
	}

	if l == 0 || l == 1 {
		res[0] = sqrt(A[0])
		return res
	}

	if A[0] >= 0 {
		for i := 0; i < l; i++ {
			res[i] = sqrt(A[i])
		}
		return res
	}
	i, j := 0, l-1
	k := l-1
	for i <= j {
		if A[i] + A[j] <= 0 {
			res[k] = sqrt(A[i])
			i++
		} else {
			res[k] = sqrt(A[j])
			j--
		}
		k--
	}
	return res
}
