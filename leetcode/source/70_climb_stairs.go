package source

func ClimbStairs(n int) int {
	mem := make([]int, n)
	m := climb(mem, n)
	return m
}

func climb(mem []int, n int) int {
	if n == 1 || n == 2 {
		return n
	}
	if mem[n-1] == 0 {
		mem[n-1] = climb(mem, n-1)
	}
	if mem[n-2] == 0 {
		mem[n-2] = climb(mem, n-2)
	}
	return mem[n-1] + mem[n-2]
}
